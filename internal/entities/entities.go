package entities

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
)

var Txn *mysql.Txn

// Initialize initializes all daos in this folder
func Initialize(dao *mysql.Dao) {
	Txn = mysql.NewTxn(dao)
	AccountDao = &accountDaoImpl{Dao: dao}
	AllowanceDao = &allowanceDaoImpl{Dao: dao}
	GatewayDao = &gatewayDaoImpl{Dao: dao}
	TransactionDao = &transactionDaoImpl{Dao: dao}
	DepositEventDao = &depositEventDaoImpl{Dao: dao}
	WithdrawEventDao = &withdrawEventDaoImpl{Dao: dao}
	ReceiptDao = &receiptDaoImpl{Dao: dao}
}
