# URL Shortener

Сервис для сокращения ссылок, написан на Go

## Возможности

- Сокращение URL и редирект
- Хранение ссылок в MySQL или в памяти
- Асинхронное логирование запросов
- REST API
- Unit-тесты

## Запуск

1. git clone https://github.com/Bastion-Scar/monitoring-demo
2. cd monitoring-demo/infrastructure
3. docker compose build
4. ../setup.sh 

## Сервисы: 

1. Пишем в терминале curl -X POST -H 'Content-Type: application/json' -d '{"url": "https://www.github.com"}' http://localhost/shorten (ссылочку можно любую, я взял гитхаб дот ком для примера типа крутой (я же крутой? нет? ладно))
2. После получения кода пишем в адресной строке localhost/код и вас перенесет на указанный ранее сайт
3. На адресе localhost:9090 можно посмотреть метрики