#### A small pet-project about films and actors

#### Swagger
Спецификация полностью описана в файле server/cmd/docs/swagger.yaml

#### Запуск тестов и проверка покрытия в папке internal

```go1.19 test -coverprofile=c.out ./...
   go1.19 tool cover -func c.out | grep total```