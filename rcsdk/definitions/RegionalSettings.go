package definitions

type RegionalSettings struct {
	FormattingLocale FormattingLocaleInfo `json:"formattingLocale,omitempty"`
	GreetingLanguage GreetingLanguageInfo `json:"greetingLanguage,omitempty"`
	HomeCountry      CountryInfo          `json:"homeCountry,omitempty"`
	Language         LanguageInfo         `json:"language,omitempty"`
	Timezone         TimezoneInfo         `json:"timezone,omitempty"`
}
