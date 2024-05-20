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
    CREATE TABLE seller_profile(
    picture_id SERIAL PRIMARY KEY,
    contact_number int NOT NULL,
    profile_picture BYTEA,
    CONSTRAINT cnumber_fk FOREIGN KEY (contact_number) REFERENCES seller(contactnumber) ON DELETE CASCADE ON UPDATE CASCADE
)
);

CREATE TABLE product(
    product_id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    product_img1 BYTEA,
    product_img2 BYTEA,
    product_img3 BYTEA,
    product_img4 BYTEA,
    product_title varchar(300) NOT NULL,
    actual_price int DEFAULT NULL,
    selling_price int NOT NULL,
    category varchar(45) NOT NULL,
    product_description varchar(500) NOT NULL,
    CONSTRAINT seller_fk FOREIGN KEY (user_id) REFERENCES seller(UserId) ON DELETE CASCADE ON UPDATE CASCADE 
)