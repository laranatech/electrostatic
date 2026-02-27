package pages

type Page struct {
	Content  []byte
	Filepath string
	Route    string
	Meta     map[string]string
}

type PageResponse struct {
	Content []byte
	Code    int
}
