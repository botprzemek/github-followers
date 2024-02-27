package models

import "os/user"

type Cache struct {
	Quota int8        `json:"qouta"`
	Users []user.User `json:"users"`
}

func (cache Cache) GetQuota() int8 {
	return cache.Quota
}

func (cache Cache) GetUsers() []user.User {
	return cache.Users
}
