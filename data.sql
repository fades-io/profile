DROP TABLE IF EXISTS public.user CASCADE;

-- init DATA

CREATE TABLE public.user
(
    id   VARCHAR(100),
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    level int NOT NULL,
    daysinrow int NOT NULL,
    daysinweek json NOT NULL,
    doessendpushups boolean NOT NULL,
    theme VARCHAR(100) NOT NULL,
    language VARCHAR(100) NOT NULL,
    image VARCHAR(100) NOT NULL
);


INSERT INTO public.user (id, username, email, password, level, daysinrow, daysinweek, doessendpushups, theme, language, image)
VALUES ('707f69e0-edac-4c3e-bb7f-918d3f190e19','u1', 'mail1', 'p1', 11, 1, ' { "first" : "John" , "middle" : "K", "last" : "Doe" }', false, 'theme1,', 'lang1', 'http1');
INSERT INTO public.user (id, username, email, password, level, daysinrow, daysinweek, doessendpushups, theme, language, image)
VALUES ('1ad0596d-e43d-4093-a7fe-a6c1074f6219','u2', 'mail2', 'p2', 22, 2, ' { "first" : "John" , "middle" : "K", "last" : "Doe" }', false, 'theme2,', 'lang2', 'http2');
INSERT INTO public.user (id, username, email, password, level, daysinrow, daysinweek, doessendpushups, theme, language, image)
VALUES ('62af3986-0963-465c-8a86-dd23ac240523','u3', 'mail3', 'p3', 33, 3, ' { "first" : "John" , "middle" : "K", "last" : "Doe" }' , false, 'theme3,', 'lang3', 'http3');
