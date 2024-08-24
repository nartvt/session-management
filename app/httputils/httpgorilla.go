package httputils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/nartvt/session-management/app/util"
)

type SessionGorilla struct {
	store      sessions.Store
	name       string
	secretKey  string
	expiration time.Duration
}

func NewGorillaStore(secretKey string) sessions.Store {
	return sessions.NewCookieStore([]byte(secretKey))
}

func NewSessionGorilla(store sessions.Store, name string, expiration time.Duration) *SessionGorilla {
	return &SessionGorilla{
		store:      store,
		name:       name,
		expiration: expiration,
	}
}

func NewSessionGorillaDefault() *SessionGorilla {
	secretKey := util.GenerateUUID()
	store := sessions.NewCookieStore([]byte(secretKey))
	return &SessionGorilla{
		store:      store,
		name:       util.String(15),
		secretKey:  secretKey,
		expiration: 15 * time.Minute,
	}
}

func (s *SessionGorilla) GetSecretKey() string {
	return s.secretKey
}

func (s *SessionGorilla) SetExpiration(expiration time.Duration) {
	s.expiration = expiration
}

func (s *SessionGorilla) Set(r *http.Request, w http.ResponseWriter, key string, value string) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}
	session.Values[key] = value
	session.Options.MaxAge = int(s.expiration.Seconds())
	return session.Save(r, w)
}

func (s *SessionGorilla) Get(r *http.Request, key string) (string, error) {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return "", err
	}
	value, ok := session.Values[key].(string)
	if !ok {
		return "", fmt.Errorf("%s", "Key not found or not a string")
	}
	return value, nil
}

func (s *SessionGorilla) Remove(r *http.Request, w http.ResponseWriter, key string) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}
	delete(session.Values, key)
	return session.Save(r, w)
}

func (s *SessionGorilla) Destroy(r *http.Request, w http.ResponseWriter) error {
	session, err := s.store.Get(r, s.name)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return session.Save(r, w)
}
