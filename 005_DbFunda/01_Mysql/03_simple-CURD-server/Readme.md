# MySQL CURD operations Webserver

This is basic **C**reate **U**pdate **R**ead and **D**elete operations example with a webserver.

Here we use a sample Data base structure as follows:

```SQL
CREATE TABLE customer_trial (
		cID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
		cName text NOT NULL,
		cPoints int NOT NULL
    )
```

We are creating a simple tabled called `customer_trial` where 
 
 - Customer ID `cID` is the Primary key and Auto generated
 - Customer Name `cName` is a string field
 - Customer Points `cPoints` stores the number of points they have earned

The webserver in this project is to provide a way to do the following actions with this table:

1. Open / Create the Table (`customer_trial`) from the Database
2. Add Records to the Table
3. Read All records from the Table
4. Find Specific records in the table
5. Update / modify records already in the table
6. Delete records from the table
7. Delete the table itself

## Attribution & Thanks

The main goal of the project was to test my Knowhow and Learning 
from Todd Mcleod's (@GoesToEleven) Courses:

1. [Learn How to Code : Google's Go pogramming language](https://greatercommons.com/learn/golang)

This Couse is also available Free after [Sign Up @ GreaterCommons](https://greatercommons.com/register): https://greatercommons.com/cwg

2. [HTML & CSS Course on Udemy](https://www.udemy.com/html-tutorial/)

This is a Great Course and worth every penny. Its more of a investment in yourself than a course.

3. [Golang Web Dev Course on Udemy](https://www.udemy.com/go-programming-language/)

I have only finished till around *Section 12* as of (28th March 2018) about **Relational Databases**. 

Hope to finish the complete course and build some thing useful.