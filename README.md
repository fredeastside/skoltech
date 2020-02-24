# Skoltech

Тестовое задание

## Запуск проекта

Требуется установленный и запущенный docker. Выполнить команды:

```bash
git clone https://github.com/fredeastside/skoltech.git ./skoltech
cd skoltech
make start
```

## Описание

В сборку входит пять контейнеров. Запросы необходимо посылать контейнеру skoltech_api. Контейнер skoltech_worker выступает в роли консьюмера. Контейнер skoltech_partner выступает в роли партнера и пишет принимаемые запросы в stdout (Данный контейнер добавлен просто для удобства, для замены партнерского url необходимо изменить переменные среды в файле .env).

Порты контейнеров и партнерский урл можно изменить в файле .env. После изменения необходимо пересобрать и перезапустить сборку. Для этого нужно выполнить команды

```bash
make down
make build
```

## Пример

После отправки запроса вида

```bash
curl -XPOST '0.0.0.0:8080' -H "Content-Type: application/json" -d '{"ap_id":"A8-F9-4B-B6-87-FF","version":"1.0","probe_requests":[{"mac":"88-1D-FC-DF-6F-C1","timestamp":"1579782767"},{"mac":"F8-59-71-PK-95-36","bssid":"04-BF-6D-04-09-8C","ssid":"SKOLTECH"},{"mac":"F8-59-71-PK-95-BB"}]}' -v
```

И выполнив

```bash
make logs
```

Можно увидеть логи движения запроса:

```bash
api_1        | 2020/02/24 13:45:36 Write message to kafka: {"uid":"21a64b48-5ac7-4eb9-b14a-c77d3282d06d","ap_id":"A11-F9-4B-B6-87-FF","version":"1.0","probe_requests":[{"mac":"88-1D-FC-DF-6F-C1","timestamp":"1579782767"},{"mac":"F8-59-71-PK-95-36","bssid":"04-BF-6D-04-09-8C","ssid":"SKOLTECH"},{"mac":"F8-59-71-PK-95-BB"}]}

worker_1     | 2020/02/24 13:45:41 Read message from kafka: {"uid":"21a64b48-5ac7-4eb9-b14a-c77d3282d06d","ap_id":"A11-F9-4B-B6-87-FF","version":"1.0","probe_requests":[{"mac":"88-1D-FC-DF-6F-C1","timestamp":"1579782767"},{"mac":"F8-59-71-PK-95-36","bssid":"04-BF-6D-04-09-8C","ssid":"SKOLTECH"},{"mac":"F8-59-71-PK-95-BB"}]}

partner_1    | 2020/02/24 13:45:41 Get request with body: {"uid":"21a64b48-5ac7-4eb9-b14a-c77d3282d06d","ap_id":"A11-F9-4B-B6-87-FF","version":"1.0","probe_requests":[{"mac":"88-1D-FC-DF-6F-C1","timestamp":"1579782767","bssid":"FF-FF-FF-FF-FF-FF","ssid":"Unknown"},{"mac":"F8-59-71-PK-95-36","bssid":"04-BF-6D-04-09-8C","ssid":"SKOLTECH"},{"mac":"F8-59-71-PK-95-BB","bssid":"FF-FF-FF-FF-FF-FF","ssid":"Unknown"}]}
```

А так же его ошибки:

```bash
curl -XPOST '0.0.0.0:8080' -H "Content-Type: application/json" -d '{"version":"1.0","probe_requests":[{"mac":"88-1D-FC-DF-6F-C1","timestamp":"1579782767"},{"mac":"F8-59-71-PK-95-36","bssid":"04-BF-6D-04-09-8C","ssid":"SKOLTECH"},{"mac":"F8-59-71-PK-95-BB"}]}' -v
```

```bash
api_1        | 2020/02/24 13:48:43 Write message to kafka: {"uid":"fc622d67-e304-4bbb-a384-8063eaa7712b","ap_id":"","version":"1.0","probe_requests":[{"mac":"88-1D-FC-DF-6F-C1","timestamp":"1579782767"},{"mac":"F8-59-71-PK-95-36","bssid":"04-BF-6D-04-09-8C","ssid":"SKOLTECH"},{"mac":"F8-59-71-PK-95-BB"}]}

worker_1     | 2020/02/24 13:48:50 Read message from kafka: {"uid":"fc622d67-e304-4bbb-a384-8063eaa7712b","ap_id":"","version":"1.0","probe_requests":[{"mac":"88-1D-FC-DF-6F-C1","timestamp":"1579782767"},{"mac":"F8-59-71-PK-95-36","bssid":"04-BF-6D-04-09-8C","ssid":"SKOLTECH"},{"mac":"F8-59-71-PK-95-BB"}]}

worker_1     | 2020/02/24 13:48:50 ERROR: device has empty app id
```

## Тестирование

Требуется установленный go

```bash
go test -v ./tests
```

## Улучшения

1. Компрессия данных
2. Строгая валидация данных
3. Масштабируемость консьюмеров (kafka group)
4. Увеличение test-coverage
5. Мониторинг
