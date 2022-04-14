CREATE TABLE `hero`(
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'PRIMARY KEY', 
  name varchar(25) NOT NULL,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE INDEX `name_index` ON `hero` (`name`);
DROP INDEX `name_index` ON `hero`;
CREATE INDEX `name_index` ON `hero` (`name`(4));