CREATE DATABASE  IF NOT EXISTS `archer_club_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `archer_club_db`;
-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: localhost    Database: archer_club_db
-- ------------------------------------------------------
-- Server version	8.0.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `archer`
--

DROP TABLE IF EXISTS `archer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `archer` (
  `ArcherID` int NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(255) NOT NULL,
  `LastName` varchar(255) NOT NULL,
  `DOB` date NOT NULL,
  `Gender` varchar(255) NOT NULL,
  PRIMARY KEY (`ArcherID`)
) ENGINE=InnoDB AUTO_INCREMENT=653 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `archerevent`
--

DROP TABLE IF EXISTS `archerevent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `archerevent` (
  `EventID` int NOT NULL AUTO_INCREMENT,
  `EventName` varchar(255) DEFAULT NULL,
  `StartDate` date NOT NULL,
  `EndDate` date NOT NULL,
  PRIMARY KEY (`EventID`),
  KEY `EventName_indx` (`EventName`)
) ENGINE=InnoDB AUTO_INCREMENT=607 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `catergory`
--

DROP TABLE IF EXISTS `catergory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `catergory` (
  `CatergoryID` int NOT NULL AUTO_INCREMENT,
  `RoundID` int DEFAULT NULL,
  `AgeGroup` varchar(255) NOT NULL,
  `GenderBracket` varchar(255) NOT NULL,
  `BowType` varchar(255) NOT NULL,
  PRIMARY KEY (`CatergoryID`),
  KEY `AgeGroup_indx` (`AgeGroup`),
  KEY `RoundID` (`RoundID`),
  CONSTRAINT `catergory_ibfk_1` FOREIGN KEY (`RoundID`) REFERENCES `round` (`RoundID`)
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `end`
--

DROP TABLE IF EXISTS `end`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `end` (
  `EndID` int NOT NULL AUTO_INCREMENT,
  `RoundNum` int DEFAULT NULL,
  `RoundID` int DEFAULT NULL,
  `EventID` int NOT NULL,
  `CatergoryID` int DEFAULT NULL,
  `ArcherID` int NOT NULL,
  `TargetType` varchar(255) NOT NULL,
  `Distance` varchar(255) NOT NULL,
  `Date` date NOT NULL,
  `Arrow1` int NOT NULL,
  `Arrow2` int NOT NULL,
  `Arrow3` int NOT NULL,
  `Arrow4` int NOT NULL,
  `Arrow5` int NOT NULL,
  `Arrow6` int NOT NULL,
  `TOTAL` int GENERATED ALWAYS AS ((((((`Arrow1` + `Arrow2`) + `Arrow3`) + `Arrow4`) + `Arrow5`) + `Arrow6`)) VIRTUAL,
  `Approved` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`EndID`),
  KEY `EventID` (`EventID`),
  KEY `ArcherID` (`ArcherID`),
  KEY `CatergoryID` (`CatergoryID`),
  KEY `RoundID` (`RoundID`),
  CONSTRAINT `end_ibfk_1` FOREIGN KEY (`EventID`) REFERENCES `archerevent` (`EventID`),
  CONSTRAINT `end_ibfk_2` FOREIGN KEY (`ArcherID`) REFERENCES `archer` (`ArcherID`),
  CONSTRAINT `end_ibfk_4` FOREIGN KEY (`CatergoryID`) REFERENCES `catergory` (`CatergoryID`),
  CONSTRAINT `end_ibfk_5` FOREIGN KEY (`RoundID`) REFERENCES `round` (`RoundID`)
) ENGINE=InnoDB AUTO_INCREMENT=7161 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `round`
--

DROP TABLE IF EXISTS `round`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `round` (
  `RoundName` varchar(255) DEFAULT NULL,
  `RoundID` int NOT NULL AUTO_INCREMENT,
  `TotalArrows` int NOT NULL,
  `PossibleScore` int NOT NULL,
  `10m` varchar(255) DEFAULT NULL,
  `20m` varchar(255) DEFAULT NULL,
  `30m` varchar(255) DEFAULT NULL,
  `40m` varchar(255) DEFAULT NULL,
  `50m` varchar(255) DEFAULT NULL,
  `60m` varchar(255) DEFAULT NULL,
  `70m` varchar(255) DEFAULT NULL,
  `90m` varchar(255) DEFAULT NULL,
  `DateAdded` date DEFAULT NULL,
  `DateChange` date DEFAULT NULL,
  PRIMARY KEY (`RoundID`),
  KEY `RoundName_indx` (`RoundName`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-12-10 14:42:39
