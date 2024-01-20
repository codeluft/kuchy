package app

import (
	"embed"
	"errors"
	"gopkg.in/yaml.v2"
)

var (
	ErrCannotOpenTranslationDirectory = errors.New("cannot open translation directory")
	ErrCannotReadTranslationFile      = errors.New("cannot read translation file")
	ErrCannotUnmarshalTranslationFile = errors.New("cannot unmarshal translation file")
	ErrLanguageNotFound               = errors.New("language not found")
)

const (
	DefaultLanguage = "en"
)

type Translator struct {
	dict map[string]map[string]string
	lang string
}

// TranslatorFunc is a function that translates a key to a string.
type TranslatorFunc func(string, string) string

// NewTranslator returns a new Translator.
func NewTranslator(dir embed.FS) (*Translator, error) {
	var t = &Translator{
		dict: make(map[string]map[string]string),
		lang: DefaultLanguage,
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
func (t *Translator) Translate(key string) string {
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
func (t *Translator) SetLanguage(lang string) error {
	if _, ok := t.dict[lang]; !ok {
		return ErrLanguageNotFound
	}
	t.lang = lang

	return nil
}

// GetLanguage returns the current language.
func (t *Translator) GetLanguage() string {
	return t.lang
}

// GetDict returns the current dictionary.
func (t *Translator) GetDict() map[string]string {
	return t.dict[t.lang]
}
