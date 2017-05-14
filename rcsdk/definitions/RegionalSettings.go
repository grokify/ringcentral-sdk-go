package definitions

type RegionalSettings struct {
	GreetingLanguage GreetingLanguageInfo `json:"greetingLanguage,omitempty"`
	FormattingLocale FormattingLocaleInfo `json:"formattingLocale,omitempty"`
	HomeCountry      CountryInfo          `json:"homeCountry,omitempty"`
	Timezone         TimezoneInfo         `json:"timezone,omitempty"`
	Language         LanguageInfo         `json:"language,omitempty"`
}
