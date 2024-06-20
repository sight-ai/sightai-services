package mysql

import (
	"database/sql/driver"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/jinzhu/gorm"
	"reflect"
	"regexp"
	"time"
	"unicode"
)

const (
	timeFormat = "2006-01-02T15:04:05-07:00"
)

var (
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
	sqlRegexp                = regexp.MustCompile(`\?`)
)

// MySQLLogger return logger for MySQL DB
type MySQLLogger gorm.Logger

// Print implements the gorm logger
func (o MySQLLogger) Print(args ...interface{}) {
	log.Debug().Msgf("%v %vsql=%v", MySQLLogFormatterDetail(args...)...)
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// MySQLLogFormatter defines the format of the SQL log
var MySQLLogFormatter = func(values ...interface{}) (messages []interface{}) {
	// TODO: send stats to datadog
	if len(values) >= 4 {
		var (
			level   = values[0]
			latency = values[2]
			sql     = values[3]
		)
		if level == "sql" {
			messages = []interface{}{"sql: "}
			messages = append(messages, fmt.Sprintf("latency=%.2fms, ", float64(latency.(time.Duration).Nanoseconds()/1e4)/100.0))
			messages = append(messages, fmt.Sprintf("\"%v\", ", sql))
		} else {
			messages = append(messages, values[1:]...)
		}
	} else {
		messages = append(messages, values[2:]...)
	}
	return
}

// MySQLLogFormatterDetail defines the detail format of the SQL log
var MySQLLogFormatterDetail = func(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
		)

		messages = []interface{}{"mysql query,"}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf("latency=%.2fms, ", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.UTC().Format(timeFormat)))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value)
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, fmt.Sprintf("\"%v\"", sql))
		} else {
			messages = append(messages, values[2:]...)
		}
	}

	return
}
