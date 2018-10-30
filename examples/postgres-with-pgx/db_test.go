package mydb

import (
	"context"
	"testing"
)

func TestDB_SelectOne(t *testing.T) {
	db := DB{Conn: testConnPool}
	// initDBForTests(t, db)

	one, err := db.SelectOne(context.TODO())
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
	if one != 1 {
		t.Fatalf("Expected 1, got %v", one)
	}
}
