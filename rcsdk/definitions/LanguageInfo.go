package definitions

type LanguageInfo struct {
	FormattingLocale bool   `json:"formattingLocale,omitempty"`
	Greeting         bool   `json:"greeting,omitempty"`
	Id               string `json:"id,omitempty"`
	LocaleCode       string `json:"localeCode,omitempty"`
	Name             string `json:"name,omitempty"`
	Ui               bool   `json:"ui,omitempty"`
	Uri              string `json:"uri,omitempty"`
}
