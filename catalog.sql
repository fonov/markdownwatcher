PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "items"
(
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL,
    url VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    price int NOT NULL,
    oldPrice int NOT NULL
);
CREATE UNIQUE INDEX items_id_uindex ON "items" (id);
COMMIT;
