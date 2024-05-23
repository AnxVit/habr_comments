package repositories

const (
	createPost = `
		INSERT INTO posts(content, author, title, commented)
		VALUES($1, $2, $3, $4)
		RETURNING id;
	`
	getAllPosts = `
		SELECT
			id,
			author,
			content,
			title,
			created_at,
			updated_at,
			commented
		FROM posts;
	`
	createComment = `
		INSERT INTO comments (content, author, path)
		values (
		$1, $2,
		(SELECT path FROM comments WHERE id = $3) || (select currval('comments_id_seq')::integer)
		);
	`

	createPostComment = `
		INSERT INTO comments (content, author, path)
		values (
		$1, $2, ARRAY [(select currval('comments_id_seq')::integer)])
		RETURNING id;
	`
	postcomment = `
		INSERT INTO post_comment (post_id, comment_id)
		values ($1, $2);
	`

	getPostByID = `
	SELECT
		author,
		content,
		title,
		created_at,
		updated_at,
		commented
	FROM posts
	WHERE id = $1;
	`
	getMainComments = `
	SELECT
		comment_id
	FROM post_comment
	WHERE post_id = $1;
	`
	getSubComments = `
	SELECT
		id,
		path,
		author,
		content
	FROM comments
	WHERE $1 && path
	ORDER BY id;
	`
	checkPost = `
	SELECT
		commented
	FROM posts
	WHERE id = $1;
	`
)
