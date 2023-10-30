package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type ClientError struct {
	DisplayMessage string `json:"message"`
	HTTPCode       int    `json:"-"`
}

func (ce *ClientError) Error() string {
	return fmt.Sprintf("an error which will be displayed to the client: %d %v", ce.HTTPCode, ce.DisplayMessage)
}

var ErrNoFileProvided = &ClientError{
	DisplayMessage: "no file provided in multipart form with name 'file'",
	HTTPCode:       http.StatusBadRequest,
}

var ErrInvalidContentType = &ClientError{
	DisplayMessage: "request content type must be 'multipart/form-data'",
	HTTPCode:       http.StatusBadRequest,
}

func ErrTooBigFile(maxSizeMB int64) *ClientError {
	return &ClientError{
		DisplayMessage: fmt.Sprintf("file is too big. Max file size is %dMB", maxSizeMB),
		HTTPCode:       http.StatusRequestEntityTooLarge,
	}
}

func WriteErrorResponse(w http.ResponseWriter, e error) {
	if unwrapped := errors.Unwrap(e); unwrapped != nil {
		e = unwrapped
	}
	if ce, ok := e.(*ClientError); ok {
		w.WriteHeader(ce.HTTPCode)
		json.NewEncoder(w).Encode(ce)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("ERROR: %v", e.Error())
}
