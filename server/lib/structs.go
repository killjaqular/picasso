package lib

type Config struct {
    Server struct {
        Port int `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        Host     string `yaml:"host"`
        Port     int    `yaml:"port"`
        Database string `yaml:"database"`
        Username string `yaml:"username"`
        Password string `yaml:"password"`
    } `yaml:"database"`
}

// TODO: Create type structs for uuids.uuid and accounts.user
/*
CREATE TABLE IF NOT EXISTS uuids.uuid (
    uuid UUID PRIMARY KEY,
    parentTable VARCHAR(64) -- The table that contains the actual asset of the UUID
);

CREATE TABLE IF NOT EXISTS accounts.user (
    uuid UUID PRIMARY KEY REFERENCES uuids.uuid(uuid),
    username VARCHAR(16) UNIQUE, -- No repeated usernames allowed
    password VARCHAR(256) -- SHA256
);
*/
