package records

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/luizpvas/blocks/db"
)

// Resource describes a collection of records, such as comments or
// posts.
type Resource struct {
	ID     string
	Fields map[string]Field
}

// Field is an attribute of a resource, similar to a column in a SQL database.
type Field struct {
	Type        string
	Label       string
	Placeholder string
	Required    bool
	Default     string
}

// CreateResult is the return type for the `CreateRecord` function. It describes
// all aspects of the operation, such as validation and record id.
type CreateResult struct {
	Success    bool
	Err        error
	Errors     ValidationErrors
	InsertedID string
}

// CreateRecord inserts a new record in the database for this resource.
func (r *Resource) CreateRecord(data map[string]interface{}) CreateResult {
	validation := ValidateRecord(r, data)
	if !validation.IsValid() {
		return CreateResult{
			Success: false,
			Errors:  validation,
			Err:     nil,
		}
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return CreateResult{
			Success: false,
			Err:     fmt.Errorf("failed to generate uuid for record: %v", err),
			Errors:  validation,
		}
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		return CreateResult{
			Success: false,
			Err:     fmt.Errorf("failed to encode record attributes: %v", err),
			Errors:  validation,
		}
	}

	_, err = db.DB.Exec(`insert into records(resource_id, record_id, attributes) values($1, $2, $3)`, r.ID, id.String(), encodedData)
	if err != nil {
		return CreateResult{
			Success: false,
			Err:     fmt.Errorf("failed to insert record in the database: %v", err),
			Errors:  validation,
		}
	}

	return CreateResult{
		Success:    true,
		Err:        nil,
		Errors:     validation,
		InsertedID: id.String(),
	}
}
