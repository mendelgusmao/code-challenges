package profile

import (
	"fmt"
	"net/http"

	resty "gopkg.in/resty.v1"
)

func retrieveUser(id int) (*map[string]interface{}, error) {
	uri := fmt.Sprintf("/users/%d", id)

	response, err := resty.R().
		SetResult(map[string]interface{}{}).
		Get(uri)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("profile.retrieveUser: backend returned %d", response.StatusCode())
	}

	return response.Result().(*map[string]interface{}), nil
}
