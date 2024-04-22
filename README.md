# greenApiTestTask

Этот проект состоит из двух сервисов, M1 и M2, которые предоставляют различные функции. Ниже приведены инструкции по установке, запуску и тестированию проекта.

## Подготовка

Перед запуском проекта необходимо установить RabbitMQ и запустить его. Мы рекомендуем использовать Docker для этого:

```bash
docker pull rabbitmq
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq
```

## Запуск
Перед запуском проекта надо зайти в корневую директорую проекта т.е. greenApiTestTask
```bash
cd internal/services/m1
go run cmd/main.go
```

Затем открыть вторую вкладку в терминале 
```bash
cd internal/services/m2
go run cmd/main.go
```

Проверить что все работает, можете отправить POST запроc по пути "http://localhost:8080/api/v1/double" положив в тело 3 по ключу "param"
