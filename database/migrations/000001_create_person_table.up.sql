-- file: ./migrations/000001_create_person_table.up.sql

CREATE TABLE person(
    id BIGINT NOT NULL,
    first_name VARCHAR(256),
    last_name VARCHAR(256)
)