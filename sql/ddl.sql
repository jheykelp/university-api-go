-- ddl.sql (PostgreSQL)

CREATE TABLE departments (
  department_id SERIAL PRIMARY KEY,
  name          VARCHAR(120) NOT NULL,
  description   TEXT NOT NULL
);

CREATE TABLE students (
  student_id     SERIAL PRIMARY KEY,
  first_name     VARCHAR(80)  NOT NULL,
  last_name      VARCHAR(80)  NOT NULL,
  email          VARCHAR(160) NOT NULL UNIQUE,
  address        TEXT         NOT NULL,
  date_of_birth  DATE         NOT NULL,
  password_hash  TEXT         NOT NULL,
  created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE professors (
  professor_id  SERIAL PRIMARY KEY,
  first_name    VARCHAR(80)  NOT NULL,
  last_name     VARCHAR(80)  NOT NULL,
  email         VARCHAR(160) NOT NULL UNIQUE,
  address       TEXT         NOT NULL,
  created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE courses (
  course_id     SERIAL PRIMARY KEY,
  name          VARCHAR(160) NOT NULL,
  description   TEXT         NOT NULL,
  department_id INT          NOT NULL REFERENCES departments(department_id),
  credits       SMALLINT     NOT NULL CHECK (credits >= 0)
);

CREATE TABLE enrollments (
  enrollment_id  SERIAL PRIMARY KEY,
  student_id     INT NOT NULL REFERENCES students(student_id),
  course_id      INT NOT NULL REFERENCES courses(course_id),
  enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE,
  deleted_at     TIMESTAMPTZ NULL
);

-- Prevent duplicate active enrollments; allow future re-enroll if previous was soft-deleted
CREATE UNIQUE INDEX uniq_active_enrollment
  ON enrollments(student_id, course_id)
  WHERE deleted_at IS NULL;

CREATE TABLE teachings (
  professor_id INT NOT NULL REFERENCES professors(professor_id),
  course_id    INT NOT NULL REFERENCES courses(course_id),
  PRIMARY KEY (professor_id, course_id)
);

-- Seed data (minimal, expand as needed)
INSERT INTO departments (name, description) VALUES
('Computer Science','CS Dept'), ('Mathematics','Math Dept');

INSERT INTO courses (name, description, department_id, credits) VALUES
('Algorithms','Design and analysis', 1, 3),
('Databases','Relational theory and SQL', 1, 3),
('Calculus I','Differential calculus', 2, 4);

INSERT INTO professors (first_name,last_name,email,address) VALUES
('Alan','Turing','alan.turing@uni.test','Princeton, NJ'),
('Edsger','Dijkstra','edsger.dijkstra@uni.test','Austin, TX');

INSERT INTO teachings (professor_id, course_id) VALUES
(1,1),(2,2);

-- create a demo student (password: "secret123" -> replace with hash in app or pre-hash)
INSERT INTO students (first_name,last_name,email,address,date_of_birth,password_hash)
VALUES ('Jane','Doe','jane.doe@uni.test','Dorm A','2003-05-20','TO_BE_SET_FROM_APP');