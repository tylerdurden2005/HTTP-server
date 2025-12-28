# HTTP-server

Для запуска проекта необходимо:
```shell
1) Ввести команду: git clone https://github.com/tylerdurden2005/HTTP-server.git
2) Ввести в терминале: make или make run
```

Для запуска unit-тестов:
```shell
Ввести в терминале: make test
```

# Обрабатывает эндпоинты:
POST /todos — создать новую задачу
GET /todos — получить список всех задач
GET /todos/{id} — получить задачу по идентификатору
PUT /todos/{id} — обновить задачу по идентификатору
DELETE /todos/{id} — удалить задачу по идентификатору
DELETE /todos - удалить все задачи
