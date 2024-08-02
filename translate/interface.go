package translate

type Translator interface {
	// TranslateES translates a text from English to Spanish
	TranslateES(text string) (string, error)
	// TranslateEN translates a text from Spanish to English
	TranslateEN(text string) (string, error)
}
