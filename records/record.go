package records

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Record is like a row in a database table.
type Record struct {
	ResourceID string
	RecordID   string
	Attributes Attributes
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Attributes is the custom generic type stored as a `jsonb` in postgres.
type Attributes map[string]interface{}

// Value encodes the value to store it in the database.
func (a Attributes) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan decodes the value stored in the database.
func (a *Attributes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion on []byte failed")
	}
	return json.Unmarshal(b, &a)
}
