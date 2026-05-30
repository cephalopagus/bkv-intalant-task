28.05.2026 

Здравствуйте!

Комментарии по коду:
Для удобства разработки сделал Makefile, чтобы не держать в голове длинные команды 
Также я намеренно не скрлы .env в gitignore, чтобы вручную вам не пришлось настраивать окружение, 
но я знаю в проде необходимо скрыть имеющиеся там чувствительные данные, так что надеюсь, вы не зачтете это за ошибку
создал общий файл для ошибок, middlware и логгирования, где все логи пишутся 'out/logs'
сделал один короткий тест "для наглядности"
Надеюсь вам понравится :)

/---------------------------------------------------------------/

Архитектура:
1. cmd/app/main.go - точка входа

2. internal/core/errors - общие ошибки
3. internal/core/logger - логика логгера, запись в файл
4. internal/core/middleware - общая миддлваря для логгирования запросоы

5. internal/core/response - хелперы
6. internal/core/postgres - подключение к бд

7. internal/feature/department/repository - работа с БД
8. internal/feature/department/service - логика
9. internal/feature/department/transport - хендлеры

10. internal/feature/employee/repository - работа с БД
11. internal/feature/employee/service - логика
12. internal/feature/employee/transport - хендлеры

13. migration - SQL миграции

14. out/logs - файлы с логами

/---------------------------------------------------------------/

Стек:
1. Go - 1.25.7
2. PostgreSQL - 18
3. Docker + docker-compose
4. GORM 
5. Goose migration

/---------------------------------------------------------------/

Запуск приложения:
1. Клонировать репозиторий, PowerShell >> "gh repo clone cephalopagus/bkv-intalant-task"
2. Перейти в корень папки, PowerShell >> "cd bkv-intalant-task"
3. Запустить, PowerShell >> "docker compose up --build"

/---------------------------------------------------------------/

Тестирование(Postman):
1. Создать новый запрос:
    Метод: POST
    URL: http://localhost:8080/departments/
    Вкладка Body → raw → JSON 
2. Тело запроса: 
    {"name": "Mobile", "parent_id": 1}
3. Send

/---------------------------------------------------------------/

Тестирование(PowerShell):
1. Ввести команду: 
    Invoke-WebRequest -Uri http://localhost:8080/departments/ `
    -Method POST `
    -ContentType "application/json" `
    -Body '{"name": "Mobile", "parent_id": 1}'

/---------------------------------------------------------------/

Завершить программу:
1. PowerShell >> "docker compose down"

/---------------------------------------------------------------/

Миграции(если нужно):
1. PowerShell >> docker compose run --rm -e GOOSE_COMMAND=up migrate
2. PowerShell >> docker compose run --rm -e GOOSE_COMMAND=down migrate