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
	"bytes"
)


// Shows the running kernels
type Kernel struct{
	Name string `json:"name"`
	Id string	`json:"id"`
}

type Kernels []Kernel


// Cleans up the host and path info
func get_endpoint(host, path string) string {
	//remove trailing "/" from host
	if string(host[len(host)-1]) == "/" {
		host = host[:len(host)-1]
	}
	//remove leading "/" from path
	if string(path[0]) == "/" {
		path = path[1:]
	}
	return fmt.Sprintf("%s/%s", host, path)
}


func (k *Kernels) fetch(host string) {
	
	endpoint := get_endpoint(host, "/api/kernels")
	
	// Query the /api/kernels endpoint
	resp, err := http.Get(endpoint)
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
func kill_kernel(host, id string) {
	
	endpoint := get_endpoint(host, "/api/kernels")
		
	url := fmt.Sprintf("%s/%s",endpoint, id)
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
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("Killed %s\n",id)
}

// Kill the specified kernel
func kernel_action(host, target, action string) {

	endpoint := get_endpoint(host, "/api/kernels")
		
	kernels := find(host, target)
	for _,k := range kernels {
		url := fmt.Sprintf("%s/%s/%s",endpoint, k.Id, action)
		// Query the /api/kernels endpoint

		req, err := http.NewRequest("POST", url, nil)
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
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	    fmt.Printf("%s on %s\n",action, k.Id)

	}
}

// Search for the specified kernel and kill anything that starts with a match
func kill(host, target string) { 
	kernels := find(host, target)
	for _,k := range kernels {
		kill_kernel(host, k.Id)	
	}
}


// Search for the specified kernel and kill anything that starts with a match
func find(host, target string) Kernels {
	retVal := Kernels{} 
	kernels := &Kernels{}
	kernels.fetch(host)
	for _,k := range *kernels {
		if (len(target) == 0) || (strings.Index(k.Id,target) == 0) {
			retVal = append(retVal, k)
		}
	}
	return retVal
}

func start(host, k string) {
	
	endpoint := get_endpoint(host, "/api/kernels")
		
	fmt.Printf("Starting %s\n",k)
	// Create the payload as a map
	payload := make(map[string]string)
	payload["name"] = k
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	//
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(dat))
	if err != nil {
		log.Fatal(err)
	}
	// handle err
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	// Read the results from the build request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	new_kernel := Kernel{}
	err = json.Unmarshal(body, &new_kernel)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Starting %s\n", new_kernel.Id)
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Usage: "Host",
				},
			},
			Action: func(c *cli.Context) {
				s := &Kernels{}
				s.fetch(c.String("host"))
				for _,k := range *s {
					fmt.Printf("%s \t %s \n", k.Name, k.Id)
				}
			},
		},
		{
			Name: "kill",
			Usage: "Kills a kernel based on the first few chars of its id",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Usage: "Host",
				},
			},
			Action: func(c *cli.Context) {
				kill(c.String("host"), c.Args().First())
			},
		},
		{
			Name: "restart",
			Usage: "Restart a kernel based on the first few chars of its id",
			Action: func(c *cli.Context) {
				kernel_action(c.String("host"), c.Args().First(), "restart")
			},
		},
		{
			Name: "interrupt",
			Usage: "Interrupt a kernel based on the first few chars of its id",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Usage: "Host",
				},
			},
			Action: func(c *cli.Context) {
				kernel_action(c.String("host"), c.Args().First(), "interrupt")
			},
		},
		{
			Name: "start",
			Usage: "Starts the specified kernel",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Usage: "Host",
				},
			},
			Action: func(c *cli.Context) {
				start(c.String("host"), c.Args().First())
			},
		},
	}
	
	app.Run(os.Args)
}

