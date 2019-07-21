package main

import (
	"testing"
)

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	got := getAllArticles()
	want := articleList

	// Check that the length of the list of articles returned is the same 
	// as the length of the global variable holding the list
	gotLen := len(got)
	wantLen := len(want)
	if gotLen != wantLen {
		t.Errorf("got %d length of list, wanted %d", gotLen, wantLen)
	}

	// Check that each member is identical
	for i, a := range got {
		if a.ID != want[i].ID {
			t.Errorf("got %d article ID, wanted %d", a.ID, want[i].ID)
		}
		if a.Title != articleList[i].Title {
			t.Errorf("got '%s' article title, wanted '%s'", a.Title, want[i].Title)
		}
		if a.Content != want[i].Content {
			t.Errorf("got '%s' article content, wanted '%s'", a.Content, want[i].Content)
		}
		
	}
}

// Next steps Allowing Users to Post New Articles