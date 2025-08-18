package db

import (
	"context"
	"database/sql"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	var u User
	err = db.QueryRowContext(ctx, `
		SELECT id, email, password_hash, role
		FROM users WHERE email = ? LIMIT 1
	`, email).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateCustomer(ctx context.Context, email, password string) (*User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	res, err := db.ExecContext(ctx, `
		INSERT INTO users (email, password_hash, role)
		VALUES (?, ?, 'customer')`, email, string(hash))
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &User{ID: id, Email: email, Role: "customer"}, nil
}

// seedAdmin ensures at least one admin exists (env-overridable credentials).
func seedAdmin(db *sql.DB) error {
	var exists int
	if err := db.QueryRow(`SELECT COUNT(*) FROM users WHERE role='admin'`).Scan(&exists); err != nil {
		return err
	}
	if exists > 0 {
		return nil
	}
	email := os.Getenv("ADMIN_EMAIL")
	if email == "" {
		email = "admin@example.com"
	}
	pass := os.Getenv("ADMIN_PASSWORD")
	if pass == "" {
		pass = "admin123"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		INSERT INTO users (email, password_hash, role)
		VALUES (?, ?, 'admin')`, email, string(hash))
	return err
}
