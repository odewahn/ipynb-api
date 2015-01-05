package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
)


// Shows the running kernels
type Kernels []struct{
	Name string `json:"name"`
	Id string	`json:"id"`
}


func (k *Kernels) fetch() {
	
	// Query the /api/kernels endpoint
	resp, err := http.Get("http://192.168.59.103:8888/api/kernels")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	// Read the results from the build request
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &k)
	if err != nil {
		log.Fatal(err)
	}
	
}


// Kill the specified kernel
func kill_kernel(k string) {
	url := fmt.Sprintf("http://192.168.59.103:8888/api/kernels/%s",k)
	// Query the /api/kernels endpoint
	// Query the /api/kernels endpoint
	
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// handle err
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// handle err	
	
	// Read the results from the build request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}

// Kill the specified kernel
func kill(c *cli.Context) {
	id := c.Args().First()
	fmt.Printf("Looking for %s\n",id)
	kernels := &Kernels{}
	kernels.fetch()
	for _,k := range *kernels {
		if strings.Index(k.Id,id) == 0 {
			kill_kernel(k.Id)
		}
	}
}



func main() {
	app := cli.NewApp()
	app.Name = "ipynb-api"
	app.Usage = "Client for ipynb api"
	app.Version = "0.0.0-alpha"
	app.Action = func(c *cli.Context) {
		fmt.Println("Nothing to do.  Try `help` or `-h` to see what's possible.")
	}
	app.Commands = []cli.Command{
		{
			Name:  "show",
			Usage: "Show active kernels",
			Action: func(c *cli.Context) {
				s := &Kernels{}
				s.fetch()
				for _,k := range *s {
					fmt.Printf("%s \t %s \n", k.Name, k.Id)
				}
			},
		},
		{
			Name: "kill",
			Usage: "Kills a kernel given the first 3 chars or its id",
			Action: func(c *cli.Context) {
				kill(c)
			},
		},
	}
	
	app.Run(os.Args)
}

