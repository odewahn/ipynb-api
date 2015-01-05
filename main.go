package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
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
func kill(c *cli.Context) {
	id := c.Args().First()
	fmt.Printf("Looking for %s\n",id)
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

