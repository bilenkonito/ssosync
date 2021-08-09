package google

import admin "google.golang.org/api/admin/directory/v1"

type Alias = admin.Alias
type Channel = admin.Channel
type Customer = admin.Customer
type Group = admin.Group
type Member = admin.Member
type Printer = admin.Printer
type User = admin.User
type UserPhoto = admin.UserPhoto
type Entity struct {
	Alias
	Channel
	Customer
	Group
	Member
	Printer
	User
	UserPhoto
}

type UserName = admin.UserName
