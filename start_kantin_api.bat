@echo off
REM filepath: d:\ProjectKantin\kantin_api\start_kantin_api.bat
cd /d D:\ProjectKantin\kantin_api

@REM REM Menjalankan PHP artisan serve di jendela baru
@REM start cmd /k "php artisan serve"

@REM REM Menjalankan npm run dev di jendela baru
@REM start cmd /k "npm run dev"

REM Menjalankan Go di jendela baru
start cmd /k "go run main.go"

pause