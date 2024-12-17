-- Заплнить таблицу "Role"
INSERT INTO "Role"(id, name)
VALUES (0, 'Пользователь'),
       (1, 'Администратор');

-- Заполнить таблицу "Account"
INSERT INTO "Account"(email, password, name, role)
VALUES ('user@test.ru', '$2a$10$QkjvoLbAM3bDlCgDqu/G4eMfdu0FcLAPSXj4OjwKBRXC79jiJaMtO', 'Иванов Иван Иванович', 0),
       ('admin@test.ru', '$2a$10$QkjvoLbAM3bDlCgDqu/q4eMfdu0FcLAPSXj4OjwKBRXC79jiJaMtO', 'Петров Петр Петрович', 1);