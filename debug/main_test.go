package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req, e := http.NewRequest(http.MethodGet, "/d@golang.org", nil)
	if e != nil {
		t.Error(e)
	}
	rec := httptest.NewRecorder()
	handler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected %v got %v\n", http.StatusOK, rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "/d@golang.org") {
		t.Errorf("unepected body %q", rec.Body.String())
	}
}

func TestHandler01(t *testing.T) {
	req, e := http.NewRequest(http.MethodGet, "/", nil)
	if e != nil {
		t.Error(e)
	}
	rec := httptest.NewRecorder()
	handler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected %v got %v\n", http.StatusOK, rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Hello from debug") {
		t.Errorf("unepected body %q", rec.Body.String())
	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, e := http.NewRequest(http.MethodGet, "/d@golang.org", nil)
		if e != nil {
			b.Error(e)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)
		if rec.Code != http.StatusOK {
			b.Errorf("expected %v got %v\n", http.StatusOK, rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "/d@golang.org") {
			b.Errorf("unepected body %q", rec.Body.String())
		}
	}
}
