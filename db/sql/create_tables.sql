--Create user table
CREATE TABLE IF NOT EXISTS users (
  user_id BIGSERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  msisdn VARCHAR(50) NOT NULL,
  pwd TEXT NOT NULL,
  user_loc JSONB NOT NULL,
  full_name VARCHAR(50),
  sex INT,
  last_login TIMESTAMP,
  create_time TIMESTAMP,
  update_time TIMESTAMP
);

-- Creation of user group rel table
CREATE TABLE IF NOT EXISTS user_group_rel (
  user_fb_id VARCHAR(60) NOT NULL,
  group_id VARCHAR(60) NOT NULL,
  PRIMARY KEY (user_fb_id, group_id)
);

-- Creation of group table
CREATE TABLE IF NOT EXISTS groups(
  group_id VARCHAR(60) NOT NULL PRIMARY KEY,
  group_name VARCHAR(60) NOT NULL,
  group_desc TEXT NOT NULL,
  created VARCHAR(30) NOT NULL,
  lat DOUBLE PRECISION,
  lon DOUBLE PRECISION
);
