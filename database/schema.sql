CREATE TABLE musicians (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  musician_type TEXT NOT NULL
);

CREATE TABLE albums (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  release_date DATE NOT NULL,
  genre TEXT,
  price REAL NOT NULL,
  description TEXT
);

CREATE TABLE album_musicians (
  album_id INTEGER,
  musician_id INTEGER,
  FOREIGN KEY (album_id) REFERENCES albums(id),
  FOREIGN KEY (musician_id) REFERENCES musicians(id),
  PRIMARY KEY (album_id, musician_id)
);
