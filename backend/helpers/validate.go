package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const emailAllowedRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.@_-+"
const passwordAllowedRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}:,.?"

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be null")
	}

	if len(email) > 200 {
		return fmt.Errorf("email too long")
	}

	if len(email) < 13 {
		return fmt.Errorf("email too short")
	}

	for _, r := range email {
		if r > 127 || !strings.ContainsRune(emailAllowedRunes, r) {
			return fmt.Errorf("email contains invalid character: %q", r)
		}
	}

	if strings.Count(email, "@") != 1 {
		return fmt.Errorf("email must contain exactly one @")
	}

	parts := strings.Split(email, "@")
	local, domain := parts[0], parts[1]

	if local == "" {
		return fmt.Errorf("email local part is empty")
	}

	if domain == "" {
		return fmt.Errorf("email domain is empty")
	}

	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return fmt.Errorf("email domain cannot start or end with dot")
	}

	if !strings.Contains(domain, ".") {
		return fmt.Errorf("email domain must contain a dot")
	}

	return nil
}

func EmailExists(ctx context.Context, db *sql.DB, email string) (bool, error) {
	var exists bool

	err := db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password cannot be null")
	}

	if len(password) < 8 {
		return fmt.Errorf("password too short")
	}

	if len(password) > 64 {
		return fmt.Errorf("password too long")
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, r := range password {
		if r > 127 || !strings.ContainsRune(passwordAllowedRunes, r) {
			return fmt.Errorf("password contains invalid character: %q", r)
		}

		switch {
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		case r >= '0' && r <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
	}

	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	if !hasDigit {
		return fmt.Errorf("password must contain at least one digit")
	}

	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}
