INSERT INTO curators(chat_id, name, surname, phone_number)
VALUES ('85085228', 'Алина', 'Сабирова', '+79031751212'),
       ('85333221', 'Полина', 'Колчина', '+79031753355'),
       ('81235233', 'Вера', 'Иванова', '+79045651212'),
       ('45685245', 'Мария', 'Отмахова', '+79121751212'),
       ('32385228', 'Алексей', 'Сабиров', '+79031753245');

SELECT *
FROM curators;

INSERT INTO shelters(street, house)
VALUES ('Рябинина', 12),
       ('Анникова', 3),
       ('Мира', 44),
       ('Анцифирова', 10),
       ('Баумана', 10);

SELECT *
FROM shelters;

INSERT INTO animals(name, age, height, weight, shelter_id, type, gender)
VALUES ('Шарик', 4, 3.3, 1.1, 1, 'собака', 'мужской'),
       ('Бобик', 1, 2.3, 0.4, 1, 'собака', 'мужской'),
       ('Мурка', 1, 2.2, 0.1, 3, 'кошка', 'женский'),
       ('Ириска', 11, 2.3, 0.2, 4, 'кошка', 'женский'),
       ('Кэт', 5, 2.4, 0.1, 4, 'кошка', 'женский');

SELECT *
FROM animals;

INSERT INTO purchases(name, frequency, cost, last_date, animal_id)
VALUES ('Royal Canin Sensible 33', 'раз в 2 недели', 403, '2023-06-15', 3),
       ('Royal Canin Sterilised в желе', 'раз в 2 недели', 800, '2023-06-15', 4),
       ('Royal Canin puppy', 'раз в неделю', 503, '2023-06-20', 2),
       ('Royal Canin Sensible 33', 'раз в 2 недели', 403, '2023-06-17', 5),
       ('Royal Canin medium', 'раз в 2 недели', 600, '2023-06-15', 1),
       ('pi-pi bent наполнитель', 'раз в неделю', 600, '2023-06-22', 4);

SELECT *
FROM purchases;

INSERT INTO curators_animals(animal_id, curator_id)
VALUES (1, 5),
       (2, 4),
       (3, 3),
       (4, 3),
       (4, 2),
       (5, 1);

SELECT *
FROM curators_animals;