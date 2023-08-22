// server/lib/methods.go

package lib

import (
    "github.com/google/uuid"
)

func (user *UUIDRecord) SetUUID(uuid uuid.UUID) {
    user.UUID = uuid
}

func (user *UUIDRecord) GetUUID() uuid.UUID {
    return user.UUID
}

func (user *UUIDRecord) SetParentTable(parentTable string) {
    user.ParentTable = parentTable
}

func (user *UUIDRecord) GetParentTable() string {
    return user.ParentTable
}

func (user *User) SetUUID(uuid uuid.UUID) {
    user.UUID = uuid
}

func (user *User) GetUUID() uuid.UUID {
    return user.UUID
}

func (user *User) SetUsername(username string) {
    user.Username = username
}

func (user *User) GetUsername() string {
    return user.Username
}

func (user *User) SetPassword(password string) {
    user.Password = password
}

func (user *User) GetPassword() string {
    return user.Password
}
