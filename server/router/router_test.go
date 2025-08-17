package router_test

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	rtr "cafe/router"
)

func makeTempStatic(t *testing.T, indexHTML string) string {
	t.Helper()
	dir := t.TempDir()
	// index.html
	if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte(indexHTML), fs.FileMode(0644)); err != nil {
		t.Fatalf("write index.html: %v", err)
	}
	// minimal assets dir to avoid 404s when referenced
	if err := os.MkdirAll(filepath.Join(dir, "assets"), 0755); err != nil {
		t.Fatalf("mkdir assets: %v", err)
	}
	return dir
}

func TestHealth(t *testing.T) {
	staticDir := makeTempStatic(t, "<!doctype html><title>ok</title>")
	r := rtr.NewRouter(staticDir)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if got := w.Body.String(); !strings.Contains(got, `"status"`) || !strings.Contains(got, `"ok"`) {
		t.Fatalf("unexpected body: %s", got)
	}
}

func TestServeIndexHTMLAtRoot(t *testing.T) {
	staticDir := makeTempStatic(t, "<!doctype html><h1>Cafe App</h1>")
	r := rtr.NewRouter(staticDir)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if got := w.Body.String(); !strings.Contains(got, "Cafe App") {
		t.Fatalf("index.html not served, body: %s", got)
	}
}

func TestSPAFallbackServesIndex(t *testing.T) {
	staticDir := makeTempStatic(t, "<!doctype html><h1>Fallback OK</h1>")
	r := rtr.NewRouter(staticDir)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/some/client/route", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if got := w.Body.String(); !strings.Contains(got, "Fallback OK") {
		t.Fatalf("fallback did not serve index.html, body: %s", got)
	}
}

func TestAPIUnknownIs404(t *testing.T) {
	staticDir := makeTempStatic(t, "<!doctype html>")
	r := rtr.NewRouter(staticDir)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/unknown", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for unknown API route, got %d", w.Code)
	}
}
