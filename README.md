### Как запустить

```shell script
APP_DEBUG= APP_INTERRUPTED_LIFETIME=10 APP_SERVER_PORT=8081 go run cmd/main.go
```

Где:
- APP_DEBUG - включит логгирование http запросов
- APP_INTERRUPTED_LIFETIME - время жизни трансялции в статусе interrupted, в сек
- APP_SERVER_PORT - порт для http сервера