package manager

import (
	"fmt"
	"testing"
)

func TestManager(t *testing.T) {
	c := NewClientV1()
	fmt.Println("register")
	if err := c.Register("adolf", "qwert"); err != nil {
		t.Fatalf("register: %v", err)
	}

	defer func() {
		fmt.Println("remove")
		if err := c.UserRemove(); err != nil {
			t.Fatalf("remove: %v", err)
		}
	}()

	fmt.Println("login")
	if err := c.Login("adolf", "qwert"); err != nil {
		t.Fatalf("login: %v", err)
	}

	fmt.Println("plan list")
	if err := c.PlanList(); err != nil {
		t.Fatalf("plan list: %v", err)
	}
}
