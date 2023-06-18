-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	8.0.33-0ubuntu0.22.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `code` int DEFAULT NULL,
                         `username` varchar(50) NOT NULL,
                         `email` varchar(100) NOT NULL,
                         `password` varchar(100) NOT NULL,
                         `firstname` varchar(100) NOT NULL DEFAULT '',
                         `lastname` varchar(100) NOT NULL DEFAULT '',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,381663,'oQhlUrk','uVCpMjU@loZGxIy.net','$2a$14$7Rrptpg3fC2Xc.hnyzg1BeAsdSN3ZkoD.Jk8.V4azfwxlfPMrA26u','Manuel','Jerde','2023-06-18 11:42:41'),(2,173900,'WbufiNK','NXyIEcT@XhalCEe.ru','$2a$14$zbL/PvvNdkDtqIYPk4Ue3ux2zqwS9oEpDhCIjsjzxePzRQq7t5L/u','Alisha','Stracke','2023-06-18 11:43:06'),(3,247267,'urdSpTm','HdQWumy@kYKVWpR.com','$2a$14$ElfIsF0Re2ne54mP5pT0vey8gqXBSbOEn.KeUBzDejS3abwINq3d2','Jerad','Block','2023-06-18 11:43:09'),(4,948163,'plVcebf','aadlByd@JfBoaMX.org','$2a$14$Arvesve5jYpPFMZWzp71yuKg.v7lAv5tNga/ybxnTvM4p9fG6i2zq','Robyn','Breitenberg','2023-06-18 11:43:13'),(5,756460,'oMTFkVy','PgJUdAP@OtqIMmK.ru','$2a$14$Yzn4XeluOZwreotdYMUgc.StQi2szFiLr17069ceCk/jiZHkV7LnG','Jennifer','Friesen','2023-06-18 11:47:01'),(6,875577,'dHoMxur','swcHbpQ@ZNmZhYh.ru','$2a$14$JwGYbIemAYLNuIougIm.5.pSgXMUkHP23dfAcUWEatrTid6jELg/i','Mikayla','Kertzmann','2023-06-18 12:08:17'),(7,783246,'ipoGbKB','wxecGUh@CipQlNO.biz','$2a$14$Sha2zZ5u9..t8hpv/bvnYeKzpqy1sJrFEqAdu3MKSwiCqPwGeFFe.','Adelia','Toy','2023-06-18 12:24:10'),(8,393222,'crCLcXa','NdARUmX@crbXkhX.org','$2a$14$OMxcAhYm019JHVNczAHsreCA9Sp6FLS/AOYUhSLpFmYNkjWC26cJW','Owen','Toy','2023-06-18 12:24:54'),(9,172063,'rNdkAtA','iIBDXUH@DCpVSHD.info','$2a$14$rTE4hqpURywPheVzijvm0ubl591HMRSyI8SjKxnFgWeJeKv.npDbe','Payton','Pouros','2023-06-18 12:25:23'),(10,102803,'ITGmuHM','HDDQqml@yPbiZMg.net','$2a$14$DWoBu1k9ObskZdHFzAO1HeP4NngecXLcUpj4fyV857uocnfAm7BNu','Mackenzie','Kunde','2023-06-18 12:27:00'),(11,563324,'qeZfEti','LkKmSrA@yIfNpkj.net','$2a$14$Dn96aMAggwMgOtiQPP/Oe.Bh/pQdIKO9GydY/dJ6Eibwq2yiWTxk2','Myrtice','Cummerata','2023-06-18 12:27:13'),(12,878949,'veeavMa','KfWBvpO@QSUphcK.com','$2a$14$PuPrZBME2TFEfOMHJMmD0.6M1L7p3M1q8Wf7oSMH/OU7vd9w6pKL.','Lillie','O\'Conner','2023-06-18 12:30:16'),(13,132408,'ySACEYF','WHpKIki@dEhVSjC.org','$2a$14$UObFM5ZHOyqm/9w1qjt4kO1/bXqcVKQD2S.KAMDor5XF4iRPqzky6','Alvena','Koelpin','2023-06-18 12:33:20'),(14,442537,'IxnKFMF','wAeRJvB@jyZEVpQ.biz','$2a$14$zuMNgLcWdSjbzDVW/L0TnuVkpgzYl3P/DurYRftZWbFllNVjf5/4K','Kris','Gutmann','2023-06-18 12:33:31'),(15,954546,'ImjlwZf','eSRQZdf@pXVheMp.com','$2a$14$p718kGNyasaRvodXtwwgguYyQWfQy.r9BH35QZgnN/9JYO8JJnLLi','Stuart','Franecki','2023-06-18 12:34:15'),(16,632259,'iZNVEtM','ALwwdtR@OMUjrKx.info','$2a$14$n/q0/LQ9Vct6vLx.L0pA1.ACQcxBH92RUK2e8Ssz5U7H0c.giYNs2','Gladys','Brakus','2023-06-18 12:34:38'),(17,843939,'RQrAyxc','aERMKNa@GqygAke.org','$2a$14$6X0.RwxBPS994tIiQIUC5ee/U/CqFGjxxIRxgw1qyBjuuu78Vg4b2','Oran','Weimann','2023-06-18 12:41:22'),(18,175246,'vTbODeR','tPoOSUx@ZCcBxsM.ru','$2a$14$hyD/6ntcno.gebXnOC/yaertLlNK1v8cP24EfO.exAwqULiaUYX8W','Juana','Lindgren','2023-06-18 12:41:35'),(19,452188,'jiqUvqu','tSLTQGq@odCQkrX.com','$2a$14$t8bg.inn4kWvi.dGLcpnUu9ffl2cdE6ADhL3sXZhkPZn3zT8X5Nz.','Tatyana','Franecki','2023-06-18 12:42:57'),(20,389559,'LALgwBw','EkYFWNP@gKKBKtX.net','$2a$14$HhT7UpzZ3iimcp8yNR8omO4xF/AUM61XFWi./uFgq3INkgzObQTMe','Holden','Ziemann','2023-06-18 12:46:36'),(21,614111,'DIOylMC','wNVjfUE@PdURRpW.ru','$2a$14$y20pyJyqcCda7eSrcrF0JOht7cDdTDp43CVcsVh8VBml/.KgVaNPe','Eloisa','Lindgren','2023-06-18 12:52:22');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-18 15:58:36
