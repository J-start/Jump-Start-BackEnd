@echo off

docker-compose down -v

docker-compose -f docker-compose.test.yaml up -d

echo Aguardando o container subir...

for /f %%i in ('docker-compose ps -q mysql') do set CONTAINER_ID=%%i

for /f "tokens=*" %%i in ('docker inspect -f "{{.Name}}" %CONTAINER_ID%') do set CONTAINER_NAME=%%i

set CONTAINER_NAME=%CONTAINER_NAME:~1%

echo O container %CONTAINER_NAME% iniciou.

:wait_loop
docker inspect -f "{{.State.Health.Status}}" %CONTAINER_NAME% | find "healthy" >nul
if %errorlevel% neq 0 (
    timeout /t 2 >nul
    goto wait_loop
)

timeout /t 10 >nul

echo Container %CONTAINER_NAME% está pronto. Executando os testes...

go test ./...

docker-compose down -v