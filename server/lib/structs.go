package lib

import (
    // Non-standard Golang
    "github.com/google/uuid"
)

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

/*
CREATE TABLE IF NOT EXISTS uuids.uuid (
    uuid UUID PRIMARY KEY,
    parentTable VARCHAR(64) -- The table that contains the actual asset of the UUID
);
*/
type UUIDRecord struct {
    UUID        uuid.UUID `json:"uuid"`
    ParentTable string    `json:"parentTable"`
}

/*
CREATE TABLE IF NOT EXISTS accounts.user (
    uuid UUID PRIMARY KEY REFERENCES uuids.uuid(uuid),
    username VARCHAR(16) UNIQUE, -- No repeated usernames allowed
    password VARCHAR(256) -- SHA256
);
*/
type User struct {
    UUID     uuid.UUID `json:"uuid"`
    Username string    `json:"username"`
    Password string    `json:"password"`
}
