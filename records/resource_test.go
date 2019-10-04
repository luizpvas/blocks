package records

import (
	"testing"

	"github.com/luizpvas/blocks/db"
)

func TestCreateRecord(t *testing.T) {
	clean := db.SetTestingConnection(t)
	defer clean()

	posts := Resource{
		ID: "posts",
		Fields: map[string]Field{
			"title": Field{
				Required: true,
			},
		},
	}

	// =====================================
	// Success case
	// =====================================
	result := posts.CreateRecord(map[string]interface{}{
		"title": "Awesome Post",
	})
	if !result.Success {
		t.Fatalf("failed to insert record in the database: %v", result.Err)
	}

	// =====================================
	// Failure with validation error
	// =====================================
	result = posts.CreateRecord(map[string]interface{}{
		"title": "",
	})
	if result.Success {
		t.Fatalf("create operation should have failed with validation")
	}
	if result.Errors.Errors["title"] != "required" {
		t.Fatalf("expected title to be required")
	}
}
