package entities

import (
	"context"
	"errors"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

const accountTableName = "account"

// AccountDao is the exported data access object for queries, e.g., CRUD
var AccountDao accountDao

// accountDao is the interface lists all functions of the entity
type accountDao interface {
	Create(ctx context.Context, e *Account) error
	Update(ctx context.Context, e *Account) error
	Get(ctx context.Context, id uint) (*Account, error)
	GetByAddress(ctx context.Context, address string) (*Account, error)
	GetForUpdate(ctx context.Context, id uint) (*Account, error)
	GetByAddressForUpdate(ctx context.Context, address string) (*Account, error)
	CreateOrGetByAddress(ctx context.Context, e *Account) (bool, error)
	GetAndIncreaseUserNonce(ctx context.Context, id uint) (uint, error)
	SetUserNonce(ctx context.Context, id, nonce uint) error
	HoldBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error
	AddBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error
	DecreaseHoldBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error
	IncreaseHoldBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error
	UnholdBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error
	DecreaseBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error
}

type accountDaoImpl struct {
	*mysql.Dao
}

// Account is a struct
type Account struct {
	gorm.Model
	Address   string
	Nonce     uint
	Role      string
	Hold      decimal.Decimal
	Available decimal.Decimal

	FromAllowances []*Allowance `gorm:"-"`
	ToAllowances   []*Allowance `gorm:"-"`
}

// TableName get/set the table name of the entity
func (e *Account) TableName() string {
	return accountTableName
}

func (e *Account) FillAllowance(ctx context.Context) error {
	fromAllowances, err := AllowanceDao.GetsByFromAccountID(ctx, e.ID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	toAllowances, err := AllowanceDao.GetsByToAccountID(ctx, e.ID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	e.FromAllowances = fromAllowances
	e.ToAllowances = toAllowances
	return nil
}

func (i *accountDaoImpl) Create(ctx context.Context, e *Account) error {
	return i.Db(ctx).Create(e).Error
}

func (i *accountDaoImpl) Update(ctx context.Context, e *Account) error {
	return i.Db(ctx).Save(e).Error
}

func (i *accountDaoImpl) CreateOrGetByAddress(ctx context.Context, e *Account) (bool, error) {
	isNewUser := false
	db := i.Db(ctx)

	a := &Account{}
	db = i.Db(ctx).Unscoped().
		Where("address = ?", e.Address).First(a)
	if db.RowsAffected == 0 {
		e.Nonce = 0
		db = i.Db(ctx).Create(e)
		isNewUser = true
	} else {
		*e = *a
	}

	return isNewUser, db.Error
}

func (i *accountDaoImpl) Get(ctx context.Context, accountId uint) (*Account, error) {
	res := &Account{}
	if err := i.Db(ctx).Where("id = ?", accountId).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (i *accountDaoImpl) GetByAddress(ctx context.Context, address string) (*Account, error) {
	res := &Account{}
	if err := i.Db(ctx).Where("address = ?", address).First(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (i *accountDaoImpl) GetForUpdate(ctx context.Context, accountId uint) (*Account, error) {
	res := &Account{}
	err := i.Db(ctx).Raw("SELECT * FROM account WHERE id = ? FOR UPDATE", accountId).Scan(res).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return res, err
}

func (i *accountDaoImpl) GetByAddressForUpdate(ctx context.Context, address string) (*Account, error) {
	res := &Account{}
	err := i.Db(ctx).Raw("SELECT * FROM account WHERE address = ? FOR UPDATE", address).Scan(res).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return res, err
}

func (i *accountDaoImpl) GetAndIncreaseUserNonce(ctx context.Context, accountId uint) (uint, error) {
	account := &Account{Model: gorm.Model{ID: accountId}}
	err := i.Db(ctx).Find(account).Error
	if err != nil {
		return 0, err
	}
	currNonce := account.Nonce

	err = i.Db(ctx).Model(account).Update("nonce", currNonce+1).Error
	if err != nil {
		return 0, err
	}

	return currNonce + 1, nil
}

func (i *accountDaoImpl) SetUserNonce(ctx context.Context, accountId, nonce uint) error {
	account := &Account{Model: gorm.Model{ID: accountId}}
	err := i.Db(ctx).Find(account).Error
	if err != nil {
		return err
	}
	if nonce <= account.Nonce {
		return errors.New("nonce too small")
	}

	err = i.Db(ctx).Model(account).Update("nonce", nonce).Error
	return err
}

// HoldBalance
// 1. decrease user's available balance
// 2. increase user's hold balance
func (i *accountDaoImpl) HoldBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error {
	if size.LessThanOrEqual(decimal.Zero) {
		return errors.New("HoldBalance size less than 0")
	}

	enough, err := i.hasEnoughBalance(ctx, account.ID, size)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if !enough {
		return errors.New(fmt.Sprintf("less balance than: %v", size))
	}

	account, err = i.GetForUpdate(ctx, account.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("no account")
	}

	account.Available = account.Available.Sub(size)
	account.Hold = account.Hold.Add(size)

	transaction := &Transaction{
		AccountId: account.ID,
		Available: size.Neg(),
		Hold:      size,
		Type:      transactionType.String(),
		Notes:     null.StringFrom(note),
	}
	err = TransactionDao.BatchCreate(ctx, []*Transaction{transaction})
	if err != nil {
		return err
	}

	err = i.Update(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

// AddBalance
// 1. increase user's available balance in DB
func (i *accountDaoImpl) AddBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error {
	if size.LessThanOrEqual(decimal.Zero) {
		return errors.New("AddBalance size less than 0")
	}

	existingAccount, err := i.GetByAddressForUpdate(ctx, account.Address)
	if err != nil {
		return err
	}
	if existingAccount == nil {
		newAccount := &Account{
			Address:   account.Address,
			Role:      types.AccountRoleUser.String(),
			Available: decimal.Zero,
			Hold:      decimal.Zero,
		}
		err = i.Create(ctx, newAccount)
		if err != nil {
			return err
		}
		account, err = i.GetForUpdate(ctx, newAccount.ID)
		if err != nil {
			return err
		}
	} else {
		account = existingAccount
	}
	account.Available = account.Available.Add(size)

	transaction := &Transaction{
		AccountId: account.ID,
		Available: size,
		Hold:      decimal.Zero,
		Type:      transactionType.String(),
		Notes:     null.StringFrom(note),
	}
	err = TransactionDao.BatchCreate(ctx, []*Transaction{transaction})
	if err != nil {
		return err
	}

	err = i.Update(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

// DecreaseHoldBalance
// 1. decrease user's hold balance in DB
func (i *accountDaoImpl) DecreaseHoldBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error {
	if size.LessThanOrEqual(decimal.Zero) {
		return errors.New("DecreaseHoldBalance size less than 0")
	}

	enough, err := i.hasEnoughHoldBalance(ctx, account.ID, size)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if !enough {
		return errors.New(fmt.Sprintf("no enough hold balance for %v", size))
	}

	account, err = i.GetForUpdate(ctx, account.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("no account")
	}
	account.Hold = account.Hold.Sub(size)

	transaction := &Transaction{
		AccountId: account.ID,
		Available: decimal.Zero,
		Hold:      size.Neg(),
		Type:      transactionType.String(),
		Notes:     null.StringFrom(note),
	}
	err = TransactionDao.BatchCreate(ctx, []*Transaction{transaction})
	if err != nil {
		return err
	}

	err = i.Update(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

// IncreaseHoldBalance
// 1. increase user's hold balance in DB
func (i *accountDaoImpl) IncreaseHoldBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error {
	var err error
	if size.LessThanOrEqual(decimal.Zero) {
		return errors.New("DecreaseHoldBalance size less than 0")
	}

	account, err = i.GetForUpdate(ctx, account.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("no account")
	}
	account.Hold = account.Hold.Add(size)

	transaction := &Transaction{
		AccountId: account.ID,
		Available: decimal.Zero,
		Hold:      size,
		Type:      transactionType.String(),
		Notes:     null.StringFrom(note),
	}
	err = TransactionDao.BatchCreate(ctx, []*Transaction{transaction})
	if err != nil {
		return err
	}

	err = i.Update(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

// UnholdBalance
// 1. decrease user's hold balance
// 2. increase user's available balance
func (i *accountDaoImpl) UnholdBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error {
	if size.LessThanOrEqual(decimal.Zero) {
		return errors.New("UnholdBalance size less than 0")
	}

	enough, err := i.hasEnoughHoldBalance(ctx, account.ID, size)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if !enough {
		return errors.New(fmt.Sprintf("no enough hold balance for %v", size))
	}

	account, err = i.GetForUpdate(ctx, account.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("no account")
	}

	account.Hold = account.Hold.Sub(size)
	account.Available = account.Available.Add(size)

	transaction := &Transaction{
		AccountId: account.ID,
		Available: size,
		Hold:      size.Neg(),
		Type:      transactionType.String(),
		Notes:     null.StringFrom(note),
	}
	err = TransactionDao.BatchCreate(ctx, []*Transaction{transaction})
	if err != nil {
		return err
	}

	err = i.Update(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

// DecreaseBalance
// 1. decrease user's balance in DB
func (i *accountDaoImpl) DecreaseBalance(ctx context.Context, account *Account, size decimal.Decimal, transactionType types.TransactionType, note string) error {
	if size.LessThanOrEqual(decimal.Zero) {
		return errors.New("DecreaseBalance size less than 0")
	}

	enough, err := i.hasEnoughBalance(ctx, account.ID, size)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if !enough {
		return errors.New(fmt.Sprintf("no enough balance for %v", size))
	}

	account, err = i.GetForUpdate(ctx, account.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("no account")
	}

	account.Available = account.Available.Sub(size)

	transaction := &Transaction{
		AccountId: account.ID,
		Available: size.Neg(),
		Hold:      decimal.Zero,
		Type:      transactionType.String(),
		Notes:     null.StringFrom(note),
	}
	err = TransactionDao.BatchCreate(ctx, []*Transaction{transaction})
	if err != nil {
		return err
	}

	err = i.Update(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func (i *accountDaoImpl) hasEnoughBalance(ctx context.Context, accountID uint, size decimal.Decimal) (bool, error) {
	account, err := i.Get(ctx, accountID)
	if err != nil {
		return false, err
	}
	if account == nil {
		return false, nil
	}
	return account.Available.GreaterThanOrEqual(size), nil
}

func (i *accountDaoImpl) hasEnoughHoldBalance(ctx context.Context, accountID uint, size decimal.Decimal) (bool, error) {
	account, err := i.Get(ctx, accountID)
	if err != nil {
		return false, err
	}
	if account == nil {
		return false, nil
	}
	return account.Hold.GreaterThanOrEqual(size), nil
}
