package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Franogales/gpt-translator/translate"
	"github.com/joho/godotenv"
	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"github.com/ncruces/zenity"
	"golang.design/x/clipboard"
)

func main() {
	// validate if the .env file exists
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		// create the .env file
		_, err := os.Create(".env")
		if err != nil {
			zenity.Error(fmt.Sprintf("Error tratando de crear el archivo .env. intentalo de nuevo [%v]", err), zenity.Title("Error"))
			os.Exit(1)
		}
	}

	err = godotenv.Load()
	if err != nil {
		zenity.Error(fmt.Sprintf("Error tratando de cargar el archivo .env. intentalo de nuevo [%v]", err), zenity.Title("Error"))
		dir, err := os.Getwd()
		if err != nil {
			zenity.Error(fmt.Sprintf("Error tratando de obtener el directorio actual. intentalo de nuevo [%v]", err), zenity.Title("Error"))
			os.Exit(1)
		}
		zenity.Info(fmt.Sprintf("El archivo .env no se ha encontrado en el directorio %s. Por favor, crea un archivo .env con tu API Key de ChatGPT", dir), zenity.Title("Info"))
	}

	err = clipboard.Init()
	if err != nil {
		zenity.Error(fmt.Sprintf("Error tratando de inicializar el portapapeles. intentalo de nuevo [%v]", err), zenity.Title("Error"))
		os.Exit(1)
	}

	// chatgptApiKey := os.Getenv("GROQ_API_KEY")
	chatgptApiKey := os.Getenv("GEMINI_API_KEY")

	if chatgptApiKey == "" {
		chatgptApiKey, err = zenity.Entry("Por favor, introduce tu API Key de ChatGPT",
			zenity.Title("GROQ_API_KEY"),
		)
		if err != nil {
			zenity.Error(fmt.Sprintf("error tratando de obtener la API Key. intentalo de nuevo [%v]", err), zenity.Title("Error"))
			os.Exit(1)
		}

		// Save the API Key in the environment file
		err = godotenv.Write(map[string]string{
			"GROQ_API_KEY": chatgptApiKey,
		}, ".env")
		if err != nil {
			zenity.Error(fmt.Sprintf("error tratando de guardar la API Key. intentalo de nuevo [%v]", err), zenity.Title("Error"))
			os.Exit(1)
		}

		zenity.Info("API Key guardada correctamente", zenity.Title("Info"))
	}
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer keyboard.Uninstall()

	gptTranslator := translate.NewGeminiChat(chatgptApiKey)
	// gptTranslator := translate.NewLocalChat()

	// gptTranslator := translate.NewGroqChat(chatgptApiKey)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing keyboard input")
	pressedCtrl := false
	pressedAlt := false
	pressedT := false
	pressedQ := false
	pressedY := false
	for {
		select {
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return
		case k := <-keyboardChan:
			updatePressedKeys(&pressedCtrl, &pressedAlt, &pressedT, &pressedQ, &pressedY, k)
			if pressedCtrl && pressedAlt && pressedT {
				pressedCtrl = false
				pressedAlt = false
				pressedT = false
				fmt.Println("Traduciendo al inglés...")
				err = translateFromClipBoard(gptTranslator, "en")
				if err != nil {
					zenity.Error(fmt.Sprintf("un error ocurrio [%v] preciona ctrl+alt+q  para detener el programa", err), zenity.Title("Error"))
				}
				fmt.Println("Texto traducido y copiado al portapapeles")
				continue
			}

			if pressedCtrl && pressedAlt && pressedY {
				pressedCtrl = false
				pressedAlt = false
				pressedY = false
				fmt.Println("Traduciendo al español...")
				err = translateFromClipBoard(gptTranslator, "es")
				if err != nil {
					zenity.Error(fmt.Sprintf("un error ocurrio [%v] preciona ctrl+alt+q  para detener el programa", err), zenity.Title("Error"))
				}
				fmt.Println("Texto traducido y copiado al portapapeles")
				continue
			}

			if pressedCtrl && pressedAlt && pressedQ {
				fmt.Println("Received shutdown signal")
				return
			}
			// fmt.Printf("Received %v %v\n", k.Message, k.VKCode)
			continue
		}
	}
}

func translateFromClipBoard(gptTranslator translate.Translator, lang string) (err error) {
	text := clipboard.Read(clipboard.FmtText)

	var translatedText string

	switch lang {
	case "es":
		translatedText, err = gptTranslator.TranslateES(string(text))
		if err != nil {
			return fmt.Errorf("error tratando de traducir el texto. intentalo de nuevo [%v]", err)
		}
	default:
		translatedText, err = gptTranslator.TranslateEN(string(text))
		if err != nil {
			return fmt.Errorf("error tratando de traducir el texto. intentalo de nuevo [%v]", err)
		}
	}

	clipboard.Write(clipboard.FmtText, []byte(translatedText))
	zenity.Notify("Texto traducido y copiado al portapapeles", zenity.Title("Info"))
	return nil
}

func updatePressedKeys(pressedCtrl, pressedAlt, pressedT, pressedQ, pressedY *bool, k types.KeyboardEvent) {
	if ctrlPressed := k.VKCode == types.VK_LCONTROL && k.Message == types.WM_KEYDOWN; ctrlPressed {
		*pressedCtrl = true
	}
	if altPressed := k.VKCode == types.VK_LMENU && k.Message == types.WM_KEYDOWN; altPressed {
		*pressedAlt = true
	}
	if altPressed := k.VKCode == types.VK_LMENU && k.Message == types.WM_SYSKEYDOWN; altPressed {
		*pressedAlt = true
	}
	if tPressed := k.VKCode == types.VK_T && k.Message == types.WM_KEYDOWN; tPressed {
		*pressedT = true
	}
	if qPressed := k.VKCode == types.VK_Q && k.Message == types.WM_KEYDOWN; qPressed {
		*pressedQ = true
	}
	if yPressed := k.VKCode == types.VK_Y && k.Message == types.WM_KEYDOWN; yPressed {
		*pressedY = true
	}

	if ctrlNotPressed := k.VKCode == types.VK_LCONTROL && k.Message == types.WM_KEYUP; ctrlNotPressed {
		*pressedCtrl = false
	}
	if altNotPressed := k.VKCode == types.VK_LMENU && k.Message == types.WM_KEYUP; altNotPressed {
		*pressedAlt = false
	}
	if altNotPressed := k.VKCode == types.VK_LMENU && k.Message == types.WM_SYSKEYUP; altNotPressed {
		*pressedAlt = false
	}
	if tNotPressed := k.VKCode == types.VK_T && k.Message == types.WM_KEYUP; tNotPressed {
		*pressedT = false
	}
	if qNotPressed := k.VKCode == types.VK_Q && k.Message == types.WM_KEYUP; qNotPressed {
		*pressedQ = false
	}
	if yNotPressed := k.VKCode == types.VK_Y && k.Message == types.WM_KEYUP; yNotPressed {
		*pressedY = false
	}
}
