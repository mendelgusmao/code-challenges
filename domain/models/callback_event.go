package models

import (
	"fmt"
	"io"
	"net/http"
)

type CallbackEvent struct {
	Data      string
	Signature string
}

func NewCallbackEvent(data io.ReadCloser, headers http.Header) (*CallbackEvent, error) {
	content, err := io.ReadAll(data)

	if err != nil {
		return nil, fmt.Errorf("reading event data: %v", err)
	}

	digitalSignature := headers["Digital-Signature"]

	if len(digitalSignature) == 0 {
		return nil, fmt.Errorf("digital signature not present in header")
	}

	return &CallbackEvent{
		Data:      string(content),
		Signature: digitalSignature[0],
	}, nil
}
