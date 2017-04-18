package storage

import (
	"github.com/wpferg/services/structs"
)

var store structs.MessageList
var currentMaxId = 1

func Get() structs.MessageList {
	return store
}

func Add(message structs.Message) int {
	message.ID = currentMaxId
	currentMaxId++
	store = append(store, message)
	return message.ID
}

func Remove(id int) {
	index := -1

	for i, message := range store {
		if message.ID == id {
			index = i
		}
	}

	store = append(store[:index], store[index+1:]...)
}
