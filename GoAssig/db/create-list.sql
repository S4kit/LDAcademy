DROP TABLE IF EXISTS Product;
DROP TABLE IF EXISTS Shop;
DROP TABLE IF EXISTS Category;
DROP TABLE IF EXISTS User;


CREATE TABLE Shop (ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,Name VARCHAR(20),Address VARCHAR(20));

CREATE TABLE Product (ID  INT NOT NULL PRIMARY KEY AUTO_INCREMENT, ShopID int , Name VARCHAR(20),Description VARCHAR(20) , Categories VARCHAR (128), FOREIGN KEY (`ShopID`) REFERENCES Shop(`ID`));
 
CREATE Table Category (ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,Name VARCHAR(20),Description VARCHAR(128));

CREATE TABLE User (ID  INT NOT NULL PRIMARY KEY AUTO_INCREMENT,Email VARCHAR(50) , Password VARCHAR(128));


INSERT INTO Category (Name, Description) VALUES ("Clothes","For men and women");
INSERT INTO Category (Name, Description) VALUES ("Automobile","Old and New Brands for Cars");
INSERT INTO Category (Name, Description)  VALUES ("Tools & Accessories","Tools for your home project and Accessories for your decoration idea");
