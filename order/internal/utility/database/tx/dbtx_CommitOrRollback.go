package tx

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func (r *DbTx) CommitOrRollback(ctx context.Context, err error) error {
	tx, ok := r.txFromContext(ctx)
	if !ok {
		return err
	}

	if err != nil {
		log.Println("[DB-TX] rolling back transaction")
		errRb := tx.Rollback()
		errRb = r.checkErrTxDone(errRb)
		if errRb != nil {
			log.Printf("[DB-TX] failed to roll back transaction: %s", errRb.Error())
		}
		log.Println("[DB-TX] transaction rolled back")
		return err
	}

	log.Println("[DB-TX] committing transaction")
	err = tx.Commit()
	err = r.checkErrTxDone(err)
	if err != nil {
		log.Printf("[DB-TX] failed to commit transaction: %s", err.Error())
	}
	log.Println("[DB-TX] transaction committed")

	return err
}

func (r *DbTx) checkErrTxDone(err error) error {
	if errors.Is(err, sql.ErrTxDone) {
		log.Println("[DB-TX] transaction has already been committed or rolled back")
		return nil
	}

	return err
}
