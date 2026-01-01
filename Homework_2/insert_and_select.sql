INSERT INTO public.groups VALUES (1, 'CS-212');
INSERT INTO public.groups VALUES (2, 'CS-213');
INSERT INTO public.groups VALUES (3, 'CS-214');
INSERT INTO public.groups VALUES (4, 'ROB-212');
INSERT INTO public.groups VALUES (5, 'MATH-212');


INSERT INTO public.class_schedule VALUES (1, 4, '{"Calculus 1","CSCI 151","ROBT 206","CSCI 231","MATH 271"}');
INSERT INTO public.class_schedule VALUES (2, 5, '{"CSCI 151","ROBT 206","CSCI 231","MATH 271","Calculus 1"}');
INSERT INTO public.class_schedule VALUES (3, 1, '{"ROBT 206","CSCI 231","MATH 271","Calculus 1","CSCI 151"}');
INSERT INTO public.class_schedule VALUES (4, 2, '{"CSCI 231","MATH 271","Calculus 1","CSCI 151","ROBT 206"}');
INSERT INTO public.class_schedule VALUES (5, 3, '{"MATH 271","Calculus 1","CSCI 151","ROBT 206","CSCI 231"}');


INSERT INTO public.students VALUES (11, 'student1', 'student1', 'Computer Science', 2, 3, 'M', '2007-12-10');
INSERT INTO public.students VALUES (12, 'student2', 'student2', 'Computer Science', 2, 2, 'F', '2007-10-12');
INSERT INTO public.students VALUES (13, 'student3', 'student3', 'Computer Science', 2, 3, 'F', '2007-08-15');
INSERT INTO public.students VALUES (14, 'student4', 'student4', 'Robotics', 1, 4, 'M', '2007-08-10');
INSERT INTO public.students VALUES (15, 'student5', 'student5', 'Mathematics', 1, 5, 'M', '2006-12-15');
INSERT INTO public.students VALUES (16, 'student6', 'student6', 'Computer Science', 2, 1, 'F', '2007-01-26');
INSERT INTO public.students VALUES (17, 'student7', 'student7', 'Computer Science', 2, 2, 'M', '2006-03-08');
INSERT INTO public.students VALUES (18, 'student8', 'student8', 'Computer Science', 2, 2, 'M', '2005-12-29');
INSERT INTO public.students VALUES (19, 'student9', 'student9', 'Computer Science', 3, 3, 'M', '2007-10-13');
INSERT INTO public.students VALUES (20, 'student10', 'student10', 'Computer Science', 2, 1, 'M', '2006-11-28');

SELECT s.firstname, s.lastname, s.major, s.course_year,s.gender,s.birth_date, c.lessons 
	FROM public.students AS s
	JOIN public.class_schedule AS c
	ON s.group_id = c.group_id
	WHERE s.gender = 'F'
	ORDER BY s.birth_date DESC;

