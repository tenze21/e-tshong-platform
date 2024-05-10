CREATE TABLE seller(
    UserId int NOT NULL,
    FirstName varchar(45) NOT NULL,
    LastName varchar(45) DEFAULT NULL,
    ContactNumber int NOT NULL,
    Email varchar(45) DEFAULT NULL,
    Gender varchar(10) NOT NULL,
    Password varchar(45) NOT NULL,
    PRIMARY KEY (UserId)
)