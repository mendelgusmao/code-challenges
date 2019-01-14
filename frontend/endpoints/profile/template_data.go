package profile

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	resty "gopkg.in/resty.v1"
)

type templateData struct {
	User   map[string]interface{}
	Action string
}

func buildTemplateData(w http.ResponseWriter, r *http.Request, formAction string) *templateData {
	session := context.Get(r, "session").(*sessions.Session)
	id := int(session.Values["id"].(float64))

	user, err := retrieveUser(id)

	if err != nil {
		log.Printf("profile: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return nil
	}

	return &templateData{
		User:   *user,
		Action: formAction,
	}
}

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
