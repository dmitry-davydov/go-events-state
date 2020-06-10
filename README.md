### Как запустить

```shell script
APP_DEBUG= APP_INTERRUPTED_LIFETIME=10 APP_SERVER_PORT=8081 go run cmd/main.go
```

Где:
- APP_DEBUG - включит логгирование http запросов
- APP_INTERRUPTED_LIFETIME - время жизни трансялции в статусе interrupted, в сек
- APP_SERVER_PORT - порт для http сервера


### Методы Api

##### Список событий
```
GET /api/v1/events
```

```json
{
  "data": [
    {
      "type": "stream",
      "id": "a23626f1-8db6-4e73-86b2-fc10903bdd08",
      "attributes": {
        "created": "2020-06-11T00:51:01.553117+03:00",
        "state": "finished"
      }
    }
  ]
}
```

##### Получить одно событие
```
POST /api/v1/event/a23626f1-8db6-4e73-86b2-fc10903bdd08
```

```json
{
  "data": {
    "type": "stream",
    "id": "06d606b6-661a-458f-9553-a7046ef68f1d",
    "attributes": {
      "created": "2020-06-11T01:05:16.935298+03:00",
      "state": "created"
    }
  }
}
```

##### Добавить событие
```
POST /api/v1/event
```

```json
{
  "data": [
    {
      "type": "stream",
      "id": "a23626f1-8db6-4e73-86b2-fc10903bdd08",
      "attributes": {
        "created": "2020-06-11T00:51:01.553117+03:00",
        "state": "created"
      }
    }
  ]
}
```

##### Обновить стейт события

```
PUT /api/v1/event/a23626f1-8db6-4e73-86b2-fc10903bdd08
{
	"state": "interrupted"
}
```
```json
{
  "data": [
    {
      "type": "stream",
      "id": "a23626f1-8db6-4e73-86b2-fc10903bdd08",
      "attributes": {
        "created": "2020-06-11T00:51:01.553117+03:00",
        "state": "interrupted"
      }
    }
  ]
}
```

##### Удалить событие
```
DELETE /api/v1/event/a23626f1-8db6-4e73-86b2-fc10903bdd08
```
```
No Content
```