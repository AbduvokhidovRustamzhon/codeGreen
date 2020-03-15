CREATE TABLE burgers (
   id BIGSERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   price INT NOT NULL CHECK ( price > 0 ),
   -- don't remove any data from db
   phone    TEXT NOT NULL,
   removed BOOLEAN DEFAULT FALSE,
   fileName TEXT NOT NULL
);


drop table burgers;