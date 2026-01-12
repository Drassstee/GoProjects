# Homework_4
## Получить информацию о посещении студента
```
/attendanceByStudentId/:id
```
Пример:
```
/attendanceByStudentId/11
```
Результат:
```JSON
[
    {
        "subject_name": "ROBT 206",
        "subject_id": 2,
        "visit_day": "2026-01-16T00:00:00Z",
        "visited": false,
        "student_firstname": "student1",
        "student_surname": "student1",
        "student_id": 11
    },
    {
        "subject_name": "MATH 321",
        "subject_id": 3,
        "visit_day": "2026-01-12T00:00:00Z",
        "visited": true,
        "student_firstname": "student1",
        "student_surname": "student1",
        "student_id": 11
    }
]
```
## Получить информацию о посещении cтудентов по конкретной дисциплине
Надо просто сделать запрос на:
```
/attendanceBySubjectId/:id
```
Пример:
```
/attendanceBySubjectId/2
```
Результат:
```JSON
[
    {
        "subject_name": "ROBT 206",
        "subject_id": 2,
        "visit_day": "2026-01-16T00:00:00Z",
        "visited": false,
        "student_firstname": "student1",
        "student_surname": "student1",
        "student_id": 11
    },
    {
        "subject_name": "ROBT 206",
        "subject_id": 2,
        "visit_day": "2026-01-16T00:00:00Z",
        "visited": true,
        "student_firstname": "student5",
        "student_surname": "student5",
        "student_id": 15
    }
]
```
## Добавить запись посещения
Сделать POST запрос на:
```
/attendance/subject
```
с json:
```JSON
{
"subject_id": 3,
"visit_day": "2026-01-07T00:00:00Z",
"visited": true,
"student_id": 11
}
```
Результат:
```JSON
"success"
```