package mobile

import "testing"

func TestPlanetConnect(t *testing.T) {
	pc := New()
	if err := pc.Connect(); err != nil {
		t.Fatalf("connect(): failed: %s", err)
	}

	if err := pc.Register("some_user@example.com", "qwerty", "sendgrid"); err != nil {
		t.Logf("Register(): failed: %s", err)
	}

	if err := pc.Login("some_user@example.com", "qwerty", "sendgrid"); err != nil {
		t.Fatalf("Login(): failed: %s", err)
	}

	if err := pc.PersonalDataAdd("name", "jakub"); err != nil {
		t.Errorf("PersonalDataAdd(): failed: %s", err)
	}

	data, err := pc.PersonalDataGet("name")
	if err != nil {
		t.Errorf("PersonalDataGet(): failed: %s", err)
	}

	if data != "jakub" {
		t.Fatalf("failed to get personal data")
	}
}
