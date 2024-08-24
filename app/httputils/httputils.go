package httputils

import (
	"net/http"
	"time"
)

type SessionHttp struct {
	cookieName string
	expiration time.Duration
}

func NewSessionHttp(cookieName string, expiration time.Duration) *SessionHttp {
	return &SessionHttp{
		cookieName: cookieName,
		expiration: expiration,
	}
}

func NewSessionHttpDefault() *SessionHttp {
	return &SessionHttp{
		cookieName: "default-session",
		expiration: 15 * time.Minute,
	}
}

func (s *SessionHttp) SetExpiration(expiration time.Duration) {
	s.expiration = expiration
}

func (s *SessionHttp) Set(w http.ResponseWriter, key string, value string) error {
	cookie := &http.Cookie{
		Name:    key,
		Value:   value,
		Expires: time.Now().Add(s.expiration),
		Path:    "/",
	}
	http.SetCookie(w, cookie)
	return nil
}

func (s *SessionHttp) Get(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil && err != http.ErrNoCookie {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func (s *SessionHttp) Remove(w http.ResponseWriter, key string) error {
	cookie := &http.Cookie{
		Name:    key,
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		MaxAge:  -1,
	}
	http.SetCookie(w, cookie)
	return nil
}

func (s *SessionHttp) Destroy(r *http.Request, w http.ResponseWriter) error {
	for _, cookie := range r.Cookies() {
		c := &http.Cookie{
			Name:    cookie.Name,
			Value:   "",
			Expires: time.Unix(0, 0),
			Path:    "/",
			MaxAge:  -1,
		}
		http.SetCookie(w, c)
	}
	return nil
}
