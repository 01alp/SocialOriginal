CREATE TABLE post (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  author INTEGER NOT NULL,
  message_ TEXT NOT NULL,
  image_ TEXT NOT NULL,
  created_at DATETIME NOT NULL,
  privacy INTEGER NOT NULL DEFAULT 0,
  FOREIGN KEY (author) REFERENCES user(id)
);
