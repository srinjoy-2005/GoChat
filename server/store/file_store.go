package store

import ("gochat/utils";
		"fmt";
		_ "modernc.org/sqlite";
		"database/sql";
		"log";
		"encoding/json"
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
}

type Store interface {
	Store(message utils.Message) error
}

type FileStore struct {
	filename string
}

func (f *FileStore) Store(message utils.Message) error {
	// Implementation for storing message
	stats, _ := json.MarshalIndent(DB.Stats(), "", "  ")
	fmt.Println(string(stats))
	fmt.Println("stored")
	return nil
}