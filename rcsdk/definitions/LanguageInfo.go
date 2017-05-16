package definitions

type LanguageInfo struct {
	FormattingLocale bool   `json:"formattingLocale,omitempty"`
	LocaleCode       string `json:"localeCode,omitempty"`
	Name             string `json:"name,omitempty"`
	Ui               bool   `json:"ui,omitempty"`
	Id               string `json:"id,omitempty"`
	Uri              string `json:"uri,omitempty"`
	Greeting         bool   `json:"greeting,omitempty"`
}
