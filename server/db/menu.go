package db

import (
	"context"
	"database/sql"
)

type MenuItem struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PriceCents  int64  `json:"priceCents"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

func GetAllMenuItems(ctx context.Context) ([]MenuItem, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, `
		SELECT id, name, price_cents, description, active
		FROM menu_items
		ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []MenuItem
	for rows.Next() {
		var m MenuItem
		var activeInt int
		if err := rows.Scan(&m.ID, &m.Name, &m.PriceCents, &m.Description, &activeInt); err != nil {
			return nil, err
		}
		m.Active = activeInt == 1
		items = append(items, m)
	}
	return items, rows.Err()
}

func CreateMenuItem(ctx context.Context, m *MenuItem) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	res, err := db.ExecContext(ctx, `
		INSERT INTO menu_items (name, price_cents, description, active)
		VALUES (?, ?, ?, ?)
	`, m.Name, m.PriceCents, m.Description, boolToInt(m.Active))
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = id
	return nil
}

func UpdateMenuItem(ctx context.Context, m *MenuItem) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, `
		UPDATE menu_items
		SET name = ?, price_cents = ?, description = ?, active = ?
		WHERE id = ?
	`, m.Name, m.PriceCents, m.Description, boolToInt(m.Active), m.ID)
	return err
}

func DeleteMenuItem(ctx context.Context, id int64) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, `DELETE FROM menu_items WHERE id = ?`, id)
	return err
}

func seedIfEmpty(db *sql.DB) error {
	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM menu_items`).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	_, err := db.Exec(`
		INSERT INTO menu_items (name, price_cents, description, active) VALUES
			('Espresso', 300, 'Rich single shot', 1),
			('Latte', 450, 'Velvety milk & espresso', 1),
			('Croissant', 350, 'Buttery, flaky pastry', 1)
	`)
	return err
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
