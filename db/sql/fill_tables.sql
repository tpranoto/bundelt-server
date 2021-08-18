-- Filling of users
INSERT INTO users
select id, concat('User ', id) 
FROM GENERATE_SERIES(1, 20) as id;
