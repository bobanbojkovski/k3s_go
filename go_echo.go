package main

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	var (
		text = flag.String("text", "default", "text to print out")
		port = flag.String("port", ":1323", "http port")
	)
	flag.Parse()

	hostip := ""
	hostname, _ := os.Hostname()
	ips, _ := net.LookupIP(hostname)
	for _, ip := range ips {
		hostip = ip.String()
	}

	e := echo.New()
	e.GET("/"+*text, func(c echo.Context) error {
		return c.String(http.StatusOK, *text+"\n"+hostname+"\n"+hostip)
	})

	e.Logger.Fatal(e.Start(*port))

}
