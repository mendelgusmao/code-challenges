package sessions

import (
	"log"

	"bitbucket.org/mendelgusmao/me_gu/frontend/config"
	"github.com/gorilla/sessions"
)

var store sessions.Store

func init() {
	config.AfterLoad(createSessionStore)
}

func createSessionStore(c *config.Specification) error {
	log.Println("creating session store with specified session key")
	store = sessions.NewCookieStore([]byte(c.SessionKey))

	return nil
}
