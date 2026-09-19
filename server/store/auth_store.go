package store

import (
	_ "gochat/utils"
	"log"
)

type Auth interface{
	Store (username, password string) error
	GetByUsername (username string) (string, error)
}

//implements Auth
type FileAuth struct{
	filename string
}

func (f *FileAuth) Store(username, hashedPassword string) error{
	//TODO: FORGOT TO DEDUPLICATE! DO NEXT
	query := `INSERT INTO users (username, password_hash) VALUES (?, ?)`
    _, err := DB.Exec(query, username, hashedPassword)
	if (err!=nil){
		return err
	}
	log.Println("stored in auth")
	return nil
}

func (f *FileAuth) GetByUsername(username string) (string, error) {
    var hash string
    err := DB.QueryRow(
        "SELECT password_hash FROM users WHERE username = ?",
        username,
    ).Scan(&hash)
	
    if err != nil {
        return "", err
    }
    return hash, nil
}

