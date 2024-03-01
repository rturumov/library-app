package library_app

type BooksList struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type UsersList struct {
	Id     int
	UserId int
	BookId int
}

type Book struct {
	Id                int    `json:"id"`
	Title             string `json:"title"`
	Author            string `json:"author"`
	Genre             string `json:"genre"`
	YearOfPublication string `json:"year_of_publication"`
}

type ListBooks struct {
	Id     int
	UserId int
	BookId int
}
