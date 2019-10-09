package records

import (
	"fmt"
	"math"

	"github.com/luizpvas/blocks/db"
)

// IndexResult is the return type from the `ListRecords` function.
type IndexResult struct {
	Err         error
	Records     []Record
	TotalPages  int
	CurrentPage int
}

// IndexParams describes the paramaters we need to run the `ListRecords` function.
type IndexParams struct {
	Page int
}

// Index lists records from the database according to some filters.
func (r *Resource) Index(params IndexParams) IndexResult {
	skip := (params.Page - 1) * r.RecordsPerPage
	if skip < 0 {
		skip = 0
	}

	total, err := r.countTotalRecords(params)
	if err != nil {
		return IndexResult{
			Err: fmt.Errorf("could not count total records for pagination: %v", err),
		}
	}

	rows, err := db.DB.Query(`
        select resource_id, record_id, attributes, created_at, updated_at
        from records
        where resource_id = $1
        order by created_at desc
        offset $2
        limit $3
    `, r.ID, skip, r.RecordsPerPage)

	if err != nil {
		return IndexResult{
			Err: err,
		}
	}
	defer rows.Close()

	records := make([]Record, 0)
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.ResourceID, &record.RecordID, &record.Attributes, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return IndexResult{
				Err: fmt.Errorf("failed to scan database columns: %v", err),
			}
		}
		records = append(records, record)
	}

	return IndexResult{
		Err:         nil,
		Records:     records,
		CurrentPage: params.Page,
		TotalPages:  total,
	}
}

func (r *Resource) countTotalRecords(params IndexParams) (int, error) {
	row := db.DB.QueryRow(`
        select count(*) from records
        where resource_id = $1
    `, r.ID)

	var totalRecords float64
	err := row.Scan(&totalRecords)

	totalPages := math.Ceil(float64(totalRecords / float64(r.RecordsPerPage)))

	return int(totalPages), err
}
