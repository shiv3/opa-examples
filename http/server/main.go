package main

//go:generate sqlite3 schema/models.sqlite < schema/ddl.sql
//go:generate sqlboiler sqlite3

import (
	"context"
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shiv3/opa-examples/http/server/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"net/http"
	"strconv"
)

type Ping struct {
	Status string
}

//ping health check用途のエンドポイント
func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, Ping{Status: "pong"})
}

type TestServer struct {
	db *sql.DB
}

//showUser リクエストされたUserIDを返す
func (s *TestServer) showUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	users, err := models.Users(qm.Where("id = ?", int64(id))).One(context.Background(), s.db)
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		log.Error(err)
		return err
	}
	if users != nil {
		return c.JSON(http.StatusOK, users)
	}

	return c.NoContent(http.StatusNotFound)
}

type showUsersResponse struct {
	Users []*models.User `json:"users"`
}

//showUsers DBにある全Userを返す
func (s *TestServer) showUsers(c echo.Context) error {
	users, err := models.Users().All(context.Background(), s.db)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, showUsersResponse{Users: users})
}

//registUser リクエストされたUserを登録する
func (s *TestServer) registUser(c echo.Context) error {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := u.Insert(context.Background(), tx, boil.Infer()); err != nil {
		log.Error(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func main() {
	db, err := sql.Open("sqlite3", "schema/models.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	s := TestServer{
		db: db,
	}

	e.Logger.SetLevel(log.DEBUG)

	e.GET("/ping", ping)
	e.POST("/user/:id", s.registUser)
	e.GET("/user/:id", s.showUser)
	e.GET("/users", s.showUsers)
	e.Logger.Fatal(e.Start(":8080"))
}
