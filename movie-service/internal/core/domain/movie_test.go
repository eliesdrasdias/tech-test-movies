package domain

import "testing"

func TestMovie(t *testing.T) {
	m := Movie{ID: 1, Title: "Test Movie", Year: "2022"}
	if m.ID != 1 {
		t.Errorf("Expected ID 1, got %d", m.ID)
	}
	if m.Title != "Test Movie" {
		t.Errorf("Expected title 'Test Movie', got '%s'", m.Title)
	}
	if m.Year != "2022" {
		t.Errorf("Expected year '2022', got '%s'", m.Year)
	}
}
