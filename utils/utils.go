package utils

type Message struct {
	Name string //sender name
	Time uint32 //unix time
	Content string
	Status int //0:sent, 1:received at proxy, 2:received at server
}


