CREATE TABLE [contact] (
    [id]        INTEGER         NOT NULL PRIMARY KEY AUTOINCREMENT,
    [name]      NVARCHAR(50)    NOT NULL,
    [email]     NVARCHAR(50)    NOT NULL,
    [phone]     NVARCHAR(20)    NULL,
    [message]   NVARCHAR(500)   NOT NULL
);