Test task BackDev
# Тестовое задание: Сервис аутентификации

## Используемые технологии:

- Go
- JWT
- MongoDB
- Docker

## Запуск проекта:

1. Убедитесь, что у вас установлен Docker и Docker Compose.
2. Запустить проект

   ```bash
   docker-compose up

## Тестирование:
### Ручка с запросом токенов:
http://localhost:8002/auth/login:
```JSON
{
    "id":"3121bd43-24bc-426a-bc61-28ec7aca2127"
}
```
### Ручка с обновлением токенов:
http://localhost:8002/auth/refresh:
```JSON
{
    "accessToken": "",
    "refreshToken": ""
}
```
