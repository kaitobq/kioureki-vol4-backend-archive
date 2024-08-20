package query

import (
	"backend/domain/errors"
	"database/sql"
	"fmt"
	"os"
	"sync"
)

func executeSQLFile(db *sql.DB, filePath string) error {
	query, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("%v(path: %s) : %v", errors.ErrReadFile, filePath, err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		return fmt.Errorf("%v(path: %s) : %v", errors.ErrExecFile, filePath, err)
	}

	return nil
}

func ExecuteSQLFiles(db *sql.DB, filePaths []string) []error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(filePaths))
	for _, filePath := range filePaths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			err := executeSQLFile(db, path)
			if err != nil {
				errChan <- err
			}
		}(filePath)
	}
	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	return errs
}