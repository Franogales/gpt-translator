package translate

type Translator interface {
	// TranslateES translates a text from English to Spanish
	TranslateES(text string) (string, error)
	// TranslateEN translates a text from Spanish to English
	TranslateEN(text string) (string, error)
}

var (
	PromptSystemSpanishToEnglish = "Eres un traductor profesional experto en varios idiomas. Tu tarea es traducir cualquier mensaje proporcionado por el usuario del español al ingles. Asegúrate de mantener el significado original, el tono y el contexto del mensaje. La traducción debe ser fluida y natural, como si hubiera sido escrita originalmente en el idioma de destino."
	PromptUserSpanishToEnglish   = `Por favor, traduce el siguiente texto del español al ingles, si tienes mas de una manera de traducirlo, por favor, proporciona solamente la primera opcion:

%s`
	PromptSystemEnglishToSpanish = "Eres un traductor profesional experto en varios idiomas. Tu tarea es traducir cualquier mensaje proporcionado por el usuario del ingles al español. Asegúrate de mantener el significado original, el tono y el contexto del mensaje. La traducción debe ser fluida y natural, como si hubiera sido escrita originalmente en el idioma de destino."
	PromptUserEnglishToSpanish   = `Por favor, traduce el siguiente texto del ingles al español:

%s`
)
