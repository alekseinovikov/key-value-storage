package main

import (
	"fmt"
	"key-value-storage/log"
)

var logger log.TransactionLogger

func InitializeTransactionLog() error {
	var err error
	logger, err = log.NewFileTransactionLogger("transaction.log")
	if err != nil {
		return fmt.Errorf("failed to create event logger: %w", err)
	}

	events, errors := logger.ReadEvents()
	e, ok := log.Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errors:
		case e, ok = <-events:
			switch e.EventType {
			case log.EventDelete:
				err = Delete(e.Key)
			case log.EventPut:
				err = Put(e.Key, e.Value)
			}

		}
	}

	logger.Run()
	return err
}
