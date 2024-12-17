-- Таблица аккаунтов
CREATE TABLE "Account" (
      "email" VARCHAR PRIMARY KEY,
      "password" VARCHAR,
      "name" VARCHAR,
      "role" INTEGER NOT NULL REFERENCES "Role"("id")
);

-- Таблица ролей
CREATE TABLE "Role" (
    "id" serial PRIMARY KEY,
    "name" VARCHAR
);


