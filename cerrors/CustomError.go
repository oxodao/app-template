package cerrors

import (
	"encoding/json"
	"net/http"
)

type CError struct {
	Message   string `json:"message"`
	Errcode   int    `json:"errcode"` // Errcode is a app-specific / internal error code to be handled by the client
	HttpError int    `json:"-"`
}

func (c *CError) Error() string {
	return c.Message
}

func (c *CError) Json() []byte {
	tx, err := json.Marshal(c)
	if err != nil {
		return []byte(`{ "message": "Something went wrong encoding the error!", "errcode": -1 }"`)
	}

	return tx
}

func (c *CError) Write(w http.ResponseWriter) {
	w.Header().Del("Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c.HttpError)
	w.Write(c.Json())
}


func WriteError(w http.ResponseWriter, err error) bool {
	if err != nil {
		cast, ok := err.(*CError)
		if ok {
			cast.Write(w)
			return true
		}

		newUnknown(err).Write(w)
		return true
	}

	return false
}

func newUnknown(err error) *CError {
	return &CError{
		Message:   err.Error(),
		Errcode:   -1,
		HttpError: http.StatusInternalServerError,
	}
}

func New(msg string, errCode, err int) *CError {
	return &CError{
		Message:   msg,
		Errcode:   errCode,
		HttpError: err,
	}
}
