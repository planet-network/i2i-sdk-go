package mobile

import (
	"fmt"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	manager := NewManager()

	if _, err := manager.Register("qwerty000000000111111111111111111"); err != nil {
		t.Fatalf("Register(): failed: %s", err)
	}

	defer func() {
		if err := manager.ClientRemove(); err != nil {
			t.Fatalf("ClientRemove(): failed: %s", err)
		}
	}()

	if err := manager.Login(); err != nil {
		t.Fatalf("Login(): failed: %s", err)
	}

	token := GenerateToken()
	if err := manager.NodeOder(token); err != nil {
		t.Fatalf("NodeOder(): failed: %s", err)
	}
	time.Sleep(time.Second)

	node, err := manager.NodeShow()
	if err != nil {
		t.Fatalf("NodeShow(): failed: %s", err)
	}

	fmt.Printf("%+v\n", node)

	defer func() {
		if err := manager.NodeRemove(); err != nil {
			t.Fatalf("ClientRemove(): failed: %s", err)
		}
	}()

}

//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestManager(t *testing.T) {
//	c := NewManager()
//	fmt.Println("register")
//	if err := c.Register("test_user", "qwert"); err != nil {
//		t.Fatalf("register: %v", err)
//	}
//
//	defer func() {
//		fmt.Println("remove")
//		if err := c.UserRemove(); err != nil {
//			t.Fatalf("remove: %v", err)
//		}
//	}()
//
//	fmt.Println("login:")
//	if err := c.Login("test_user", "qwert"); err != nil {
//		t.Fatalf("login: %v", err)
//	}
//
//	fmt.Println("plan list:")
//	if err := c.PlanList(); err != nil {
//		t.Fatalf("plan list: %v", err)
//	}
//}
