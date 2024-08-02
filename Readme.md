## Build in windows in silent
go build -ldflags "-H windowsgui" -o translator.exe .\cmd\

# Combinaciones de teclas
- Ctrl + Shift + T = Traduce el texto en el portapapeles al ingles
- Ctrl + Shift + Y = Traduce el texto en el portapapeles al español
- Ctrl + Shift + Q = Cierra la aplicación


## Agregar la tarea al scheduler de windows
```
PowerShell Script para Crear una Tarea Programada
Crea un archivo .ps1 (por ejemplo, create-task.ps1) con el siguiente contenido:

```powershell
# Ruta del ejecutable de tu programa Go
$action = New-ScheduledTaskAction -Execute "C:\ruta\a\tu\programa\myprogram.exe"

# Trigger para ejecutar la tarea al iniciar sesión
$trigger = New-ScheduledTaskTrigger -AtLogon

# Configuración adicional de la tarea
$settings = New-ScheduledTaskSettingsSet -AllowStartIfOnBatteries -DontStopIfGoingOnBatteries

# Crear la tarea en el Programador de tareas
Register-ScheduledTask -Action $action -Trigger $trigger -Settings $settings -TaskName "MyGoProgram" -Description "Ejecuta mi traductor Go al iniciar sesión" -User "NT AUTHORITY\INTERACTIVE" -RunLevel Highest
```

Pasos para Ejecutar el Script de PowerShell
Asegúrate de que la política de ejecución de PowerShell permissions to run scripts:
Abre PowerShell como administrador y ejecuta:

```powershell
Set-ExecutionPolicy RemoteSigned
```

Ejecuta el script de PowerShell:
Navega al directorio donde guardaste el archivo create-task.ps1 y ejecútalo:

```powershell
.\create-task.ps1
```
Explicación del Script
Define la acción: Especifica el ejecutable que deseas ejecutar. Asegúrate de cambiar C:\ruta\a\tu\programa\myprogram.exe a la ruta correcta de tu programa compilado.
Define el trigger: Configura el disparador para que la tarea se ejecute al iniciar sesión.
Configura los ajustes adicionales: Configura la tarea para que permita el inicio si está en batería y no se detenga si pasa a batería.
Registra la tarea: Utiliza Register-ScheduledTask para crear la tarea con el nombre "MyGoProgram" y una descripción, y ejecutarla con los privilegios más altos.