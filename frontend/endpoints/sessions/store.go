package sessions

import (
	"fmt"

	"bitbucket.org/mendelgusmao/me_gu/frontend/config"
	"github.com/gorilla/sessions"
)

var store sessions.Store

func init() {
	config.AfterLoad(createSessionStore)
}

func createSessionStore(c *config.Specification) error {
	fmt.Println("session key", c.SessionKey)
	store = sessions.NewCookieStore([]byte(c.SessionKey))

	return nil
}
