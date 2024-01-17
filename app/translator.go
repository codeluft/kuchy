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

type translator struct {
	dict map[string]map[string]string
	lang string
}

// TranslatorFunc is a function that translates a key to a string.
type TranslatorFunc func(string, string) string

// NewTranslator returns a new Translator.
func NewTranslator(dir embed.FS) (*translator, error) {
	var t = &translator{
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
func (t *translator) Translate(key, placeholder string) string {
	if _, ok := t.dict[t.lang]; !ok {
		return placeholder
	}

	if _, ok := t.dict[t.lang][key]; !ok {
		return placeholder
	}

	return t.dict[t.lang][key]
}

// SetLanguage sets the current language.
func (t *translator) SetLanguage(lang string) error {
	if _, ok := t.dict[lang]; !ok {
		return ErrLanguageNotFound
	}
	t.lang = lang

	return nil
}

// GetLanguage returns the current language.
func (t *translator) GetLanguage() string {
	return t.lang
}

// GetDict returns the current dictionary.
func (t *translator) GetDict() map[string]string {
	return t.dict[t.lang]
}
