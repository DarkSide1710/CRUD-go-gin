package view

type R struct {
	Status    string      `json:"status"`     // success
	ErrorCode int         `json:"error_code"` // error_code
	ErrorNote string      `json:"error_note"` // error_note
	Data      interface{} `json:"data"`       // data
}
