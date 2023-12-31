# File Storage

Простое и эффективное решение для хранения и извлечения файлов через HTTP API.

# Введение

File Storage - это серверное приложение, созданное для управления файлами. Основная идея заключается в предоставлении API для загрузки, хранения и извлечения файлов. Приложение также предлагает возможность мониторинга с помощью Prometheus.

# Основные возможности

* Многократное Хранение: Поддерживает несколько директорий для хранения, что обеспечивает гибкость и расширяемость.
* Prometheus Интеграция: Включает в себя конечные точки для мониторинга метрик.
* Расширяемый: Легко адаптировать и модифицировать для различных нужд и окружений.

# Установка и Запуск

Для запуска тестов необходимо выполнить одну из команд:
1. `make test` запуск всех тестов
2. `make cover` для запуска тестов с покрытием
3. `make cover-html` для запуска тестов с покрытием и получения отчёта в html формате

Для запуска линтера необходимо выполнить команду `make lint`

Остальные команды можно получить выполнив команду `make help`

## Требования:
* Go (версия 1.15 и выше)
* Подключение к интернету (для загрузки зависимостей)

## Шаги для установки
* Клонировать репозиторий:  
`git clone [URL вашего репозитория]`
* Перейти в директорию проекта:
`cd file-storage`
* Настроить config.yaml:  
Убедитесь, что вы настроили файл конфигурации согласно вашим требованиям. Пример содержания:
```
ServerPort: ":8080"
StorageDirs:
- "./storage1"
- "./storage2"
```
* Запустить приложение:
`go run .`

# HTTP API

* Загрузить файл
Маршрут: /upload  
Метод: POST  
Тело запроса: Бинарные данные файла.  
Описание: Загрузите файл на сервер. Файл должен быть включен в тело запроса.  
* Скачать файл  
Маршрут: /download?fileid=[ID файла]  
Метод: GET  
Описание: Скачайте файл с сервера по его ID.  
* Проверка статуса  
Маршрут: /status  
Метод: GET  
Описание: Получите текущий статус сервера.  
* Метрики Prometheus  
Маршрут: /metrics  
Метод: GET  
Описание: Используйте для мониторинга с Prometheus.  

# Содействие

Мы рады вашим предложениям и исправлениям. Пожалуйста, создайте pull request или issue в репозитории.