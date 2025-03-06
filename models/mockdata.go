package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "first comment",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}

	Article1 = Article{
		ID:        1,
		Title:     "first article",
		Contents:  "this is the test article",
		UserName:  "jhon",
		NiceNum:   1,
		Comments:  []Comment{Comment1, Comment2},
		CreatedAt: time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "second article",
		Contents:  "this is the test article",
		UserName:  "jhon",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
