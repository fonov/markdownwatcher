 CREATE TABLE items
 (
   id VARCHAR(36) PRIMARY KEY NOT NULL,
   title VARCHAR(100) NOT NULL,
   url VARCHAR(50) NOT NULL,
   description VARCHAR(100),
   price int NOT NULL,
   oldPrice int
 );
 CREATE UNIQUE INDEX sqlite_autoindex_items ON items (id);