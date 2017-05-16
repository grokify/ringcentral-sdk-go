package definitions

type ExtensionInfo_Request_ContactInfo_RegionalSettings struct {
	FormattingLocale ExtensionInfo_Request_ContactInfo_RegionalSettings_FormattingLocale `json:"formattingLocale,omitempty"`
	GreetingLanguage ExtensionInfo_Request_ContactInfo_RegionalSettings_GreetingLanguage `json:"greetingLanguage,omitempty"`
	Language         ExtensionInfo_Request_ContactInfo_RegionalSettings_Language         `json:"language,omitempty"`
	Timezone         ExtensionInfo_Request_ContactInfo_RegionalSettings_Timezone         `json:"timezone,omitempty"`
}
