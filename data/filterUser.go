package data

type FilterUserRequest struct {
	Age int32 `json:"age" query:"age"`
}
