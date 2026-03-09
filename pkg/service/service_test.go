package service

import (
	"bytes"
	"strings"
	"testing"
)

func TestGenerateSQLAcceptsUppercaseOperation(t *testing.T) {
	input := strings.NewReader("yes\nSELECT\nusers\n")

	sql, err := generateSQL([]string{"111", "222"}, input, &bytes.Buffer{})
	if err != nil {
		t.Fatalf("generateSQL returned error: %v", err)
	}

	want := "SELECT * FROM users WHERE user_id IN (111, 222);"
	if sql != want {
		t.Fatalf("unexpected SQL:\nwant: %s\ngot:  %s", want, sql)
	}
}

func TestGenerateSQLReadsFullUpdateClause(t *testing.T) {
	input := strings.NewReader("yes\nupdate\nusers\nset status = 'active'\n")

	sql, err := generateSQL([]string{"111", "222"}, input, &bytes.Buffer{})
	if err != nil {
		t.Fatalf("generateSQL returned error: %v", err)
	}

	want := "UPDATE users set status = 'active' WHERE user_id IN (111, 222);"
	if sql != want {
		t.Fatalf("unexpected SQL:\nwant: %s\ngot:  %s", want, sql)
	}
}

func TestGenerateSQLErrorsWhenIDsAreEmpty(t *testing.T) {
	_, err := generateSQL(nil, strings.NewReader("yes\n"), &bytes.Buffer{})
	if err == nil {
		t.Fatal("expected error for empty IDs, got nil")
	}
}
