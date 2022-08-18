package repository

import (
	"boosters/api/post"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewPosts(db *sqlx.DB) *PostgresDB {
	return &PostgresDB{db: db}
}

func (p *PostgresDB) CreatePost(post post.Post) (int, error) {
	var id int
	create_at := time.Now().UTC()
	query := "insert into posts(title, content, create_at, update_at) values($1,$2,$3, $4) returning id;"
	row := p.db.QueryRow(query, post.Title, post.Content, create_at, create_at)
	logrus.Info(row)
	if err := row.Scan(&id); err != nil {
		logrus.Error(err)
		return 0, err
	}
	return id, nil
}

func (p *PostgresDB) GetAllPosts() ([]post.Post, error) {
	var posts []post.Post
	quary := "select id, title, content, create_at, update_at from posts;"
	if err := p.db.Select(&posts, quary); err != nil {

		return nil, err
	}
	return posts, nil
}

func (p *PostgresDB) GetById(id string) (post.Post, error) {
	var post post.Post
	quary := "select id, title,content,create_at, update_at from posts where id=$1;"
	if err := p.db.Get(&post, quary, id); err != nil {
		logrus.Error(err)
		return post, err
	}
	return post, nil

}

func (p *PostgresDB) UpdatePost(post post.Post, id string) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if post.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, post.Title)
		argId++
	}
	if post.Content != "" {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, post.Content)
		argId++
	}
	update_at := time.Now().UTC()
	upd := fmt.Sprintf("update_at=$%d", argId)
	setValues = append(setValues, upd)
	args = append(args, update_at)

	// title=$1
	// content=$2
	// update_at=$3
	// title=$1, content=$2 , update_at=$3
	set_query := strings.Join(setValues, ", ")
	logrus.Info(set_query)
	logrus.Info(args...)
	query := fmt.Sprintf("UPDATE posts SET %s WHERE id=%s;", set_query, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := p.db.Exec(query, args...)
	return err

}
func (p *PostgresDB) DeletePost(id string) error {
	query := "DELETE FROM posts WHERE id=$1"
	_, err := p.db.Exec(query, id)
	return err
}
