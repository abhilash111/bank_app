CREATE TABLE "customers"
(
    "customer_id" bigserial NOT NULL PRIMARY KEY,
    "name" varchar NOT NULL,
    "date_of_birth" date NOT NULL,
    "city" varchar NOT NULL,
    "zipcode" varchar NOT NULL,
    "status" int NOT NULL DEFAULT '1'
);

INSERT INTO "customers"
VALUES
    (2000, 'Steve', '1978-12-15', 'Delhi', '110075', 1),
    (2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550', 1),
    (2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', 1),
    (2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', 0),
    (2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', 1),
    (2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', 0);


CREATE TABLE "acconts"
(
    "id" bigserial PRIMARY KEY,
    "owner" VARCHAR NOT NULL,
    "balance" VARCHAR NOT NULL,
    "currency" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);