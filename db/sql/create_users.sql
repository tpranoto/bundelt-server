--Create user for users table
CREATE USER bd_user WITH PASSWORD 'password';
GRANT CONNECT ON DATABASE bundelt TO bd_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO bd_user;
GRANT USAGE, SELECT ON SEQUENCE users_user_id_seq TO bd_user;
GRANT USAGE, SELECT ON SEQUENCE groups_group_id_seq TO bd_user;
GRANT USAGE, SELECT ON SEQUENCE group_messages_group_msg_id_seq TO bd_user;
GRANT USAGE, SELECT ON SEQUENCE events_event_id_seq TO bd_user;
