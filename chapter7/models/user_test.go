package models

import (
	"database/sql"
	"testing"

	"github.com/pkg/errors"
)

func TestGetUserByUsername(t *testing.T) {
	db := &MockDB{
		mockQuery: func(query string, args ...interface{}) (*sql.Rows, error) {
			return nil, errors.New("test query failure!")
		},
	}

	_, err := GetUserByUsername(db, "test")
	if err != nil {
		if errors.Cause(err).Error() != "test query failure!" {
			t.Errorf("incorrect failure expected: %s", err.Error())
		}
	}
}
