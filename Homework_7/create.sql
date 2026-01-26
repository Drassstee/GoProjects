CREATE DATABASE University;

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    groupname TEXT NOT NULL
);

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    major TEXT NOT NULL,
    course_year SMALLINT NOT NULL,
    gender CHAR(1) NOT NULL,
    birth_date DATE NOT NULL,
    group_id INTEGER NOT NULL REFERENCES groups(id)
);

CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    subname TEXT NOT NULL
);

CREATE TABLE class_schedule (
    id SERIAL PRIMARY KEY,
    lessons TEXT[] NOT NULL,
    group_id INTEGER NOT NULL REFERENCES groups(id)
);

CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    visit_day DATE NOT NULL,
    visit BOOLEAN NOT NULL,
    student_id INTEGER REFERENCES students(id),
    subject_id INTEGER NOT NULL REFERENCES subjects(id)
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);