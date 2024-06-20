package data

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
)

func checkErr(err error) {
	if err != nil {
		if mysql.GetGormErrorCode(err) == mysql.GormErrorCodeDuplicateEntry {
			log.Info(err).Msgf("")
		} else {
			panic(err)
		}
	}
}

func logErr(err error) {
	if err != nil {
		log.Error().Err(err)
	}
}
