CREATE TABLE IF NOT EXISTS public.groups
(
    id integer NOT NULL DEFAULT nextval('groups_id_seq'::regclass),
    groupname text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT groups_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.students
(
    id integer NOT NULL DEFAULT nextval('students_id_seq'::regclass),
    firstname text COLLATE pg_catalog."default" NOT NULL,
    lastname text COLLATE pg_catalog."default" NOT NULL,
    major text COLLATE pg_catalog."default" NOT NULL,
    course_year smallint NOT NULL,
    group_id integer NOT NULL,
    gender character(1) COLLATE pg_catalog."default" NOT NULL,
    birth_date date NOT NULL,
    CONSTRAINT students_pkey PRIMARY KEY (id),
    CONSTRAINT students_groupid_fkey FOREIGN KEY (group_id)
        REFERENCES public.groups (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS public.class_schedule
(
    id integer NOT NULL DEFAULT nextval('class_schedule_id_seq'::regclass),
    group_id integer NOT NULL,
    lessons text[] COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT class_schedule_pkey PRIMARY KEY (id),
    CONSTRAINT class_schedule_group_id_fkey FOREIGN KEY (group_id)
        REFERENCES public.groups (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

