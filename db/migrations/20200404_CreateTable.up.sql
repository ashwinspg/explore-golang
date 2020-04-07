CREATE TABLE IF NOT EXISTS movie_info(
	movie_uuid UUID,
	info JSON NOT NULL,
	PRIMARY KEY (movie_uuid)
);