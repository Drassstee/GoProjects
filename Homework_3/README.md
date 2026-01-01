# Homework_3
## Получить информацию о студенте
Чтобы получить данные студента нужно сделать запрос в постмане на:
```
localhost:8080/student/:id
```
Надо заменить ":id" на айди студента. Пример:
```
localhost:8080/student/11
```
Результат:
```JSON
{
    "firstname": "student1",
    "surname": "student1",
    "group_name": "CS-214",
    "major": "Computer Science",
    "course_year": 2,
    "gender": "M",
    "birth_date": "2007-12-10T00:00:00Z"
}
```
## Получить расписания всех групп
Надо просто сделать запрос на:
```
localhost:8080/all_class_schedule
```
Результат:
```JSON
[
    {
        "id": 4,
        "group_name": "ROB-212",
        "schedule": [
            "Calculus 1",
            "CSCI 151",
            "ROBT 206",
            "CSCI 231",
            "MATH 271"
        ]
    },
    {
        "id": 5,
        "group_name": "MATH-212",
        "schedule": [
            "CSCI 151",
            "ROBT 206",
            "CSCI 231",
            "MATH 271",
            "Calculus 1"
        ]
    },
    {
        "id": 1,
        "group_name": "CS-212",
        "schedule": [
            "ROBT 206",
            "CSCI 231",
            "MATH 271",
            "Calculus 1",
            "CSCI 151"
        ]
    },
    {
        "id": 2,
        "group_name": "CS-213",
        "schedule": [
            "CSCI 231",
            "MATH 271",
            "Calculus 1",
            "CSCI 151",
            "ROBT 206"
        ]
    },
    {
        "id": 3,
        "group_name": "CS-214",
        "schedule": [
            "MATH 271",
            "Calculus 1",
            "CSCI 151",
            "ROBT 206",
            "CSCI 231"
        ]
    }
]
```
## Получить расписание конкретной группы
Сделать запрос на:
```
localhost:8080/schedule/group/:id
```
Заменить ":id" на айди группы
```
localhost:8080/schedule/group/3
```
Результат:
```JSON
{
    "id": 3,
    "group_name": "CS-214",
    "schedule": [
        "MATH 271",
        "Calculus 1",
        "CSCI 151",
        "ROBT 206",
        "CSCI 231"
    ]
}
```