package main

import (
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/golang/glog"
)

var (
	url  *string
	host *string
	port *string
)

// helm is a wrapper function to execute the helm command on the
// shell.
func helm(arguments []string) (output []byte, err error) {
	command := "helm"
	cmd := exec.Command(command, arguments...)

	// Combine stdout and stderr
	glog.Info("updating helm repository index")
	output, err = cmd.CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
		return output, err
	}

	return output, nil
}

// initRepo initialize a helm repository generating a index.yaml file
func initRepo() error {
	// generate helm index
	_, err := helm([]string{"repo", "index", "./charts/", "--url", *url})
	if err != nil {
		glog.Error(err.Error())
		return err
	}
	return nil
}

// upload uploads a given file to the the charts directory
func upload(c echo.Context) error {
	chartName := c.Param("chartName")
	// TODO - do some sanitising on chartName

	glog.Info("uploading " + chartName)
	f, err := os.Create("charts/" + chartName)
	defer f.Close()
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	_, err = io.Copy(f, c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	glog.Info("done uploading " + chartName)

	// generate helm index
	initRepo()

	return nil
}

// repo  serves back any files in the charts directory
// with content-type header set to text/yaml
func repo(c echo.Context) error {
	//c.Response().Header().Set("content-type", "text/yaml")
	c.Response().Header().Set("content-type", "text/plain; charset=utf-8")
	return c.File(path.Join("charts", c.Param("*")))
}

func init() {
	// Get command line options
	// repoURL is also the url that get's used to generate the helm repo index file
	url = flag.String("url", "http://localhost:1323/", "The URL where Helmet runs as a repository")
	host = flag.String("host", "127.0.0.1", "The address that Helmet listens on")
	port = flag.String("port", "1323", "The port that Helmet listens on")
	flag.Parse()

	// initialize the helm repository on startup.
	err := initRepo()
	if err != nil {
		glog.Fatal(err.Error())
	}
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoints
	e.PUT("/upload/:chartName", upload)

	// Serve the charts directory
	e.GET("/*", repo)

	// Start server
	e.Logger.Fatal(e.Start(*host + ":" + *port))
}
