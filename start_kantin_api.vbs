Set WhsShell = CreateObject("WScript.Shell")
projectPath = "D:\ProjectKantin\kantin_api"
WshShell.Run "cmd.exe /k cd /d """ & projectPath & """ & && go run main.go", 1,false