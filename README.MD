28.05.2026 

Здравствуйте!

Комментарии по коду:
Для удобства разработки сделал Makefile, чтобы не держать в голове длинные команды 
Также я намеренно не скрлы .env в gitignore, чтобы вручную вам не пришлось настраивать окружение, 
но я знаю в проде необходимо скрыть имеющиеся там чувствительные данные, так что надеюсь, вы не зачтете это за ошибку

Запуск приложения:
1. Клонировать репозиторий, PowerShell >> "gh repo clone cephalopagus/bkv-intalant-task"
2. Перейти в корень папки, PowerShell >> "cd bkv-intalant-task"
3. Запустить, PowerShell >> "docker compose up --build"

Тестирование(Postman):
1. Создать новый запрос:
    Метод: POST
    URL: http://localhost:8080/departments/
    Вкладка Body → raw → JSON 
2. Тело запроса: 
    {"name": "Mobile", "parent_id": 1}
3. Send

Тестирование(PowerShell):
1. Ввести команду: 
    Invoke-WebRequest -Uri http://localhost:8080/departments/ `
    -Method POST `
    -ContentType "application/json" `
    -Body '{"name": "Mobile", "parent_id": 1}'

Завершить программу:
1. PowerShell >> "docker compose down"


docker compose run --rm -e GOOSE_COMMAND=up migrate
docker compose run --rm -e GOOSE_COMMAND=down migrate