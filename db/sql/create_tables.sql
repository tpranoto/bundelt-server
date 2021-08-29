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

-- Creation of product table
CREATE TABLE IF NOT EXISTS product (
  product_id INT NOT NULL,
  name varchar(250) NOT NULL,
  PRIMARY KEY (product_id)
);
