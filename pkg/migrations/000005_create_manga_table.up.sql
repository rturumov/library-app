CREATE TABLE IF NOT EXISTS mangas (
                                     id bigserial PRIMARY KEY,
                                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                     title text NOT NULL,
                                     year integer NOT NULL,
                                     author_id integer NOT NULL,
                                     genres text[] NOT NULL,
                                     version integer NOT NULL DEFAULT 1,
                                     FOREIGN KEY (author_id) REFERENCES authors(id)
    );