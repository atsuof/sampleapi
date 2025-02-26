package main

import (
	"time"
)

type Comment struct {
	CommentID int
	ArticleID int
	Message   string
	CreatedAt time.Time
}

type Article struct {
	ID        int
	Title     string
	Contents  string
	UserName  string
	NiceNum   int
	Comments  []Comment
	CreatedAt time.Time
}

//func main() {
//	comment1 := Comment{
//		CommentID: 1,
//		ArticleID: 2,
//		Message:   "aaa",
//		CreatedAt: time.Now(),
//	}
//
//	comment2 := Comment{
//		CommentID: 2,
//		ArticleID: 2,
//		Message:   "bbb",
//		CreatedAt: time.Now(),
//	}
//
//	article1 := Article{
//		ID:        1,
//		Title:     "Artincle01",
//		Contents:  "Content",
//		UserName:  "UserName",
//		NiceNum:   0,
//		Comments:  []Comment{comment1, comment2},
//		CreatedAt: time.Now(),
//	}
//
//	fmt.Printf("%+v\n", article1)
//
//	j, err := json.Marshal(article1)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	var out bytes.Buffer
//	jsonerr := json.Indent(&out, j, "", "  ")
//	if jsonerr != nil {
//		fmt.Println("err occurred.", jsonerr)
//	}
//	_, e := out.WriteTo(os.Stdout)
//	if e != nil {
//		fmt.Println("err occurred.", e)
//	}
//}
