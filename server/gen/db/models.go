// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package gendb

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type UsersUserRole string

const (
	UsersUserRolePlayer UsersUserRole = "player"
	UsersUserRoleAdmin  UsersUserRole = "admin"
)

func (e *UsersUserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersUserRole(s)
	case string:
		*e = UsersUserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersUserRole: %T", src)
	}
	return nil
}

type NullUsersUserRole struct {
	UsersUserRole UsersUserRole
	Valid         bool // Valid is true if UsersUserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UsersUserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersUserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersUserRole), nil
}

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type BenchmarkLog struct {
	ID        int64
	ServerID  int64
	Score     sql.NullInt64
	LogOutput json.RawMessage
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServerInfo struct {
	ID            int64
	ServerAddress string
	TeamID        int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Team struct {
	ID        int64
	TeamName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID          int64
	UserID      string
	DisplayName string
	IconPath    string
	UserRole    UsersUserRole
	TeamID      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}