package sessions

import (
	"rushflow/api/dbops"
	"rushflow/api/defs"
	"rushflow/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

// 初始化
func init() {
	sessionMap = &sync.Map{}
}

func NowInMilli() int64 {
	return time.Now().UnixNano() / 100000
}

func DeleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// 加载session
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

// 生成新session
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := NowInMilli()     // 毫秒
	ttl := ct + 30*60*1000 // Severside session valid time
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}

// session是否过期
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := NowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expired session
			DeleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
