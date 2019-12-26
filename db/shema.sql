-- Drop tables.
-- DROP TABLE IF EXISTS hostels CASCADE;
-- DROP TABLE IF EXISTS specialities CASCADE;
-- DROP TABLE IF EXISTS students CASCADE;

--Create tables.
CREATE TABLE hostels
(
  id   SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);


CREATE TABLE specialities
(
   id   SERIAL PRIMARY KEY,
  name  VARCHAR(50) NOT NULL
);
CREATE TABLE students
(
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(50) NOT NULL,
  hostelid      INTEGER,
  specialityid  INTEGER,
  FOREIGN KEY (hostelid) REFERENCES hostels(id),
  FOREIGN KEY (specialityid) REFERENCES specialities(id)
);


INSERT INTO hostels (name) VALUES ('1');
INSERT INTO hostels (name) VALUES ('2');
INSERT INTO hostels (name) VALUES ('3');
INSERT INTO hostels (name) VALUES ('4');
INSERT INTO specialities (name) VALUES ('biology');
INSERT INTO specialities (name) VALUES ('literature');
INSERT INTO specialities (name) VALUES ('computerScience');

INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Max',
  (SELECT id FROM hostels WHERE name = '1'),
  (SELECT id FROM specialities WHERE name = 'biology')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Ihor',
  (SELECT id FROM hostels WHERE name = '1'),
  (SELECT id FROM specialities WHERE name = 'biology')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Stas',
  (SELECT id FROM hostels WHERE name = '2'),
  (SELECT id FROM specialities WHERE name = 'biology')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Masha',
  (SELECT id FROM hostels WHERE name = '3'),
  (SELECT id FROM specialities WHERE name = 'literature')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Daniel',
  (SELECT id FROM hostels WHERE name = '2'),
  (SELECT id FROM specialities WHERE name = 'computerScience')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Sergey',
  (SELECT id FROM hostels WHERE name = '3'),
  (SELECT id FROM specialities WHERE name = 'literature')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Andrey',
  (SELECT id FROM hostels WHERE name = '3'),
  (SELECT id FROM specialities WHERE name = 'computerScience')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Misha',
  (SELECT id FROM hostels WHERE name = '4'),
  (SELECT id FROM specialities WHERE name = 'biology')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Katya',
  (SELECT id FROM hostels WHERE name = '1'),
  (SELECT id FROM specialities WHERE name = 'literature')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Anna',
  (SELECT id FROM hostels WHERE name = '4'),
  (SELECT id FROM specialities WHERE name = 'computerScience')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Orest',
  (SELECT id FROM hostels WHERE name = '1'),
  (SELECT id FROM specialities WHERE name = 'literature')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Anna',
  (SELECT id FROM hostels WHERE name = '4'),
  (SELECT id FROM specialities WHERE name = 'biology')
  );
INSERT INTO students
  (name,hostelid,specialityid)
VALUES
  ('Max',
  (SELECT id FROM hostels WHERE name = '3'),
  (SELECT id FROM specialities WHERE name = 'computerScience')
  );
