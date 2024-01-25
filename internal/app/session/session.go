package session

import (
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

const (
	SessionKeyName = "APPLICATION_SESSION_KEY"
)

// Manager is a session manager.
type Manager struct {
	sessions map[string]Dictionary
	mutex    sync.RWMutex
}

// Dictionary is a session.
type Dictionary map[string]interface{}

// NewManager returns a new session manager.
func NewManager() *Manager {
	return &Manager{sessions: map[string]Dictionary{}}
}

// GetSession returns the session for the given request.
func (m *Manager) GetSession(w http.ResponseWriter, r *http.Request) Dictionary {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var key string

	if cookie, err := r.Cookie(SessionKeyName); err == nil {
		key = cookie.Value
	}

	if session, ok := m.sessions[key]; ok {
		return session
	}

	return m.createSession(w, r)
}

func (m *Manager) createSession(w http.ResponseWriter, r *http.Request) Dictionary {
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

	if _, ok := m.sessions[key]; !ok {
		m.sessions[key] = Dictionary{}
	}

	return m.sessions[key]
}

// Get returns the value for the given key.
func (d Dictionary) Get(key string) interface{} {
	return d[key]
}

// Set sets the value for the given key.
func (d Dictionary) Set(key string, value interface{}) {
	d[key] = value
}
