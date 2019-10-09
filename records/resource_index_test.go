package records

import (
	"testing"

	"github.com/luizpvas/blocks/db"
)

func TestListRecordsWithNoFilter(t *testing.T) {
	clean := db.SetTestingConnection(t)
	defer clean()

	posts := Resource{
		ID:             "posts",
		RecordsPerPage: 2,
		Fields: map[string]Field{
			"title": Field{
				Required: true,
			},
		},
	}

	posts.CreateRecord(map[string]interface{}{
		"title": "Awesome Post 1",
	})
	posts.CreateRecord(map[string]interface{}{
		"title": "Awesome Post 2",
	})
	posts.CreateRecord(map[string]interface{}{
		"title": "Awesome Post 3",
	})

	// =============================
	// Fetching the first page
	// =============================
	result := posts.Index(IndexParams{
		Page: 1,
	})
	if len(result.Records) != 2 {
		t.Fatalf("failed to fetch 2 posts per page - got %v results (%v)", len(result.Records), result.Err)
	}
	if result.TotalPages != 2 {
		t.Fatalf("failed to calculate total amount of pages - expected 2, got %v", result.TotalPages)
	}

	// =============================
	// Fetching the second page
	// =============================
	result = posts.Index(IndexParams{
		Page: 2,
	})
	if len(result.Records) != 1 {
		t.Fatalf("failed to fetch 1 post in the second page - got %v results (%v)", len(result.Records), result.Err)
	}
}
