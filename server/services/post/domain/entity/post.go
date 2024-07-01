package entity

type Post struct {
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	FilePath  string `json:"file_path"`
	CoverPath string `json:"cover_path"`
}
