package records

// ValidationErrors is the type returned from `ValidateRecord` that
// describes.
type ValidationErrors struct {
	Errors map[string]string
}

// ValidateRecord validates the given data against the resource configuration.
func ValidateRecord(r *Resource, data map[string]interface{}) ValidationErrors {
	validation := ValidationErrors{
		Errors: make(map[string]string),
	}

	for key, field := range r.Fields {
		if field.Required {
			if val, ok := data[key]; ok {
				switch v := val.(type) {
				case string:
					if v == "" {
						validation.addRequired(key)
					}
				}
			} else {
				validation.addRequired(key)
			}
		}
	}

	return validation
}

func (ve *ValidationErrors) addRequired(key string) {
	ve.Errors[key] = "required"
}

// IsValid returns true if the validation passed, false if there are any errors.
func (ve *ValidationErrors) IsValid() bool {
	return len(ve.Errors) == 0
}
