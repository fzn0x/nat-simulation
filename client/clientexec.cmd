@echo off
rem Starts the client. Extra flags are passed through; run with -h to list them.
cd /d "%~dp0.."
go run ./client %*
