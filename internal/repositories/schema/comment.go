package schema

type Comment struct {
	ID      int    `db:"id"`
	Path    []int  `db:"path"`
	Author  string `db:"author"`
	Content string `db:"content"`
}
