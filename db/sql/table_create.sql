-- Create db
-- CREATE DATABASE cnsgateway;

-- Create table
DROP TABLE session;
DROP TABLE errorLog;

CREATE TABLE session (
	session_id varchar(100) NOT NULL,
    pk varchar(1000) NOT NULL,
    timestamp datetime NOT NULL    
);

CREATE TABLE errorLog (
    index_id varchar(1000) NOT NULL, 
    message varchar(1000) NOT NULL, 
    timestamp datetime NOT NULL
);