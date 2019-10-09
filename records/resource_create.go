package records

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/luizpvas/blocks/db"
)

// CreateResult is the return type for the `CreateRecord` function. It describes
// all aspects of the operation, such as validation and record id.
type CreateResult struct {
	Success    bool
	Err        error
	Errors     ValidationErrors
	InsertedID string
}

// CreateRecordFromJSON parses the given JSON and calls `CreateRecord` with the result map.
func (r *Resource) CreateRecordFromJSON(rawJSON []byte) CreateResult {
	var data map[string]interface{}
	err := json.Unmarshal(rawJSON, &data)
	if err != nil {
		return CreateResult{
			Err: fmt.Errorf("failed to parse record JSON: %v", err),
		}
	}

	return r.CreateRecord(data)
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

	_, err = db.DB.Exec(`
        insert into records(resource_id, record_id, attributes, created_at, updated_at)
        values($1, $2, $3, current_timestamp, current_timestamp)
    `, r.ID, id.String(), encodedData)
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
