package repositories_test

import (
	"database/sql"
	"github.com/atsuof/sampleapi/models"
	"github.com/atsuof/sampleapi/repositories"
	"reflect"
	"testing"
	"time"
)

func TestInsertArticle(t *testing.T) {
	tx, err := testDB.Begin()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err = tx.Rollback()
		if err != nil {
			t.Fatal(err)
		}
	})
	type args struct {
		db      *sql.DB
		article models.Article
	}
	tests := []struct {
		name    string
		args    args
		want    models.Article
		wantErr bool
	}{
		{
			name: "insert01",
			args: args{
				db: testDB,
				article: models.Article{
					Title:    "insert01",
					Contents: "contents01",
					UserName: "aaa",
				},
			},
			want: models.Article{
				ID:       8,
				Title:    "insert01",
				Contents: "contents01",
				UserName: "aaa",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repositories.InsertArticle(tt.args.db, tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertArticle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectArticleDetail(t *testing.T) {

	exceptTime, _ := time.Parse("2025-02-28 15:38:59", "2025-02-28 15:38:59")

	type args struct {
		db        *sql.DB
		articleId int
	}

	tests := []struct {
		name    string
		args    args
		want    models.Article
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{
				db:        testDB,
				articleId: 1,
			},
			want: models.Article{
				ID:        1,
				Title:     "firstPost",
				Contents:  "This is my first blog",
				UserName:  "saki",
				NiceNum:   3,
				Comments:  nil,
				CreatedAt: exceptTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(tt.args.db, tt.args.articleId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectArticleDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectArticleDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	type args struct {
		db   *sql.DB
		page int
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Article
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repositories.SelectArticleList(tt.args.db, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectArticleList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectArticleList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateNiceNum(t *testing.T) {
	type args struct {
		db        *sql.DB
		articleID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.UpdateNiceNum(tt.args.db, tt.args.articleID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateNiceNum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
