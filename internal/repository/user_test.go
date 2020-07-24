package repository

import (
	"database/sql"
	"testing"
)

func TestMySQL_Connect(t *testing.T) {
	Repository = mockedRepository
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MySQL{
				db: tt.fields.db,
			}
		})
	}
}

func TestMySQL_Fetch(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MySQL{
				db: tt.fields.db,
			}
		})
	}
}

func TestMySQL_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MySQL{
				db: tt.fields.db,
			}
		})
	}
}

func TestMySQL_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MySQL{
				db: tt.fields.db,
			}
		})
	}
}

func TestMySQL_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MySQL{
				db: tt.fields.db,
			}
		})
	}
}