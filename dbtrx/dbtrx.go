package dbtrx

import "database/sql"

func Trx(tx *sql.Tx, errorOccurred error) {
	if tx == nil {
		return
	}
	if errorOccurred != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
