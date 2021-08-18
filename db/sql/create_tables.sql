--Create user table
CREATE TABLE IF NOT EXISTS users (
  user_id INT NOT NULL,
  name varchar(250) NOT NULL,
  PRIMARY KEY (user_id)
);

-- Creation of product table
CREATE TABLE IF NOT EXISTS product (
  product_id INT NOT NULL,
  name varchar(250) NOT NULL,
  PRIMARY KEY (product_id)
);
