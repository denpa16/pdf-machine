# PDF Machine
Fast and secure based on C++ and Golang service for generate PDF from URL

## Локальный запуск приложения

```shell
docker-compose -f docker-compose.yml -f docker-compose.override.yml up --remove-orphans -d --build
```

## Пример

```shell
http://localhost:8000/convert?url=https://neometria.ru/about
```
