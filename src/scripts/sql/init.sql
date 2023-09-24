CREATE TABLE IF NOT EXISTS curators
(
    id           INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    chat_id      TEXT NOT NULL,
    name         TEXT NOT NULL,
    surname      TEXT NOT NULL CHECK (surname != ''),
    phone_number TEXT NOT NULL CHECK (phone_number != '')
);

CREATE TABLE IF NOT EXISTS shelters
(
    id     INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    street TEXT NOT NULL CHECK (street != ''),
    house  INT NOT NULL CHECK (house > 0)
);

CREATE TABLE IF NOT EXISTS animals
(
    id         INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name       TEXT NOT NULL CHECK (name != ''),
    age        INT NOT NULL CHECK (age BETWEEN 1 AND 40),
    height     DOUBLE PRECISION NOT NULL,
    weight     DOUBLE PRECISION NOT NULL,
    shelter_id INT NOT NULL,
    type       TEXT NOT NULL,
    gender     TEXT NOT NULL,

    FOREIGN KEY (shelter_id) REFERENCES shelters(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS purchases
(
    id        INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name      TEXT NOT NULL CHECK (name != ''),
    frequency TEXT,
    cost      DOUBLE PRECISION NOT NULL,
    last_date DATE NOT NULL,
    animal_id INT NOT NULL,

    FOREIGN KEY (animal_id) REFERENCES animals(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS curators_animals
(
    id         INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    animal_id  INT NOT NULL,
    curator_id INT NOT NULL,

    FOREIGN KEY (animal_id) REFERENCES animals(id) ON DELETE CASCADE,
    FOREIGN KEY (curator_id) REFERENCES curators(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS diseases
(
    id         INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    diagnosis  TEXT NOT NULL CHECK (diagnosis != ''),
    symptoms   TEXT NOT NULL,
    cause      TEXT,
    is_chronic BOOLEAN NOT NULL,
    animal_id  INT NOT NULL,

    FOREIGN KEY (animal_id) REFERENCES animals(id) ON DELETE CASCADE
);