package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()

		if errorRollback != nil {
			panic(errorRollback)
		}

	} else {
		errorCommit := tx.Commit()

		if errorCommit != nil {
			panic(errorCommit)
		}
	}
}
