package models

import (
	"time"
)

type BaseEvent struct {
	Event Event `json:",omitempty"`
}

type Event struct {
	Id           string     `json:",omitempty"`
	Log          EventLog   `json:",omitempty"`
	Created      *time.Time `json:",omitempty"`
	IsDelivered  bool       `json:",omitempty"`
	Subscription string     `json:",omitempty"`
	WorkspaceId  string     `json:",omitempty"`
}

type EventLog struct {
	Id      string     `json:",omitempty"`
	Invoice Invoice    `json:",omitempty"`
	Errors  []string   `json:",omitempty"`
	Type    string     `json:",omitempty"`
	Created *time.Time `json:",omitempty"`
}
