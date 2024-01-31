# PDF Machine
Fast and secure based on C++ and Golang service for generate PDF from URL

## Локальный запуск приложения

```shell
docker-compose -f docker-compose.yml -f docker-compose.override.yml up --remove-orphans -d --build
```

## Пример

```shell
http://localhost:8000/convert?url=https://neometria.ru/about&format=A4
```

параметр ```format``` необходимо передавать в виде A1, A2, и т.д.

## Проект основан на основе наработок
```https://github.com/SebastiaanKlippert/go-wkhtmltopdf/blob/master/wkhtmltopdf.go```
