# Traductor desde el portapapeles
obten tu groq api key desde aqui https://console.groq.com/docs/api-keys

## Build in windows in silent
go build -ldflags "-H windowsgui" -o translator.exe .\cmd\
ejecuta el translator.exe

# Combinaciones de teclas
- Ctrl + Shift + T = Traduce el texto en el portapapeles al ingles
- Ctrl + Shift + Y = Traduce el texto en el portapapeles al español
- Ctrl + Shift + Q = Cierra la aplicación

# Uso
Selecciona el texto que quieras traducir
Presiona Ctrl + Shift + T para traducir al ingles
Presiona Ctrl + Shift + Y para traducir al español
el texto sera copiado al portapapeles y se intentara pegar en el campo de texto en el que se encuentre el cursor