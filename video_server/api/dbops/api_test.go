package dbops

import "testing"

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_inf")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// 用户测试工作流
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
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

// 视频测试工作流
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo:%v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo:%v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo:%v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo:%v", err)
	}
}
