-- User
INSERT INTO UserType (Label) VALUES
('Admin'),
('Staff'),
('Member');

INSERT INTO UserStatus (Label) VALUES
('Active'),
('Out');

-- Stock
INSERT INTO StockType (Label) VALUES
('Finished'),
('Mixture');

INSERT INTO `CategoryUnit` (Label) VALUES
(''),
('cl'),
('kg');

INSERT INTO Category (Label, UnitId) VALUES
('Drink', 2),
('Food', 1);

INSERT INTO StockFormat (Label) VALUES
(''),
('Can'),
('Bottle');

-- Product
INSERT INTO ProductType (Label) VALUES
(''),
('Cocktail');

-- Order
INSERT INTO OrderStatus (Label) VALUES
('Pending'),
('Ready'),
('Done'),
('Cancel');

INSERT INTO PaymentMethod (Label) VALUES
('Credit'),
('Cash'),
('Card');

-- Table
INSERT INTO `Table` (Label) VALUES
('Default');