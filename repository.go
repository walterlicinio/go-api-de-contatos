package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func connect() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	return session, err
}

func addUser(user User) error {
	session, err := connect()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.DB("API").C("Users").Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func getUser(email string) (User, error) {
	session, err := connect()
	if err != nil {
		return User{}, err
	}
	defer session.Close()

	var user User
	err = session.DB("API").C("Users").Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func getAllUsers() ([]User, error) {
	session, err := connect()
	if err != nil {
		return []User{}, err
	}
	defer session.Close()

	var users []User
	err = session.DB("API").C("Users").Find(bson.M{}).All(&users)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func deleteUser(email string) error {
	session, err := connect()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.DB("API").C("Users").Remove(bson.M{"email": email})
	if err != nil {
		return err
	}

	return nil
}
