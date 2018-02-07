# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.21)
# Database: ideaparLog
# Generation Time: 2018-02-07 11:10:45 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table goal
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goal`;

CREATE TABLE `goal` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `dataUrl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `goal` WRITE;
/*!40000 ALTER TABLE `goal` DISABLE KEYS */;

INSERT INTO `goal` (`id`, `name`, `desc`, `dataUrl`)
VALUES
	(1,'testName','testDesc','http://www.ideapar.com'),
	(2,'test2','testDesc22','http://www.baidu.com'),
	(3,'chang','chang','chang'),
	(5,'注册人数','网站一周内注册人数','http://www.ideapar.com/users/num'),
	(6,'注册人数','网站一周内注册人数','http://www.ideapar.com/users/num');

/*!40000 ALTER TABLE `goal` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table pv
# ------------------------------------------------------------

DROP TABLE IF EXISTS `pv`;

CREATE TABLE `pv` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `date` timestamp NULL DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `referrerUrl` varchar(255) DEFAULT NULL,
  `browserUuId` varchar(255) DEFAULT NULL,
  `userAgent` varchar(255) DEFAULT '',
  `userId` varchar(255) DEFAULT NULL,
  `userIp` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `pv` WRITE;
/*!40000 ALTER TABLE `pv` DISABLE KEYS */;

INSERT INTO `pv` (`id`, `date`, `url`, `referrerUrl`, `browserUuId`, `userAgent`, `userId`, `userIp`)
VALUES
	(113,'2018-02-07 16:26:17','http://www.icl.main/','','c3ed0d1e5b4a14268fa396e7a78ea530','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36','1-5a530898bcd0b','172.19.0.12'),
	(114,'2018-02-07 16:26:21','http://www.icl.main/','http://www.icl.main/','c3ed0d1e5b4a14268fa396e7a78ea530','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36','1-5a530898bcd0b','172.19.0.12'),
	(115,'2018-02-07 16:26:24','http://www.icl.main/','http://www.icl.main/','c3ed0d1e5b4a14268fa396e7a78ea530','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36','1-5a530898bcd0b','172.19.0.12'),
	(116,'2018-02-07 16:44:30','http://www.icl.main/','','c3ed0d1e5b4a14268fa396e7a78ea530','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36','1-5a530898bcd0b','172.19.0.12'),
	(117,'2018-02-07 16:44:58','http://www.icl.main/','','c3ed0d1e5b4a14268fa396e7a78ea530','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36','','172.19.0.12');

/*!40000 ALTER TABLE `pv` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
