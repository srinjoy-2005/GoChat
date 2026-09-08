package store

import ("gochat/utils";
		"fmt"
	)


type Store interface {

	Store(message utils.Message) error
}

type FileStore struct {
	filename string
}

func (f *FileStore) Store(message utils.Message) error {
	// Implementation for storing message
	fmt.Println("stored")
	return nil
}