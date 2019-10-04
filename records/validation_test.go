package records

import "testing"

func TestValidationRequired(t *testing.T) {
	resource := Resource{
		Fields: map[string]Field{
			"title": Field{
				Required: true,
			},
			"body": Field{
				Required: false,
			},
		},
	}

	// ==============================
	// case: record is valid
	// ==============================
	validation := ValidateRecord(&resource, map[string]interface{}{
		"title": "My post",
		"body":  "My body",
	})
	if !validation.IsValid() {
		t.Fatalf("record with title and body should be valid, but it's not: %v", validation)
	}

	// ===============================================================
	// case: record is invalid - key is present but value is empty
	// ===============================================================
	validation = ValidateRecord(&resource, map[string]interface{}{
		"title": "",
		"body":  "My body",
	})
	if validation.IsValid() {
		t.Fatalf("record without title should not be valid")
	}
	if validation.Errors["title"] != "required" {
		t.Fatalf("record title should be `required`, got: %v", validation.Errors["title"])
	}

	// ===================================================
	// case: record is invalid - key is missing
	// ===================================================
	validation = ValidateRecord(&resource, map[string]interface{}{
		"body": "My body",
	})
	if validation.IsValid() {
		t.Fatalf("record without title should not be valid")
	}
	if validation.Errors["title"] != "required" {
		t.Fatalf("record title should be `required`, got: %v", validation.Errors["title"])
	}
}
