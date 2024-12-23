# Testing System API
Серверная часть для модуля "Тестирование сотрудника" системы "Компетентум" предназначено для управления процессом оценки компетенций сотрудников.

## Настройка сервера

### Cхема базы данных
* Схема базы данных распологается `migrations/create.tables.testing-system.sql`
* Данные для заполнения таблиц находятся `migrations/fill.tables.testing-system.sql`

### Переменные окружения
Сервер использует несколько переменных окружения для конфигурации. Эти переменные загружаются из файла `.env.local`, в другом случае из окружения.

Структура переменных окружения следующая:
* DB_PASSWORD: Пароль для подключения к базе данных.

### Конфигурация сервера
Для API необходимо создать файл конфигурации `configs/config.local.json` для настройки параметров сервера и базы данных. Без него сервер не запустится.
Пример:
``` json
{
  "server": {
    "port": ""
  },
  "testing-system-database": {
    "host": "",
    "port": "",
    "username": "",
    "db_name": "",
    "ssl_mode": ""
  },
}
```

## Аутентификация
Все запросы к API требуют авторизации и защищены с использованием токенов.

### Аутентификация с помощью JWT
Для получения токена отправьте запрос на соответствующую конечную точку аутентификации.

Добавьте полученный токен в заголовок запроса:
Authorization: Bearer <ваш-токен>
