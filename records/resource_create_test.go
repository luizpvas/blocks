package records

import (
	"testing"

	"github.com/luizpvas/blocks/db"
)

func TestRecordsCreateSuccess(t *testing.T) {
	clean := db.SetTestingConnection(t)
	defer clean()

	posts := PostsResource()

	result := posts.CreateRecord(map[string]interface{}{
		"title": "Awesome Post",
	})

	if !result.Success {
		t.Fatalf("failed to insert record in the database: %v", result.Err)
	}
}

func TestRecordsCreateValidationError(t *testing.T) {
	clean := db.SetTestingConnection(t)
	defer clean()

	posts := PostsResource()

	result := posts.CreateRecord(map[string]interface{}{
		"title": "",
	})

	if result.Success {
		t.Fatalf("create operation should have failed with validation")
	}

	if result.Errors.Errors["title"] != "required" {
		t.Fatalf("expected title to be required")
	}
}

func PostsResource() Resource {
	return Resource{
		ID: "posts",
		Fields: map[string]Field{
			"title": Field{
				Required: true,
			},
		},
	}
}
