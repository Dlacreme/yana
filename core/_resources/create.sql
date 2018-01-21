-- ---
-- Globals
-- ---

-- SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- SET FOREIGN_KEY_CHECKS=0;

-- ---
-- Table 'User'
-- Default user
-- ---

DROP TABLE IF EXISTS `User`;
		
CREATE TABLE `User` (
  `UserId` INTEGER NOT NULL AUTO_INCREMENT,
  `Email` VARCHAR(200) NOT NULL DEFAULT 'NULL',
  `Password` VARCHAR(250) NULL DEFAULT NULL,
  `Name` VARCHAR(200) NOT NULL DEFAULT 'NULL',
  `StatusId` SMALLINT(1) NOT NULL DEFAULT 1,
  `Token` INTEGER NULL DEFAULT NULL,
  `TypeId` SMALLINT(1) NOT NULL,
  `CreditAmount` DOUBLE NOT NULL DEFAULT 0,
  PRIMARY KEY (`UserId`)
) COMMENT 'Default user';

-- ---
-- Table 'Article'
-- 
-- ---

DROP TABLE IF EXISTS `Article`;
		
CREATE TABLE `Article` (
  `ArticleId` INTEGER NOT NULL AUTO_INCREMENT,
  `CreatedBy` INTEGER NOT NULL DEFAULT 0,
  `UpdatedBy` INTEGER NULL DEFAULT NULL,
  `CreatedAt` DATE NULL DEFAULT NULL,
  `UpdatedAt` DATE NULL DEFAULT NULL,
  `IsDisabled` SMALLINT(1) NOT NULL DEFAULT 0,
  `ContentId` INTEGER NULL DEFAULT NULL,
  `PublishedAt` DATE NULL DEFAULT NULL,
  `IsPublished` SMALLINT(1) NOT NULL DEFAULT 0,
  `Title` VARCHAR(500) NULL DEFAULT NULL,
  PRIMARY KEY (`ArticleId`)
);

-- ---
-- Table 'Tag'
-- 
-- ---

DROP TABLE IF EXISTS `Tag`;
		
CREATE TABLE `Tag` (
  `TagId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(200) NOT NULL,
  `IsDisabled` SMALLINT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`TagId`)
);

-- ---
-- Table 'ArticleTag'
-- 
-- ---

DROP TABLE IF EXISTS `ArticleTag`;
		
CREATE TABLE `ArticleTag` (
  `ArticleTagId` INTEGER NOT NULL AUTO_INCREMENT,
  `ArticleId` INTEGER NOT NULL,
  `TagId` INTEGER NOT NULL,
  PRIMARY KEY (`ArticleTagId`)
);

-- ---
-- Table 'ArticleContent'
-- 
-- ---

DROP TABLE IF EXISTS `ArticleContent`;
		
CREATE TABLE `ArticleContent` (
  `ArticleContentId` INTEGER NOT NULL AUTO_INCREMENT,
  `Introduction` TEXT NULL DEFAULT NULL,
  `Content` TEXT NULL DEFAULT NULL,
  PRIMARY KEY (`ArticleContentId`)
);

-- ---
-- Table 'Member'
-- 
-- ---

DROP TABLE IF EXISTS `Member`;
		
CREATE TABLE `Member` (
  `MemberId` INTEGER NOT NULL AUTO_INCREMENT,
  `UserId` INTEGER NULL DEFAULT NULL,
  `MemberTypeId` INTEGER NULL DEFAULT NULL,
  PRIMARY KEY (`MemberId`)
);

-- ---
-- Table 'Event'
-- 
-- ---

DROP TABLE IF EXISTS `Event`;
		
CREATE TABLE `Event` (
  `EventId` INTEGER NOT NULL AUTO_INCREMENT,
  `ArticleId` INTEGER NOT NULL,
  `EventTypeId` INTEGER NOT NULL,
  `StartDate` DATE NULL DEFAULT NULL,
  `EndDate` DATE NULL DEFAULT NULL,
  `Slots` INTEGER NULL DEFAULT NULL,
  PRIMARY KEY (`EventId`)
);

-- ---
-- Table 'EventType'
-- 
-- ---

DROP TABLE IF EXISTS `EventType`;
		
CREATE TABLE `EventType` (
  `EventTypeId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(200) NOT NULL,
  PRIMARY KEY (`EventTypeId`)
);

-- ---
-- Table 'Contact'
-- 
-- ---

DROP TABLE IF EXISTS `Contact`;
		
CREATE TABLE `Contact` (
  `ContactId` INTEGER NOT NULL AUTO_INCREMENT,
  `ContactTypeId` INTEGER NULL DEFAULT NULL,
  `Name` VARCHAR(255) NULL DEFAULT NULL,
  `Email` VARCHAR(255) NULL DEFAULT NULL,
  PRIMARY KEY (`ContactId`)
);

-- ---
-- Table 'ContactType'
-- 
-- ---

DROP TABLE IF EXISTS `ContactType`;
		
CREATE TABLE `ContactType` (
  `ContactTypeId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`ContactTypeId`)
);

-- ---
-- Table 'Guest'
-- 
-- ---

DROP TABLE IF EXISTS `Guest`;
		
CREATE TABLE `Guest` (
  `EventId` INTEGER NULL DEFAULT NULL,
  `ContactId` INTEGER NULL DEFAULT NULL
);

-- ---
-- Table 'EventMember'
-- 
-- ---

DROP TABLE IF EXISTS `EventMember`;
		
CREATE TABLE `EventMember` (
  `EventMemberId` INTEGER NOT NULL AUTO_INCREMENT,
  `EventId` INTEGER NOT NULL,
  `MemberId` INTEGER NULL DEFAULT NULL,
  `RegistrationDate` DATE NOT NULL,
  PRIMARY KEY (`EventMemberId`)
);

-- ---
-- Table 'MemberType'
-- 
-- ---

DROP TABLE IF EXISTS `MemberType`;
		
CREATE TABLE `MemberType` (
  `MemberTypeId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`MemberTypeId`)
);

-- ---
-- Table 'Privilege'
-- 
-- ---

DROP TABLE IF EXISTS `Privilege`;
		
CREATE TABLE `Privilege` (
  `PrivilegeId` INTEGER NOT NULL AUTO_INCREMENT,
  `MemberTypeId` INTEGER NOT NULL,
  `Label` VARCHAR(255) NOT NULL,
  `Description` VARCHAR(500) NULL DEFAULT NULL,
  PRIMARY KEY (`PrivilegeId`)
);

-- ---
-- Table 'Stock'
-- 
-- ---

DROP TABLE IF EXISTS `Stock`;
		
CREATE TABLE `Stock` (
  `StockId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(200) NOT NULL,
  `Quantity` INTEGER NULL DEFAULT NULL,
  `Size` INTEGER NULL DEFAULT NULL,
  `TypeId` SMALLINT(1) NOT NULL DEFAULT 1,
  `CategoryId` INTEGER NOT NULL,
  `FormatId` SMALLINT(1) NOT NULL,
  PRIMARY KEY (`StockId`)
);

-- ---
-- Table 'Product'
-- 
-- ---

DROP TABLE IF EXISTS `Product`;
		
CREATE TABLE `Product` (
  `ProductId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(255) NULL DEFAULT NULL,
  `TypeId` INTEGER NOT NULL,
  `IsDisabled` SMALLINT(1) NOT NULL DEFAULT 0,
  `new field` INTEGER NULL DEFAULT NULL,
  PRIMARY KEY (`ProductId`)
);

-- ---
-- Table 'Category'
-- 
-- ---

DROP TABLE IF EXISTS `Category`;
		
CREATE TABLE `Category` (
  `CategoryId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(200) NOT NULL,
  `UnitId` SMALLINT(1) NOT NULL,
  `ParentId` INTEGER NULL DEFAULT NULL,
  PRIMARY KEY (`CategoryId`)
);

-- ---
-- Table 'CategoryUnit'
-- 
-- ---

DROP TABLE IF EXISTS `CategoryUnit`;
		
CREATE TABLE `CategoryUnit` (
  `CategoryUnitId` SMALLINT(1) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`CategoryUnitId`)
);

-- ---
-- Table 'StockFormat'
-- 
-- ---

DROP TABLE IF EXISTS `StockFormat`;
		
CREATE TABLE `StockFormat` (
  `StockFormatId` SMALLINT(1) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(200) NOT NULL,
  PRIMARY KEY (`StockFormatId`)
);

-- ---
-- Table 'StockType'
-- 
-- ---

DROP TABLE IF EXISTS `StockType`;
		
CREATE TABLE `StockType` (
  `StockTypeId` SMALLINT(1) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`StockTypeId`)
);

-- ---
-- Table 'Price'
-- 
-- ---

DROP TABLE IF EXISTS `Price`;
		
CREATE TABLE `Price` (
  `PriceId` INTEGER NOT NULL AUTO_INCREMENT,
  `ProductId` INTEGER NOT NULL,
  `Price` DOUBLE NOT NULL DEFAULT 0,
  `StartDate` DATE NOT NULL,
  `EndDate` DATE NULL DEFAULT NULL,
  PRIMARY KEY (`PriceId`)
);

-- ---
-- Table 'ProductStock'
-- 
-- ---

DROP TABLE IF EXISTS `ProductStock`;
		
CREATE TABLE `ProductStock` (
  `ProductStockId` INTEGER NOT NULL AUTO_INCREMENT,
  `ProductId` INTEGER NOT NULL,
  `StockId` INTEGER NOT NULL,
  `Quantity` INTEGER NOT NULL,
  `StartDate` DATE NOT NULL,
  `EndDate` DATE NULL DEFAULT NULL,
  PRIMARY KEY (`ProductStockId`)
);

-- ---
-- Table 'ProductType'
-- 
-- ---

DROP TABLE IF EXISTS `ProductType`;
		
CREATE TABLE `ProductType` (
  `ProductTypeId` INTEGER NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(200) NOT NULL,
  PRIMARY KEY (`ProductTypeId`)
);

-- ---
-- Table 'Import'
-- 
-- ---

DROP TABLE IF EXISTS `Import`;
		
CREATE TABLE `Import` (
  `ImportId` INTEGER NOT NULL AUTO_INCREMENT,
  `Date` DATE NOT NULL,
  `File` VARCHAR(500) NULL DEFAULT NULL,
  PRIMARY KEY (`ImportId`)
);

-- ---
-- Table 'ImportStock'
-- 
-- ---

DROP TABLE IF EXISTS `ImportStock`;
		
CREATE TABLE `ImportStock` (
  `ImportId` INTEGER NOT NULL,
  `StockId` INTEGER NOT NULL,
  `Quantity` INTEGER NOT NULL,
  `BuyingPrice` DOUBLE NOT NULL
);

-- ---
-- Table 'Order'
-- 
-- ---

DROP TABLE IF EXISTS `Order`;
		
CREATE TABLE `Order` (
  `OrderId` INTEGER NOT NULL AUTO_INCREMENT,
  `TableId` SMALLINT(2) NOT NULL,
  `MemberId` INTEGER NULL DEFAULT NULL,
  `StaffId` INTEGER NULL DEFAULT NULL,
  `StatusId` SMALLINT(1) NOT NULL DEFAULT 1,
  `IsPaid` BOOLEAN NOT NULL,
  `TotalPrice` DOUBLE NULL DEFAULT NULL,
  `AmountPaid` DOUBLE NULL DEFAULT NULL,
  `PaymentMethodId` SMALLINT(1) NULL DEFAULT NULL,
  `InvoiceId` VARCHAR(25) NULL DEFAULT NULL,
  `CreatedBy` INTEGER NOT NULL,
  `UpdatedBy` INTEGER NOT NULL,
  `CreatedAt` DATE NOT NULL,
  `UpdatedAt` DATE NOT NULL,
  PRIMARY KEY (`OrderId`)
);

-- ---
-- Table 'OrderProduct'
-- 
-- ---

DROP TABLE IF EXISTS `OrderProduct`;
		
CREATE TABLE `OrderProduct` (
  `OrderProductId` INTEGER NOT NULL AUTO_INCREMENT,
  `ProductId` INTEGER NULL DEFAULT NULL,
  `OrderId` INTEGER NULL DEFAULT NULL,
  `ExtraOrderProductId` INTEGER NOT NULL DEFAULT NULL,
  `CreatedBy` INTEGER NOT NULL,
  `CreatedAt` DATE NOT NULL,
  PRIMARY KEY (`OrderProductId`)
);

-- ---
-- Table 'ExtraOrderProduct'
-- 
-- ---

DROP TABLE IF EXISTS `ExtraOrderProduct`;
		
CREATE TABLE `ExtraOrderProduct` (
  `ExtraOrderProductId` INTEGER NOT NULL AUTO_INCREMENT,
  `StockId` INTEGER NOT NULL,
  `Quantity` INTEGER NOT NULL,
  PRIMARY KEY (`ExtraOrderProductId`)
);

-- ---
-- Table 'Discount'
-- 
-- ---

DROP TABLE IF EXISTS `Discount`;
		
CREATE TABLE `Discount` (
  `DiscountId` INTEGER NOT NULL AUTO_INCREMENT,
  `OrderId` INTEGER NOT NULL,
  PRIMARY KEY (`DiscountId`)
);

-- ---
-- Table 'PaymentMethod'
-- 
-- ---

DROP TABLE IF EXISTS `PaymentMethod`;
		
CREATE TABLE `PaymentMethod` (
  `PaymentMethodId` SMALLINT(1) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`PaymentMethodId`)
);

-- ---
-- Table 'OrderStatus'
-- 
-- ---

DROP TABLE IF EXISTS `OrderStatus`;
		
CREATE TABLE `OrderStatus` (
  `OrderStatusId` SMALLINT(1) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`OrderStatusId`)
);

-- ---
-- Table 'Table'
-- 
-- ---

DROP TABLE IF EXISTS `Table`;
		
CREATE TABLE `Table` (
  `TableId` SMALLINT(2) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(200) NULL DEFAULT NULL,
  PRIMARY KEY (`TableId`)
);

-- ---
-- Table 'UserStatus'
-- 
-- ---

DROP TABLE IF EXISTS `UserStatus`;
		
CREATE TABLE `UserStatus` (
  `UserStatusId` SMALLINT(1) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`UserStatusId`)
);

-- ---
-- Table 'UserType'
-- 
-- ---

DROP TABLE IF EXISTS `UserType`;
		
CREATE TABLE `UserType` (
  `UserTypeId` SMALLINT(1) NOT NULL AUTO_INCREMENT,
  `Label` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`UserTypeId`)
);

-- ---
-- Table 'CreditHistory'
-- 
-- ---

DROP TABLE IF EXISTS `CreditHistory`;
		
CREATE TABLE `CreditHistory` (
  `CreditHistoryId` INTEGER NOT NULL AUTO_INCREMENT,
  `UserId` INTEGER NOT NULL,
  `Amount` DOUBLE NOT NULL,
  `Date` DATE NOT NULL,
  `Debit` BOOLEAN NOT NULL,
  PRIMARY KEY (`CreditHistoryId`)
);

-- ---
-- Table 'OrderCreditHistory'
-- 
-- ---

DROP TABLE IF EXISTS `OrderCreditHistory`;
		
CREATE TABLE `OrderCreditHistory` (
  `OrderId` INTEGER NOT NULL,
  `CreditHistoryId` INTEGER NOT NULL
);

-- ---
-- Table 'PriceVariable'
-- 
-- ---

DROP TABLE IF EXISTS `PriceVariable`;
		
CREATE TABLE `PriceVariable` (
  `PriceVariableId` INTEGER NOT NULL AUTO_INCREMENT,
  `ProductId` INTEGER NOT NULL,
  `Price` DOUBLE NOT NULL,
  `MinPrice` DOUBLE NULL DEFAULT NULL,
  `MaxPrice` DOUBLE NULL DEFAULT NULL,
  PRIMARY KEY (`PriceVariableId`)
);

-- ---
-- Foreign Keys 
-- ---

ALTER TABLE `User` ADD FOREIGN KEY (StatusId) REFERENCES `UserStatus` (`UserStatusId`);
ALTER TABLE `User` ADD FOREIGN KEY (TypeId) REFERENCES `UserType` (`UserTypeId`);
ALTER TABLE `Article` ADD FOREIGN KEY (CreatedBy) REFERENCES `User` (`UserId`);
ALTER TABLE `Article` ADD FOREIGN KEY (ContentId) REFERENCES `ArticleContent` (`ArticleContentId`);
ALTER TABLE `ArticleTag` ADD FOREIGN KEY (ArticleId) REFERENCES `Article` (`ArticleId`);
ALTER TABLE `ArticleTag` ADD FOREIGN KEY (TagId) REFERENCES `Tag` (`TagId`);
ALTER TABLE `Member` ADD FOREIGN KEY (MemberId) REFERENCES `User` (`UserId`);
ALTER TABLE `Member` ADD FOREIGN KEY (MemberTypeId) REFERENCES `MemberType` (`MemberTypeId`);
ALTER TABLE `Event` ADD FOREIGN KEY (ArticleId) REFERENCES `Article` (`ArticleId`);
ALTER TABLE `Event` ADD FOREIGN KEY (EventTypeId) REFERENCES `EventType` (`EventTypeId`);
ALTER TABLE `Contact` ADD FOREIGN KEY (ContactTypeId) REFERENCES `ContactType` (`ContactTypeId`);
ALTER TABLE `Guest` ADD FOREIGN KEY (EventId) REFERENCES `Event` (`EventId`);
ALTER TABLE `Guest` ADD FOREIGN KEY (ContactId) REFERENCES `Contact` (`ContactId`);
ALTER TABLE `EventMember` ADD FOREIGN KEY (EventId) REFERENCES `Event` (`EventId`);
ALTER TABLE `EventMember` ADD FOREIGN KEY (MemberId) REFERENCES `Member` (`MemberId`);
ALTER TABLE `Privilege` ADD FOREIGN KEY (MemberTypeId) REFERENCES `MemberType` (`MemberTypeId`);
ALTER TABLE `Stock` ADD FOREIGN KEY (TypeId) REFERENCES `StockType` (`StockTypeId`);
ALTER TABLE `Stock` ADD FOREIGN KEY (CategoryId) REFERENCES `Category` (`CategoryId`);
ALTER TABLE `Stock` ADD FOREIGN KEY (FormatId) REFERENCES `StockFormat` (`StockFormatId`);
ALTER TABLE `Product` ADD FOREIGN KEY (TypeId) REFERENCES `ProductType` (`ProductTypeId`);
ALTER TABLE `Category` ADD FOREIGN KEY (UnitId) REFERENCES `CategoryUnit` (`CategoryUnitId`);
ALTER TABLE `Category` ADD FOREIGN KEY (ParentId) REFERENCES `Category` (`CategoryId`);
ALTER TABLE `Price` ADD FOREIGN KEY (ProductId) REFERENCES `Product` (`ProductId`);
ALTER TABLE `ProductStock` ADD FOREIGN KEY (ProductId) REFERENCES `Product` (`ProductId`);
ALTER TABLE `ProductStock` ADD FOREIGN KEY (StockId) REFERENCES `Stock` (`StockId`);
ALTER TABLE `ImportStock` ADD FOREIGN KEY (ImportId) REFERENCES `Import` (`ImportId`);
ALTER TABLE `ImportStock` ADD FOREIGN KEY (StockId) REFERENCES `Stock` (`StockId`);
ALTER TABLE `Order` ADD FOREIGN KEY (TableId) REFERENCES `Table` (`TableId`);
ALTER TABLE `Order` ADD FOREIGN KEY (MemberId) REFERENCES `User` (`UserId`);
ALTER TABLE `Order` ADD FOREIGN KEY (StaffId) REFERENCES `User` (`UserId`);
ALTER TABLE `Order` ADD FOREIGN KEY (StatusId) REFERENCES `OrderStatus` (`OrderStatusId`);
ALTER TABLE `Order` ADD FOREIGN KEY (PaymentMethodId) REFERENCES `PaymentMethod` (`PaymentMethodId`);
ALTER TABLE `Order` ADD FOREIGN KEY (CreatedBy) REFERENCES `User` (`UserId`);
ALTER TABLE `Order` ADD FOREIGN KEY (UpdatedBy) REFERENCES `User` (`UserId`);
ALTER TABLE `OrderProduct` ADD FOREIGN KEY (ProductId) REFERENCES `Product` (`ProductId`);
ALTER TABLE `OrderProduct` ADD FOREIGN KEY (OrderId) REFERENCES `Order` (`OrderId`);
ALTER TABLE `OrderProduct` ADD FOREIGN KEY (ExtraOrderProductId) REFERENCES `ExtraOrderProduct` (`ExtraOrderProductId`);
ALTER TABLE `OrderProduct` ADD FOREIGN KEY (CreatedBy) REFERENCES `User` (`UserId`);
ALTER TABLE `ExtraOrderProduct` ADD FOREIGN KEY (StockId) REFERENCES `Stock` (`StockId`);
ALTER TABLE `Discount` ADD FOREIGN KEY (OrderId) REFERENCES `Order` (`OrderId`);
ALTER TABLE `CreditHistory` ADD FOREIGN KEY (UserId) REFERENCES `User` (`UserId`);
ALTER TABLE `OrderCreditHistory` ADD FOREIGN KEY (OrderId) REFERENCES `Order` (`OrderId`);
ALTER TABLE `OrderCreditHistory` ADD FOREIGN KEY (CreditHistoryId) REFERENCES `CreditHistory` (`CreditHistoryId`);
ALTER TABLE `PriceVariable` ADD FOREIGN KEY (ProductId) REFERENCES `Product` (`ProductId`);

-- ---
-- Table Properties
-- ---

-- ALTER TABLE `User` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Article` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Tag` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `ArticleTag` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `ArticleContent` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Member` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Event` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `EventType` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Contact` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `ContactType` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Guest` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `EventMember` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `MemberType` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Privilege` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Stock` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Product` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Category` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `CategoryUnit` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `StockFormat` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `StockType` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Price` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `ProductStock` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `ProductType` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Import` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `ImportStock` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Order` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `OrderProduct` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `ExtraOrderProduct` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Discount` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `PaymentMethod` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `OrderStatus` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `Table` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `UserStatus` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `UserType` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `CreditHistory` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `OrderCreditHistory` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `PriceVariable` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ---
-- Test Data
-- ---

-- INSERT INTO `User` (`UserId`,`Email`,`Password`,`Name`,`StatusId`,`Token`,`TypeId`,`CreditAmount`) VALUES
-- ('','','','','','','','');
-- INSERT INTO `Article` (`ArticleId`,`CreatedBy`,`UpdatedBy`,`CreatedAt`,`UpdatedAt`,`IsDisabled`,`ContentId`,`PublishedAt`,`IsPublished`,`Title`) VALUES
-- ('','','','','','','','','','');
-- INSERT INTO `Tag` (`TagId`,`Label`,`IsDisabled`) VALUES
-- ('','','');
-- INSERT INTO `ArticleTag` (`ArticleTagId`,`ArticleId`,`TagId`) VALUES
-- ('','','');
-- INSERT INTO `ArticleContent` (`ArticleContentId`,`Introduction`,`Content`) VALUES
-- ('','','');
-- INSERT INTO `Member` (`MemberId`,`UserId`,`MemberTypeId`) VALUES
-- ('','','');
-- INSERT INTO `Event` (`EventId`,`ArticleId`,`EventTypeId`,`StartDate`,`EndDate`,`Slots`) VALUES
-- ('','','','','','');
-- INSERT INTO `EventType` (`EventTypeId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `Contact` (`ContactId`,`ContactTypeId`,`Name`,`Email`) VALUES
-- ('','','','');
-- INSERT INTO `ContactType` (`ContactTypeId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `Guest` (`EventId`,`ContactId`) VALUES
-- ('','');
-- INSERT INTO `EventMember` (`EventMemberId`,`EventId`,`MemberId`,`RegistrationDate`) VALUES
-- ('','','','');
-- INSERT INTO `MemberType` (`MemberTypeId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `Privilege` (`PrivilegeId`,`MemberTypeId`,`Label`,`Description`) VALUES
-- ('','','','');
-- INSERT INTO `Stock` (`StockId`,`Label`,`Quantity`,`Size`,`TypeId`,`CategoryId`,`FormatId`) VALUES
-- ('','','','','','','');
-- INSERT INTO `Product` (`ProductId`,`Label`,`TypeId`,`IsDisabled`,`new field`) VALUES
-- ('','','','','');
-- INSERT INTO `Category` (`CategoryId`,`Label`,`UnitId`,`ParentId`) VALUES
-- ('','','','');
-- INSERT INTO `CategoryUnit` (`CategoryUnitId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `StockFormat` (`StockFormatId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `StockType` (`StockTypeId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `Price` (`PriceId`,`ProductId`,`Price`,`StartDate`,`EndDate`) VALUES
-- ('','','','','');
-- INSERT INTO `ProductStock` (`ProductStockId`,`ProductId`,`StockId`,`Quantity`,`StartDate`,`EndDate`) VALUES
-- ('','','','','','');
-- INSERT INTO `ProductType` (`ProductTypeId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `Import` (`ImportId`,`Date`,`File`) VALUES
-- ('','','');
-- INSERT INTO `ImportStock` (`ImportId`,`StockId`,`Quantity`,`BuyingPrice`) VALUES
-- ('','','','');
-- INSERT INTO `Order` (`OrderId`,`TableId`,`MemberId`,`StaffId`,`StatusId`,`IsPaid`,`TotalPrice`,`AmountPaid`,`PaymentMethodId`,`InvoiceId`,`CreatedBy`,`UpdatedBy`,`CreatedAt`,`UpdatedAt`) VALUES
-- ('','','','','','','','','','','','','','');
-- INSERT INTO `OrderProduct` (`OrderProductId`,`ProductId`,`OrderId`,`ExtraOrderProductId`,`CreatedBy`,`CreatedAt`) VALUES
-- ('','','','','','');
-- INSERT INTO `ExtraOrderProduct` (`ExtraOrderProductId`,`StockId`,`Quantity`) VALUES
-- ('','','');
-- INSERT INTO `Discount` (`DiscountId`,`OrderId`) VALUES
-- ('','');
-- INSERT INTO `PaymentMethod` (`PaymentMethodId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `OrderStatus` (`OrderStatusId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `Table` (`TableId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `UserStatus` (`UserStatusId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `UserType` (`UserTypeId`,`Label`) VALUES
-- ('','');
-- INSERT INTO `CreditHistory` (`CreditHistoryId`,`UserId`,`Amount`,`Date`,`Debit`) VALUES
-- ('','','','','');
-- INSERT INTO `OrderCreditHistory` (`OrderId`,`CreditHistoryId`) VALUES
-- ('','');
-- INSERT INTO `PriceVariable` (`PriceVariableId`,`ProductId`,`Price`,`MinPrice`,`MaxPrice`) VALUES
-- ('','','','','');