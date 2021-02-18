package server

import (
	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func genUUID() string {
	id := guuid.New()
	return id.String()
}

// Handlers
func (s *Server) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is response from a web server!")
}

func (s *Server) GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, storage)
}

func (s *Server) CreateUser(c echo.Context) error {
	user := User{}
	user.ID = genUUID()
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	storage = append(storage, user)
	return c.JSON(http.StatusCreated, user)
}

func (s *Server) GetUser(c echo.Context) error {
	//s.conn.GetUser()
	id := c.Param("id")
	for _, user := range storage {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func (s *Server) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	for item := range storage {
		if storage[item].ID == id {
			storage = append(storage[:item], storage[item+1:]...)
			return c.JSON(http.StatusOK, "User was deleted")
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}
