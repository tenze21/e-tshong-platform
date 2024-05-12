-- Seller Table.
CREATE TABLE seller(
    UserId SERIAL PRIMARY KEY,
    FirstName varchar(45) NOT NULL,
    LastName varchar(45) DEFAULT NULL,
    ContactNumber int UNIQUE NOT NULL,
    Email varchar(45) UNIQUE DEFAULT NULL,
    Gender varchar(10) NOT NULL,
    Password varchar(45) NOT NULL
);

INSERT INTO seller(FirstName, LastName, ContactNumber, Email, Gender, Password) VALUES('tenzin', 'choda', 17736762, 'tenzin@33.com', 'Male', 'Nov@core2023');
DELETE FROM seller WHERE UserId=1;
DROP TABLE seller;

-- Table for User profile picture.
CREATE TABLE seller_profile(
    picture_id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    profile_picture BYTEA,
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES seller(UserId) ON DELETE CASCADE ON UPDATE CASCADE
);