CREATE TABLE IF NOT EXISTS images (
	id SERIAL NOT NULL PRIMARY KEY,
	name1 VARCHAR(255),
	image_str text,
	image_type VARCHAR(255),
	created_at TIMESTAMP(0) DEFAULT NOW()
	);
