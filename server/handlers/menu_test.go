package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"cafe/db"
	rtr "cafe/router"
)

func TestMenuList(t *testing.T) {
	// Ensure default store is seeded.
	db.Default().SeedDefault()

	// router needs a static dir; we don't exercise static here.
	r := rtr.NewRouter(t.TempDir())

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/menu", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	body := w.Body.String()
	wantSnippets := []string{`"items"`, `"name"`, `"price"`, `"available"`}
	for _, s := range wantSnippets {
		if !contains(body, s) {
			t.Fatalf("response missing %q. got: %s", s, body)
		}
	}
}

func contains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})())
}
