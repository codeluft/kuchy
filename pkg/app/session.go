package app

type sessionMap map[string]interface{}

func NewSession() sessionMap {
	return sessionMap{}
}

func (s sessionMap) Set(key string, value interface{}) {
	s[key] = value
}

func (s sessionMap) Get(key string) interface{} {
	return s[key]
}
