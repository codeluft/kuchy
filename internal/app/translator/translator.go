package translator

import (
	"embed"
	"errors"
	"gopkg.in/yaml.v2"
	"sync"
)

var (
	// ErrCannotOpenTranslationDirectory is an error that occurs when the translation directory cannot be opened.
	ErrCannotOpenTranslationDirectory = errors.New("cannot open translation directory")

	// ErrCannotReadTranslationFile is an error that occurs when the translation file cannot be read.
	ErrCannotReadTranslationFile = errors.New("cannot read translation file")

	// ErrCannotUnmarshalTranslationFile is an error that occurs when the translation file cannot be unmarshalled.
	ErrCannotUnmarshalTranslationFile = errors.New("cannot unmarshal translation file")

	// ErrLanguageNotFound is an error that occurs when the language is not found.
	ErrLanguageNotFound = errors.New("language not found")
)

const (
	// DefaultLanguage is the default language.
	DefaultLanguage = "en"
)

// Loader loads translations and expose translation functions.
type Loader struct {
	dict  map[string]map[string]string
	lang  string
	mutex *sync.RWMutex
}

// Func is a function that translates a key to a string.
type Func func(string, string) string

// New returns a new TranslatorImpl.
func New(dir embed.FS) (*Loader, error) {
	var t = &Loader{
		dict:  make(map[string]map[string]string),
		lang:  DefaultLanguage,
		mutex: new(sync.RWMutex),
	}

	dirEntry, err := dir.ReadDir("translations")
	if err != nil {
		return nil, errors.Join(ErrCannotOpenTranslationDirectory, err)
	}
	for _, file := range dirEntry {
		if file.IsDir() {
			continue
		}
		fname := file.Name()
		lang := fname[:len(fname)-len(".yaml")]
		t.dict[lang] = make(map[string]string)

		yamlFile, err := dir.ReadFile("translations/" + fname)
		if err != nil {
			return nil, errors.Join(ErrCannotReadTranslationFile, err)
		}

		err = yaml.Unmarshal(yamlFile, t.dict[lang])
		if err != nil {
			return nil, errors.Join(ErrCannotUnmarshalTranslationFile, err)
		}
	}

	return t, nil
}

// Translate translates a key to the current language.
func (t *Loader) Translate(key string) string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	var dict, ok = t.dict[t.lang]
	if !ok {
		dict, ok = t.dict[DefaultLanguage]
		if !ok {
			return key
		}
	}

	if _, ok := dict[key]; !ok {
		return key
	}

	return dict[key]
}

// SetLanguage sets the current language.
func (t *Loader) SetLanguage(lang string) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if _, ok := t.dict[lang]; !ok {
		return ErrLanguageNotFound
	}
	t.lang = lang

	return nil
}

// GetLanguage returns the current language.
func (t *Loader) GetLanguage() string {
	return t.lang
}

// GetDict returns the current dictionary.
func (t *Loader) GetDict() map[string]string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.dict[t.lang]
}
