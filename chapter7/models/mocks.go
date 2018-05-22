package models

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

type MockableDB interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
	Conn(ctx context.Context) (*sql.Conn, error)
	Driver() driver.Driver
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Ping() error
	PingContext(ctx context.Context) error
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Stats() sql.DBStats
}

type MockDB struct {
	mockBegin              func() (*sql.Tx, error)
	mockBeginTx            func(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	mockClose              func() error
	mockConn               func(ctx context.Context) (*sql.Conn, error)
	mockDriver             func() driver.Driver
	mockExec               func(query string, args ...interface{}) (sql.Result, error)
	mockExecContext        func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	mockPing               func() error
	mockPingContext        func(ctx context.Context) error
	mockPrepare            func(query string) (*sql.Stmt, error)
	mockPrepareContext     func(ctx context.Context, query string) (*sql.Stmt, error)
	mockQuery              func(query string, args ...interface{}) (*sql.Rows, error)
	mockQueryContext       func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	mockQueryRow           func(query string, args ...interface{}) *sql.Row
	mockQueryRowContext    func(ctx context.Context, query string, args ...interface{}) *sql.Row
	mockSetConnMaxLifetime func(d time.Duration)
	mockSetMaxIdleConns    func(n int)
	mockSetMaxOpenConns    func(n int)
	mockStats              func() sql.DBStats
}

func (db *MockDB) Begin() (*sql.Tx, error) {
	if db.mockBegin != nil {
		return db.mockBegin()
	}
	return nil, nil
}
func (db *MockDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	if db.mockBeginTx != nil {
		return db.mockBeginTx(ctx, opts)
	}
	return nil, nil
}
func (db *MockDB) Close() error {
	if db.mockClose != nil {
		return db.mockClose()
	}
	return nil
}
func (db *MockDB) Conn(ctx context.Context) (*sql.Conn, error) {
	if db.mockConn != nil {
		return db.mockConn(ctx)
	}
	return nil, nil
}
func (db *MockDB) Driver() driver.Driver {
	if db.mockDriver != nil {
		return db.mockDriver()
	}
	return nil
}
func (db *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if db.mockExec != nil {
		return db.mockExec(query, args...)
	}
	return nil, nil

}
func (db *MockDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if db.mockExecContext != nil {
		return db.mockExecContext(ctx, query, args...)
	}
	return nil, nil

}
func (db *MockDB) Ping() error {
	if db.mockPing != nil {
		return db.mockPing()
	}
	return nil
}
func (db *MockDB) PingContext(ctx context.Context) error {
	if db.mockPingContext != nil {
		return db.mockPingContext(ctx)
	}
	return nil

}
func (db *MockDB) Prepare(query string) (*sql.Stmt, error) {
	if db.mockPrepare != nil {
		return db.mockPrepare(query)
	}
	return nil, nil

}
func (db *MockDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if db.mockPrepareContext != nil {
		return db.mockPrepareContext(ctx, query)
	}
	return nil, nil

}
func (db *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db.mockQuery != nil {
		return db.mockQuery(query, args...)
	}
	return nil, nil

}
func (db *MockDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if db.mockQueryContext != nil {
		return db.mockQueryContext(ctx, query, args...)
	}
	return nil, nil
}
func (db *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	if db.mockQueryRow != nil {
		return db.mockQueryRow(query, args...)
	}
	return nil
}
func (db *MockDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if db.mockQueryRowContext != nil {
		return db.mockQueryRowContext(ctx, query, args...)
	}
	return nil
}
func (db *MockDB) SetConnMaxLifetime(d time.Duration) {
	if db.mockSetConnMaxLifetime != nil {
		db.mockSetConnMaxLifetime(d)
	}
}
func (db *MockDB) SetMaxIdleConns(n int) {
	if db.mockSetMaxIdleConns != nil {
		db.mockSetMaxIdleConns(n)
	}
}
func (db *MockDB) SetMaxOpenConns(n int) {
	if db.mockSetMaxOpenConns != nil {
		db.mockSetMaxOpenConns(n)
	}

}
func (db *MockDB) Stats() sql.DBStats {
	if db.mockStats != nil {
		return db.mockStats()
	}
	return sql.DBStats{}
}
