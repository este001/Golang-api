package Model

type Book struct {
	Id 		string `json:"id"`
	Isbn 	string `json:"isbn"`
	Title 	string `json:"title"`
	Author 	*Author `json:"author"`
}

// Author

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}