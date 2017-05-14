package definitions

type LanguageInfo struct {
	Id               string `json:"id,omitempty"`
	Uri              string `json:"uri,omitempty"`
	Greeting         bool   `json:"greeting,omitempty"`
	FormattingLocale bool   `json:"formattingLocale,omitempty"`
	LocaleCode       string `json:"localeCode,omitempty"`
	Name             string `json:"name,omitempty"`
	Ui               bool   `json:"ui,omitempty"`
}
