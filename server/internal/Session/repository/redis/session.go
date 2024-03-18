package repository

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"server/internal/domain/entity"
	"sync"
)

type SessionManager struct {
	redisConn redis.Conn
	mu        sync.Mutex
}

func NewSessionManager(conn redis.Conn) *SessionManager {
	return &SessionManager{
		redisConn: conn,
	}
}

func (sm *SessionManager) Create(cookie *entity.Cookie) error {
	dataSerialized, _ := json.Marshal(cookie.UserID)
	mkey := "sessions:" + cookie.SessionToken
	sm.mu.Lock()
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 540000))
	sm.mu.Unlock()
	if err != nil || result != "OK" {
		return err
	}

	return nil
}

func (sm *SessionManager) Check(sessionToken string) (*entity.Cookie, error) {
	mkey := "sessions:" + sessionToken
	sm.mu.Lock()
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	sm.mu.Unlock()
	if err != nil {
		if err != redis.ErrNil {
			return nil, err
		}
		return nil, nil
	}

	cookie := &entity.Cookie{}
	err = json.Unmarshal(data, &cookie.UserID)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (sm *SessionManager) Delete(cookie *entity.DBDeleteCookie) error {
	mkey := "sessions:" + cookie.SessionToken
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		if err != redis.ErrNil {
			return err
		}
		return nil
	}

	return nil
}
