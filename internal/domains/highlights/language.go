package highlights

type Language string

const (
	LangFrenchCA  Language = "fr-CA"
	LangFrenchFR  Language = "fr-FR"
	LangFrenchBE  Language = "fr-BE"
	LangFrenchCH  Language = "fr-CH"
	LangEnglishUS Language = "en-US"
	LangEnglishGB Language = "en-GB"
	LangEnglishAU Language = "en-AU"
	LangEnglishCA Language = "en-CA"
)

func IsValidLanguage(s string) bool {
	switch Language(s) {
	case LangFrenchCA, LangFrenchFR, LangFrenchBE, LangFrenchCH,
		LangEnglishUS, LangEnglishGB, LangEnglishAU, LangEnglishCA:
		return true
	default:
		return false
	}
}

func (l Language) DisplayName() string {
	switch l {
	case LangFrenchCA:
		return "French (Canada)"
	case LangFrenchFR:
		return "French (France)"
	case LangFrenchBE:
		return "French (Belgium)"
	case LangFrenchCH:
		return "French (Switzerland)"
	case LangEnglishUS:
		return "English (US)"
	case LangEnglishGB:
		return "English (UK)"
	case LangEnglishAU:
		return "English (Australia)"
	case LangEnglishCA:
		return "English (Canada)"
	default:
		return string(l)
	}
}
