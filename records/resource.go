package records

// Resource describes a collection of records, such as comments or
// posts.
type Resource struct {
	ID             string           `json:"id"`
	RecordsPerPage int              `json:"records_per_page"`
	Fields         map[string]*Field `json:"fields"`
}

// Field is an attribute of a resource, similar to a column in a SQL database.
type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
	Required    bool   `json:"required"`
	Default     string `json:"default"`
}
