CREATE INDEX IF NOT EXISTS mangas_title_idx ON mangas USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS mangas_genres_idx ON mangas USING GIN (genres);