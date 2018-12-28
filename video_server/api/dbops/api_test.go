package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_inf")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("chensiqi", "123")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("chensiqi")
	if pwd != "123" || err != nil {
		t.Errorf("Errof of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("chensiqi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("chensiqi")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

//
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}
