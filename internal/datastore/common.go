package datastore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

type Timestamp time.Time

type PgStore struct {
	DB *sql.DB

	retrieveEvent *sql.Stmt
}

type NullString struct {
	sql.NullString
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

// NewPostgres initialises a connection to the DB - it takes in a connection string
// and returns a pointer to a PgStore{} struct (defined above)
func NewPostgres(pgConfig string) (*PgStore, error) {
	// parse connection string / env vars
	cfg, err := pgx.ParseConnectionString(strings.TrimSpace(pgConfig))
	if err != nil {
		return nil, err
	}

	// enable TCP keepalives
	dialer := &net.Dialer{
		KeepAlive: 5 * time.Second,
	}
	cfg.Dial = dialer.Dial

	// initialise the DB object (doesn't actually connect at this point)
	db := stdlib.OpenDB(cfg)
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(1)

	s := &PgStore{
		DB: db,
	}

	if err := s.prepare(); err != nil {
		_ = s.DB.Close()
		return nil, fmt.Errorf("failed to prepare statement :: %v", err)
	}

	return s, nil
}

// Test and initialise all the queries you're going to be using and
// throw an error if they don't work
func (s *PgStore) prepare() error {
	var err error
	// this query could use bindings from config in the future for the various interval states
	s.retrieveEvent, err = s.DB.Prepare(`SELECT * FROM event`)
	if err != nil {
		return fmt.Errorf("retrieveJob :: %v", err)
	}

	return nil
}
