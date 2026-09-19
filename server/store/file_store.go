package store

import (
	"database/sql"
	"fmt"
	"gochat/utils"
	"log"

	// _ "golang.org/x/tools/go/analysis/passes/defers"
	_ "modernc.org/sqlite"
	// "encoding/json"
)

var DB *sql.DB

func InitDB() {
	var err error
	
	DB, err = sql.Open("sqlite", "./server/store/messages.db") //  Use = instead of := because DB is declared above
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// 1. Create messages table
    messagesQuery := `create table if not exists messages (
        id integer primary key autoincrement,
        name text,
        time integer,
        content text,
        status integer
    );`
	_, err = DB.Exec(messagesQuery)
    if err != nil {
        log.Fatalf("Failed to create messages table: %v", err)
    }

    // 2. Create users table for authentication
    usersQuery := `create table if not exists users (
        id integer primary key autoincrement,
        username text unique not null,
        password_hash text not null
    );`

    _, err = DB.Exec(usersQuery)
    if err != nil {
        log.Fatalf("Failed to create users table: %v", err)
    }
}

type Store interface {
	Store(message utils.Message) error
	GetAll() error
}

type FileStore struct {
	filename string
}

type CloudStore struct {
	// Add fields for cloud storage configuration if needed
}

func (c *CloudStore) Store(message utils.Message) error {
	// Implementation for storing message in cloud storage
	return nil
}

func (c * CloudStore) GetAll() error{
	return nil
}

func (f *FileStore) Store(message utils.Message) error {
	// Implementation for storing message
	const query string = "INSERT INTO messages (name, time, content, status) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, message.Name, message.Time, message.Content, message.Status)
	if err != nil {
		log.Printf("Failed to store message: %v", err)
		return err
	}
	fmt.Println("stored")
	return nil
}

func (f *FileStore) GetAll () error{
	const query string = "select * from messages"
	res, err := DB.Query(query)
	defer res.Close()	
	if err != nil {
		log.Printf("Failed to get messages: %v", err)
		return err
	}

	log.Printf("res:%v ", res)
	return nil
}

