package store

import ("gochat/utils";
		"fmt";
		_ "modernc.org/sqlite";
		"database/sql";
		"log";
		// "encoding/json"
	)

var DB *sql.DB

func InitDB() {
	var err error
	
	// PUT IT HERE INSIDE A FUNCTION: 
	DB, err = sql.Open("sqlite", "./server/store/messages.db") //  Use = instead of := because DB is declared above
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	query:= `create table if not exists messages (
		id integer primary key autoincrement,
		name text,
		time integer,
		content text,
		status integer
	);`
	createTableSQL := query
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

}

type Store interface {
	Store(message utils.Message) error
}

type FileStore struct {
	filename string
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

func (f *FileStore) GetAll (message utils.Message) error{
	const query string = "select * from messages"
	res, err := DB.Exec(query)
	if err != nil {
		log.Printf("Failed to get messages: %v", err)
		return err
	}

	log.Printf("res:%v ", res)
	return nil
}