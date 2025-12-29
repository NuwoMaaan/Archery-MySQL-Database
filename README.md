## Overview

This is a basic terminal application written in Golang which uses a relational database design I completed prior that is for an archery club. This project was for me to get my feet wet in programming with Golang, the application can be run with `docker-compose`, which seeds a MySQL database with mock data and the application will prompt predefined query options. The result will display the query string and the rows.

(An explanation of the archery club requirements and database design will be added later on) 

## Quick Start

**Start application:**
```powershell
docker compose run --rm menu-service
```
**Stop application & remove database volume:**
```powershell
docker-compose down
docker volume rm archery-mysql-database_db_data
```
## Entity Relationship Diagram
<img width="1265" height="758" alt="image" src="https://github.com/user-attachments/assets/1fad53c5-b7f8-4af0-afae-c10f46f143c6" />

## MySQL Queries
**Select Best Round For Particular Round Out Of All Archers**
```MySQL
SELECT archer.ArcherID, archer.FirstName, archer.LastName, MAX(round_total_sum) AS PB_total
FROM 
(
    SELECT RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, SUM(total) AS round_total_sum
    FROM end
    INNER JOIN round on round.RoundID = end.RoundID
    INNER JOIN archer on archer.ArcherID = end.ArcherID
    WHERE round.RoundName = 'Hobart' AND Approved = 1 AND round.DateAdded IS NOT NULL AND round.DateChange IS NULL
    GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName
) AS round_sums 
INNER JOIN  archer on archer.ArcherID = round_sums.ArcherID
GROUP BY archer.ArcherID;
```
**Select Personal Best For Selected Round**
```MySQL
SELECT MAX(round_total_sum) AS PB_total
FROM 
(
    SELECT RoundNum, SUM(total) AS round_total_sum
    FROM end
    INNER JOIN round on round.RoundID = end.RoundID
    WHERE round.RoundName = 'Hobart' AND Approved = 1 AND ArcherID = 1 
    AND round.DateAdded IS NOT NULL and round.DateChange IS NULL
    GROUP BY RoundNum
) AS round_sums; 
```
**Get end count by ID**
```MySQL
SELECT ArcherID, COUNT(ArcherID) FROM end GROUP BY ArcherID;
```
**Get archer by ID**
```MySQL
SELECT FirstName, LastName, Gender FROM archer WHERE ArcherId = 1;
```
**search for round and catergory requirements by Round Name & Category Age**
```MySQL
SELECT round.RoundID, round.RoundName, catergory.AgeGroup, catergory.GenderBracket, catergory.BowType,
10m,20m,30m,40m,50m,60m,70m,90m FROM catergory
INNER JOIN round on round.RoundID = catergory.RoundID
WHERE round.RoundName = 'Perth' AND catergory.AgeGroup = "Under 21";
```
**Select Score By Round Order By Total**
```MySQL
SELECT RoundNum, round.RoundName, archer.ArcherID, archer.FirstName, Date, archer.LastName, SUM(total) AS RoundTotalSum
FROM end
INNER JOIN round on round.RoundID = end.RoundID
INNER JOIN archer on archer.ArcherID = end.ArcherID
WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20'
GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, round.RoundName, Date
ORDER BY RoundTotalSum DESC;
```
**Select Score By Round Order By Date**
```MySQL
SELECT RoundNum, round.RoundName, archer.ArcherID, archer.FirstName, Date, archer.LastName, SUM(total) AS RoundTotalSum
FROM end
INNER JOIN round on round.RoundID = end.RoundID
INNER JOIN archer on archer.ArcherID = end.ArcherID
WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20'
GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, round.RoundName, Date
ORDER BY Date DESC;
```
**Select Score By Round Order By RoundNum**
```MySQL
SELECT RoundNum, round.RoundName, archer.ArcherID, archer.FirstName, Date, archer.LastName, SUM(total) AS RoundTotalSum
FROM end
INNER JOIN round on round.RoundID = end.RoundID
INNER JOIN archer on archer.ArcherID = end.ArcherID
WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart'
GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, round.RoundName, Date
ORDER BY RoundNum DESC;
```
**Search archers score per all ends above certain value (50)**
```MySQL
SELECT end.ArcherID, COUNT(TOTAL) FROM end
LEFT JOIN archer ON archer.ArcherID = end.ArcherID
WHERE TOTAL > 50 AND Approved = 1 GROUP BY end.ArcherID;
```
**Search championship scores by particular championship event name**
```MySQL
SELECT archerevent.EventName, catergory.AgeGroup, catergory.GenderBracket, archer.FirstName, archer.LastName, SUM(TOTAL) 
AS TOTAL FROM end
INNER JOIN archer ON archer.ArcherID = end.ArcherID
INNER JOIN archerevent on archerevent.EventID = end.EventID
INNER JOIN catergory on catergory.CatergoryID = end.CatergoryID
WHERE archerevent.EventName = 'December Championship' AND Approved = 1
GROUP BY end.ArcherID, catergory.AgeGroup, catergory.GenderBracket;
```
