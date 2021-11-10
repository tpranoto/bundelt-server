-- table for user information
CREATE TABLE IF NOT EXISTS users (
  user_id BIGSERIAL PRIMARY KEY,
  email VARCHAR(50) NOT NULL UNIQUE,
  msisdn VARCHAR(50) NOT NULL,
  full_name VARCHAR(50) NOT NULL,
  pwd TEXT NOT NULL,
  lat DOUBLE PRECISION,
  lon DOUBLE PRECISION,
  last_login TIMESTAMP,
  create_time TIMESTAMP,
  update_time TIMESTAMP
);

-- table for relation between user and group
CREATE TABLE IF NOT EXISTS user_group_rel (
  user_id BIGINT NOT NULL,
  group_id BIGINT NOT NULL,
  role INT DEFAULT 0 NOT NULL, -- 0 member, 1 admin
  PRIMARY KEY (user_id, group_id)
);

-- table for group information
CREATE TABLE IF NOT EXISTS groups(
  group_id BIGSERIAL PRIMARY KEY,
  group_name VARCHAR(60) NOT NULL,
  group_desc TEXT NOT NULL,
  created VARCHAR(30) NOT NULL,
  lat DOUBLE PRECISION,
  lon DOUBLE PRECISION
);

-- table for group messages
CREATE TABLE IF NOT EXISTS group_messages(
  group_msg_id BIGSERIAL PRIMARY KEY,
  group_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  msg TEXT NOT NULL,
  create_time TIMESTAMP
);

-- table for events
CREATE TABLE IF NOT EXISTS events(
  event_id BIGSERIAL PRIMARY KEY,
  event_name VARCHAR(100) NOT NULL,
  event_desc TEXT,
  --visibility INT NOT NULL,
  --status INT NOT NULL,
  start_time TIMESTAMP,
  end_time TIMESTAMP,
  lat DOUBLE PRECISION,
  lon DOUBLE PRECISION
);

-- table for event group rel
CREATE TABLE IF NOT EXISTS event_group_rel(
  group_id BIGINT NOT NULL,
  event_id BIGINT NOT NULL,
  PRIMARY KEY (group_id, event_id)
);

-- table for event user rel
CREATE TABLE IF NOT EXISTS event_user_rel(
  event_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  PRIMARY KEY (event_id, user_id)
);

-- table for relation between users
CREATE TABLE IF NOT EXISTS user_rel(
  user_rel_id BIGSERIAL PRIMARY KEY,
  user_id1 BIGINT NOT NULL,
  user_id2 BIGINT NOT NULL,
  create_time TIMESTAMP
);

