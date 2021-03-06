// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "firstgoa": Application User Types
//
// Command:
// $ goagen
// --design=github.com/phxuecan/FirstGo/firstgoa/design
// --out=$(GOPATH)/src/github.com/phxuecan/FirstGo/firstgoa
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"time"
)

// message of user to send
type message struct {
	Info *struct {
		// content of message
		Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
		// user send message time
		SendTime *time.Time `form:"sendTime,omitempty" json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	} `form:"info,omitempty" json:"info,omitempty" xml:"info,omitempty"`
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// Validate validates the message type instance.
func (ut *message) Validate() (err error) {
	if ut.Info == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "info"))
	}
	return
}

// Publicize creates Message from message
func (ut *message) Publicize() *Message {
	var pub Message
	if ut.Info != nil {
		pub.Info = &struct {
			// content of message
			Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
			// user send message time
			SendTime *time.Time `form:"sendTime,omitempty" json:"sendTime,omitempty" xml:"sendTime,omitempty"`
		}{}
		if ut.Info.Content != nil {
			pub.Info.Content = ut.Info.Content
		}
		if ut.Info.SendTime != nil {
			pub.Info.SendTime = ut.Info.SendTime
		}
	}
	if ut.User != nil {
		pub.User = ut.User
	}
	return &pub
}

// message of user to send
type Message struct {
	Info *struct {
		// content of message
		Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
		// user send message time
		SendTime *time.Time `form:"sendTime,omitempty" json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	} `form:"info" json:"info" xml:"info"`
	User *string `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// Validate validates the Message type instance.
func (ut *Message) Validate() (err error) {
	if ut.Info == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "info"))
	}
	return
}
