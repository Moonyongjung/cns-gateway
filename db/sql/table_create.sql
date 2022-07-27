-- Create db
-- CREATE DATABASE cnsgateway;

-- Create table
DROP TABLE session;

CREATE TABLE session (
	session_id varchar(100) NOT NULL,
    pk varchar(1000) NOT NULL,
    timestamp datetime NOT NULL    
);