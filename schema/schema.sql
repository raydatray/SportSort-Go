CREATE TABLE sport_centers (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL
);

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  type USER_TYPE NOT NULL,
  sport_center_id INT,
  deleted BOOL NOT NULL,
  FOREIGN KEY (sport_center_id) REFERENCES sport_centers(id)
);

CREATE TABLE sessions (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  session_token TEXT UNIQUE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TYPE user_type AS ENUM ('admin', 'owner', 'instructor', 'customer' );

CREATE TABLE rooms (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  capacity BIGINT NOT NULL,
  sport_center_id INT NOT NULL,
  FOREIGN KEY (sport_center_id) REFERENCES sport_centers(id)
);

CREATE TABLE course_types (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  rate MONEY NOT NULL
);

CREATE TABLE course_offerings (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  starting_date DATE NOT NULL,
  ending_date DATE NOT NULL,
  price MONEY NOT NULL,
  sport_center_id INT NOT NULL,
  course_type_id BIGINT NOT NULL,
  instructor_id BIGINT NOT NULL,
  FOREIGN KEY (sport_center_id) REFERENCES sport_centers(id),
  FOREIGN KEY (course_type_id) REFERENCES course_types(id),
  FOREIGN KEY (instructor_id) REFERENCES users(id)
);

CREATE TABLE course_sessions (
  id BIGSERIAL PRIMARY KEY,
  date DATE NOT NULL,
  course_offering_id BIGINT NOT NULL,
  room_id BIGINT NOT NULL,
  instructor_id BIGINT NOT NULL,
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  FOREIGN KEY (course_offering_id) REFERENCES course_offerings(id),
  FOREIGN KEY (room_id) REFERENCES rooms(id),
  FOREIGN KEY (instructor_id) REFERENCES users(id)
);

CREATE TABLE registrations (
  user_id BIGINT NOT NULL,
  course_offering_id BIGINT NOT NULL,
  registration_date TIMESTAMPTZ NOT NULL,
  price_paid MONEY NOT NULL,
  PRIMARY KEY (user_id, course_offering_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (course_offering_id) REFERENCES course_offerings(id)
);

CREATE TABLE reviews (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  rating INTEGER NOT NULL,
  comment VARCHAR,
  review_date TIMESTAMPTZ NOT NULL,
  instructor_id BIGINT,
  course_type_id BIGINT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (instructor_id) REFERENCES users(id),
  FOREIGN KEY (course_type_id) REFERENCES course_types(id),
  CHECK (
    (instructor_id IS NOT NULL AND course_type_id IS NULL) OR
    (instructor_id IS NULL AND course_type_id IS NOT NULL)
  ),
  CHECK (rating >= 1 AND rating <= 10)
);
