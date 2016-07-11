package notify

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotify(t *testing.T) {
	n := NewNotifications()

	msg := &Notification{
		Title:    "Your Lantern time is up",
		Message:  "You have reached your data cap limit",
		ClickURL: "https://www.getlantern.org",
		IconURL:  "http://127.0.0.1:2000/img/lantern_logo.png",
	}

	err := n.Notify(msg)
	assert.Nil(t, err, "got an error notifying user")
	time.Sleep(3 * time.Second)
}