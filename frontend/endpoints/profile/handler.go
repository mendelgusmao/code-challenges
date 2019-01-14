package profile

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	resty "gopkg.in/resty.v1"
)

func handler(w http.ResponseWriter, r *http.Request, templateNames ...string) {
	session := context.Get(r, "session").(*sessions.Session)
	id := int(session.Values["id"].(float64))

	user, err := retrieveUser(id)

	if err != nil {
		log.Printf("profile: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	templateData := struct {
		User map[string]interface{}
	}{
		User: *user,
	}

	templates.NewRenderer(templateNames...).Do(w, r, templateData)
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
