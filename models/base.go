package models

/*
Pagination
| Param      | Value                                                                                                         |
|------------|---------------------------------------------------------------------------------------------------------------|
| start_date | integer (optional)  Starting date of the data. Example: "2020-01-01".                                         |
| end_date   | integer (optional)  Ending date of the data. Example: "2020-12-12". End date must be greater than start_date. |
| pagination | integer (optional)  Pagination of the data.                                                                   |
| page       | integer (optional)  The desired page of the data pagination.                                                  |
| sort_by    | string (optional)  Sort the result by the attribute.                                                          |
| sort_type  | string (optional)  You can sort in sort_desc (descending)sort_asc (ascending)                                 |

Example Pagination:
[URL]?start_date=start_date&end_date=end_date&pagination=pagination&page=page&sort_by=sort_by&sort_type=sort_type
*/
type Pagination struct {
	StartDate  string `json:"start_date,omitempty"`
	EndDate    string `json:"end_date,omitempty"`
	Pagination string `json:"pagination,omitempty"`
	Page       string `json:"page,omitempty"`
	SortBy     string `json:"sort_by,omitempty"`
	SortType   string `json:"sort_type,omitempty"`
}
