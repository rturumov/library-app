CREATE INDEX IF NOT EXISTS mangas_title_idx ON mangas USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS mangas_genres_idx ON mangas USING GIN (genres);

-- Вставка данных в таблицу authors
INSERT INTO authors (name) VALUES
                               ('Author One'),
                               ('Author Two'),
                               ('Author Three'),
                               ('Author Four'),
                               ('Author Five'),
                               ('Author Six'),
                               ('Author Seven'),
                               ('Author Eight'),
                               ('Author Nine'),
                               ('Author Ten');

-- Вставка данных в таблицу books
INSERT INTO books (title, year, author_id, genres) VALUES
                                                       ('Book One', 2020, 1, ARRAY['Fiction', 'Adventure']),
                                                       ('Book Two', 2019, 2, ARRAY['Non-Fiction', 'Biography']),
                                                       ('Book Three', 2021, 3, ARRAY['Science Fiction', 'Thriller']),
                                                       ('Book Four', 2018, 4, ARRAY['Fantasy', 'Drama']),
                                                       ('Book Five', 2017, 5, ARRAY['Mystery', 'Horror']),
                                                       ('Book Six', 2022, 6, ARRAY['Romance', 'Comedy']),
                                                       ('Book Seven', 2023, 7, ARRAY['Historical', 'Adventure']),
                                                       ('Book Eight', 2021, 8, ARRAY['Science', 'Technology']),
                                                       ('Book Nine', 2016, 9, ARRAY['Education', 'Self-help']),
                                                       ('Book Ten', 2015, 10, ARRAY['Children', 'Fantasy']);

-- Вставка данных в таблицу mangas
INSERT INTO mangas (title, year, author_id, genres) VALUES
                                                        ('Manga One', 2020, 1, ARRAY['Action', 'Adventure']),
                                                        ('Manga Two', 2019, 2, ARRAY['Fantasy', 'Romance']),
                                                        ('Manga Three', 2021, 3, ARRAY['Horror', 'Mystery']),
                                                        ('Manga Four', 2018, 4, ARRAY['Comedy', 'Drama']),
                                                        ('Manga Five', 2017, 5, ARRAY['Science Fiction', 'Thriller']),
                                                        ('Manga Six', 2022, 6, ARRAY['Fantasy', 'Action']),
                                                        ('Manga Seven', 2023, 7, ARRAY['Romance', 'Slice of Life']),
                                                        ('Manga Eight', 2021, 8, ARRAY['Adventure', 'Fantasy']),
                                                        ('Manga Nine', 2016, 9, ARRAY['Historical', 'Drama']),
                                                        ('Manga Ten', 2015, 10, ARRAY['Supernatural', 'Mystery']);