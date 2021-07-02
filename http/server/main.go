package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Ping struct {
	Status string
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, Ping{Status: "pong"})
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TestServer struct {
	Users []*User
}

func (s *TestServer) showUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if _, u := s.getUser(id); u != nil {
		return c.JSON(http.StatusOK, u)
	}

	return c.NoContent(http.StatusNotFound)
}
func (s *TestServer) getUser(id int) (int, *User) {
	for i, u := range s.Users {
		if u.ID == id {
			return i, u
		}
	}
	return 0, nil
}

type showUsersResponse struct {
	Users []*User `json:"users"`
}

func (s *TestServer) showUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, showUsersResponse{Users: s.Users})
}

func (s *TestServer) registUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	var err error
	u.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if i, user := s.getUser(u.ID); user != nil {
		s.Users[i] = u
	} else {
		s.Users = append(s.Users, u)
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	s := TestServer{
		Users: []*User{},
	}

	e.GET("/ping", ping)
	e.POST("/user/:id", s.registUser)
	e.GET("/user/:id", s.showUser)
	e.GET("/users", s.showUsers)

	e.Logger.Fatal(e.Start(":8080"))
}
