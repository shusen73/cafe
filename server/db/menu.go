package db

import "sync"

type MenuItem struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}

type Store struct {
	mu   sync.RWMutex
	menu []MenuItem
}

func New() *Store {
	return &Store{
		menu: make([]MenuItem, 0, 16),
	}
}

func (s *Store) ListMenu() []MenuItem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]MenuItem, len(s.menu))
	copy(out, s.menu)
	return out
}

// SeedDefault populates a few items for development.
func (s *Store) SeedDefault() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.menu) > 0 {
		return
	}
	s.menu = []MenuItem{
		{ID: 1, Name: "Espresso", Description: "Rich single shot", Price: 2.50, Available: true},
		{ID: 2, Name: "Americano", Description: "Espresso + hot water", Price: 3.00, Available: true},
		{ID: 3, Name: "Latte", Description: "Espresso + steamed milk", Price: 3.80, Available: true},
		{ID: 4, Name: "Cappuccino", Description: "Foamy classic", Price: 3.80, Available: true},
		{ID: 5, Name: "Croissant", Description: "Butter pastry", Price: 2.20, Available: true},
	}
}
