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

type MetaConfig struct {
	TitleTemplate       string `json:"title_template"`
	DescriptionTemplate string `json:"description_template"`
	KeywordsTemplate    string `json:"keywords_template"`
	FallbackTitle       string `json:"fallback_title"`
	FallbackDescription string `json:"fallback_description"`
	FallbackKeywords    string `json:"fallback_keywords"`
}
