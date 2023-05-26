-- Create a new table called 'TableName' in schema 'SchemaName'
-- Drop the table if it already exists
IF OBJECT_ID('db_pet.users', 'U') IS NOT NULL
DROP TABLE db_pet.users
GO
-- Create the table in the specified schema
CREATE TABLE db_pet.users
(
  id INT NOT NULL PRIMARY KEY,
  name [NVARCHAR](50) NOT NULL,
  password [NVARCHAR](50) NOT NULL
);
GO