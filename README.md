# Проект MongoDB и GoLang

## Описание
Этот проект представляет собой простое веб-приложение, которое использует MongoDB в качестве базы данных и GoLang для обработки запросов. Он позволяет создавать, получать и удалять пользователей.

## Установка
1. Клонируйте репозиторий на свой локальный компьютер.
2. Установите MongoDB и запустите сервер базы данных.
3. Установите GoLang.
4. Запустите приложение, выполнив команду `go run main.go` в корневом каталоге проекта.

## Использование
Отправьте HTTP-запросы на `localhost:8000` для взаимодействия с приложением. Вы можете использовать следующие эндпоинты:

- `GET /user/:id`: Получить информацию о пользователе с указанным ID.
- `POST /user`: Создать нового пользователя. Тело запроса должно содержать информацию о пользователе в формате JSON.
- `DELETE /user/:id`: Удалить пользователя с указанным ID.


