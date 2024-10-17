# Система портов
Задача подразумевает взаимодействие с объектами "порт", которые могут быть либо входными (IN), либо выходными (OUT), количество портов задается при запуске.

## Запуск
1. Склонировать репозиторий
```bash
git clone https://github.com/qquiqlerr/test_intelvision.git
cd test_intelvision
```
2. Запустить программу с указанием количества портов IN и OUT
```bash
go run cmd/main.go --in 8 --out 8
```
Сервис будет доступен на порте 8080

## API
### /read/{id}
GET-запрос на получение состояния IN порта по его id

Пример:
```bash
curl http://localhost:8080/read/1
```

Ответ:
```json
{
    "value": 0
}
```

### /write
POST-запрос на запись значения в OUT порт

Пример:
```bash
curl -X POST -d '{"id": 1, "value": 1}' http://localhost:8080/write
```

Ответ:
```json
{
    "status": "ok"
}
```
Значение порта выводится в консоль


