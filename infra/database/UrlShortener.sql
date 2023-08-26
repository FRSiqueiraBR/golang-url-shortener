CREATE TABLE [short_url] ( 
	[long] VARCHAR(255)  NOT NULL,
    [hash] VARCHAR(7) NOT NULL,
    [expiration] TIMESTAMP NOT NULL,
    [timestamp] TIMESTAMP NOT NULL
); 