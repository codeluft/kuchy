package app

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

const (
	SessionKeyName = "APPLICATION_SESSION_KEY"
)

// SessionManager is a session manager.
type SessionManager struct {
	sessions map[string]*Session
}

// Session is a session.
type Session map[string]interface{}

// NewSessionManager returns a new session manager.
func NewSessionManager() *SessionManager {
	return &SessionManager{sessions: map[string]*Session{}}
}

// GetSession returns the session for the given request.
func (s *SessionManager) GetSession(w http.ResponseWriter, r *http.Request) *Session {
	var key string

	if cookie, err := r.Cookie(SessionKeyName); err == nil {
		key = cookie.Value
	}

	if session, ok := s.sessions[key]; ok {
		return session
	}

	return s.createSession(w, r)
}

func (s *SessionManager) createSession(w http.ResponseWriter, r *http.Request) *Session {
	var key = uuid.NewString()

	cookie, err := r.Cookie(SessionKeyName)
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:     SessionKeyName,
			Value:    key,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			Expires:  time.Now().Add(24 * time.Hour),
		})
	} else {
		key = cookie.Value
	}

	if _, ok := s.sessions[key]; !ok {
		s.sessions[key] = &Session{}
	}

	return s.sessions[key]
}

func (s *Session) Get(key string) interface{} {
	return (*s)[key]
}

func (s *Session) Set(key string, value interface{}) {
	(*s)[key] = value
}
