CREATE TABLE todos (
    ID text PRIMARY KEY,
    Content text NOT NULL,
    Status boolean NOT NULL,
    UpdatedAt timestamp,
    CreatedAt timestamp
);
