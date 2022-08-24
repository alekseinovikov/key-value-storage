package log

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresTransactionLogger struct {
	events chan<- Event
	errors <-chan error
	db     *sql.DB
}

type PostgresDBParams struct {
	dbName   string
	host     string
	user     string
	password string
}

func NewPostgresTransactionLogger(config PostgresDBParams) (TransactionLogger, error) {
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", config.host, config.dbName, config.user, config.password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	logger := &PostgresTransactionLogger{db: db}

}

func (p *PostgresTransactionLogger) WriteDelete(key string) {
	p.events <- Event{EventType: EventDelete, Key: key}
}

func (p *PostgresTransactionLogger) WritePut(key, value string) {
	p.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (p *PostgresTransactionLogger) Err() <-chan error {
	return p.errors
}

func (p *PostgresTransactionLogger) ReadEvents() (<-chan Event, <-chan error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresTransactionLogger) Run() {
	//TODO implement me
	panic("implement me")
}
