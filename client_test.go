package xenAPI

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

var LoginSessionId = 0

func TestAuthentication(t *testing.T) {

	SessionClassLoginWithPasswordMockedCallback = func(uname string, pwd string, version string, originator string) (_retval SessionRef, _err error) {
		LoginSessionId++
		log.Printf("New connection uname: %v, pwd: %v, version: %v, originator: %v, LoginSessionId: %v", uname, pwd, version, originator, LoginSessionId)
		return SessionRef("Login " + string(LoginSessionId)), nil
	}

	SessionClassLogoutMockedCallback = func(sessionID SessionRef) (_err error) {
		log.Printf("Session Logout %v", sessionID)
		return nil
	}
	client, _ := NewClient("http://localhost:40080", nil)

	sessionRef, err := client.Session.LoginWithPassword("terraform", "testing", "1.0", "terraform")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	err = client.Session.Logout(sessionRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	log.SetOutput(os.Stdout)
	os.Exit(m.Run())
}
