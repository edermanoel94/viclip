package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.design/x/clipboard"
)

type Data struct {
	Content string `json:"content"`
}

func init() {

	if err := clipboard.Init(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	logger := log.New("viclip")

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	e.POST("/clip", func(c echo.Context) error {

		//if c.RealIP() != "192.168.0.157" {
		//	return c.NoContent(http.StatusForbidden)
		//}

		data := &Data{}

		if err := c.Bind(data); err != nil {
			return err
		}

		dataByte := []byte(data.Content)

		clipboard.Write(clipboard.FmtText, dataByte)

		logger.Infof("Data copied successfully from source app: %s", c.Request().Header.Get("X-Source-App"))

		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
