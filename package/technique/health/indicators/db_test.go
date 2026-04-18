package indicators

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"
	"testing"
)

// testDriver is a minimal sql.Driver for testing ping behavior.
type testDriver struct {
	pingErr error
}

type testConn struct {
	pingErr error
}

func (c *testConn) Prepare(query string) (driver.Stmt, error) { return nil, nil }
func (c *testConn) Close() error                             { return nil }
func (c *testConn) Begin() (driver.Tx, error)                { return nil, nil }

// Implement driver.Pinger so sql.DB.PingContext calls our mock.
func (c *testConn) Ping(ctx context.Context) error { return c.pingErr }

func (d *testDriver) Open(_ string) (driver.Conn, error) {
	return &testConn{pingErr: d.pingErr}, nil
}

func registerDriver(name string, pingErr error) {
	// Re-register only if not already registered.
	for _, existing := range sql.Drivers() {
		if existing == name {
			return
		}
	}
	sql.Register(name, &testDriver{pingErr: pingErr})
}

func newTestDB(t *testing.T, driverName string) *sql.DB {
	t.Helper()
	db, err := sql.Open(driverName, "")
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func init() {
	registerDriver("test-ok", nil)
	registerDriver("test-fail", io.ErrUnexpectedEOF)
}

func TestDBIndicator_Up(t *testing.T) {
	db := newTestDB(t, "test-ok")
	ind := NewDBIndicator("postgres", db)

	result := ind.Check(context.Background())

	if !result.Up {
		t.Errorf("Up = false, want true; error = %q", result.Error)
	}
	if result.Key != "postgres" {
		t.Errorf("Key = %q, want %q", result.Key, "postgres")
	}
}

func TestDBIndicator_Down(t *testing.T) {
	db := newTestDB(t, "test-fail")
	ind := NewDBIndicator("postgres", db)

	result := ind.Check(context.Background())

	if result.Up {
		t.Error("Up = true, want false")
	}
	if result.Error == "" {
		t.Error("Error should be set when ping fails")
	}
}

func TestDBIndicator_CustomKey(t *testing.T) {
	db := newTestDB(t, "test-ok")
	ind := NewDBIndicator("primary-db", db)

	result := ind.Check(context.Background())

	if result.Key != "primary-db" {
		t.Errorf("Key = %q, want %q", result.Key, "primary-db")
	}
}
