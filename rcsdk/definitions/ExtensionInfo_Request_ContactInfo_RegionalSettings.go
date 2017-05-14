package definitions

type ExtensionInfo_Request_ContactInfo_RegionalSettings struct {
	GreetingLanguage ExtensionInfo_Request_ContactInfo_RegionalSettings_GreetingLanguage `json:"greetingLanguage,omitempty"`
	FormattingLocale ExtensionInfo_Request_ContactInfo_RegionalSettings_FormattingLocale `json:"formattingLocale,omitempty"`
	Timezone         ExtensionInfo_Request_ContactInfo_RegionalSettings_Timezone         `json:"timezone,omitempty"`
	Language         ExtensionInfo_Request_ContactInfo_RegionalSettings_Language         `json:"language,omitempty"`
}
