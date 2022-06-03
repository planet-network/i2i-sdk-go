package mobile

import (
	"fmt"
	"testing"
)

func TestManager(t *testing.T) {
	c := NewManager()
	fmt.Println("register")
	if err := c.Register("test_user", "qwert"); err != nil {
		t.Fatalf("register: %v", err)
	}

	defer func() {
		fmt.Println("remove")
		if err := c.UserRemove(); err != nil {
			t.Fatalf("remove: %v", err)
		}
	}()

	fmt.Println("login:")
	if err := c.Login("test_user", "qwert"); err != nil {
		t.Fatalf("login: %v", err)
	}

	fmt.Println("plan list:")
	if err := c.PlanList(); err != nil {
		t.Fatalf("plan list: %v", err)
	}
}
