-- a) All students with contacts + courses they currently take
SELECT s.student_id, s.first_name, s.last_name, s.email, s.address, s.date_of_birth,
       c.course_id, c.name AS course_name
FROM students s
LEFT JOIN enrollments e
  ON e.student_id = s.student_id AND e.deleted_at IS NULL
LEFT JOIN courses c ON c.course_id = e.course_id
ORDER BY s.student_id, c.course_id;

-- b) All courses with department, professors, and enrolled students
SELECT c.course_id, c.name AS course_name, d.name AS department_name,
       p.professor_id, (p.first_name || ' ' || p.last_name) AS professor_name,
       s.student_id, (s.first_name || ' ' || s.last_name) AS student_name
FROM courses c
JOIN departments d ON d.department_id = c.department_id
LEFT JOIN teachings t ON t.course_id = c.course_id
LEFT JOIN professors p ON p.professor_id = t.professor_id
LEFT JOIN enrollments e ON e.course_id = c.course_id AND e.deleted_at IS NULL
LEFT JOIN students s ON s.student_id = e.student_id
ORDER BY c.course_id, p.professor_id, s.student_id;

-- c) All professors with contacts + courses they teach
SELECT p.professor_id, p.first_name, p.last_name, p.email, p.address,
       c.course_id, c.name AS course_name
FROM professors p
LEFT JOIN teachings t ON t.professor_id = p.professor_id
LEFT JOIN courses c ON c.course_id = t.course_id
ORDER BY p.professor_id, c.course_id;

-- d) Enrollment date and course credits per student-course enrollment
SELECT e.enrollment_id, e.student_id, e.course_id,
       e.enrollment_date, c.credits
FROM enrollments e
JOIN courses c ON c.course_id = e.course_id
WHERE e.deleted_at IS NULL
ORDER BY e.enrollment_id;

-- e) Departments with their courses
SELECT d.department_id, d.name AS department_name, c.course_id, c.name AS course_name
FROM departments d
LEFT JOIN courses c ON c.department_id = d.department_id
ORDER BY d.department_id, c.course_id;

-- f) Total students per course (active)
SELECT c.course_id, c.name AS course_name, COUNT(e.enrollment_id) AS total_students
FROM courses c
LEFT JOIN enrollments e
  ON e.course_id = c.course_id AND e.deleted_at IS NULL
GROUP BY c.course_id, c.name
ORDER BY c.course_id;

-- g) Average number of students per course by department
WITH counts AS (
  SELECT c.course_id, c.department_id, COUNT(e.enrollment_id) AS total_students
  FROM courses c
  LEFT JOIN enrollments e
    ON e.course_id = c.course_id AND e.deleted_at IS NULL
  GROUP BY c.course_id, c.department_id
)
SELECT d.department_id, d.name AS department_name,
       AVG(total_students)::NUMERIC(10,2) AS avg_students_per_course
FROM counts
JOIN departments d ON d.department_id = counts.department_id
GROUP BY d.department_id, d.name
ORDER BY d.department_id;