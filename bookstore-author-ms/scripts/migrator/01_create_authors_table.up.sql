CREATE TABLE public.authors (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    book_id VARCHAR(64),
    name VARCHAR(64)
);