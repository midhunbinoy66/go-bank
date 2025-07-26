package main

type Account struct {

	id int
	firstName string
	lastName string
	Number int64
}

func NewAccount(firstName, lastName string) *Account{
	return &Account{
		id: 1,
		firstName: firstName,
		lastName: lastName,
		Number: 5456465465465,
	}
}
