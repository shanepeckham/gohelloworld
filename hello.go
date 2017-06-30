package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	res, err := http.Get("http://gohelloworld2.default.svc.cluster.local")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	ro := string(robots[:])

	name, err := os.Hostname()
	if err == nil {
		this.Ctx.WriteString("{'app1':" + "'" + name + "'}" + ro)
	} else {
		this.Ctx.WriteString("hello world from go Kubernetes")
	}
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
