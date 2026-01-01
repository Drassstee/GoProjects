# Homework_2

## Домашка по фреймворкам

Запуск:
```bash
go run ./file_1
```
в Постмане на тег "/" отправить гет запрос. Результат:
```
Hello, World!
```
Если отправить на тег "/postreq" любой запрос методом кроме POST  
то выйдет ошибка json с статус кодом 405(Method not Allowed)
```JSON
{
    "message": "Method Not Allowed"
}
```
Чтобы получить правильный результат, надо отправить запрос методом пост
  и заполнить данные(name и surname) в виде JSON:
```JSON
{
    "name": "Dayana",
    "surname": "Dayana"
}
```
результат:
```JSON
{
    "name": "Dayana",
    "surname": "Dayana"
}
```
Если не заполнить один из филдов то высветится ошибка с статус кодом 400(Bad Request)
```JSON
{
    "message": "Name and surname cannot be empty"
}
```
## Домашка по sql
Create tables скрипт находится в create.sql.  
Insert и Select в insert_and_select.sql.  
результат селекта в select_query_result.csv.