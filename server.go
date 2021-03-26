package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

// func helloHandler(w http.ResponseWriter, r *http.Request){
// 	if r.URL.Path != "/hello"{
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello!")
// }

func info(){
	app.Name = "Simple CLI"
	app.Usage = "CLI project for VirtusLab"
	app.Version = "1.0.0"
}

func runServer(fileName string){
	fmt.Println(fileName)
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func commands(){
	app.Commands = []cli.Command{
		{
			Name: 		"version",
			Aliases: 	[]string{"v"},
			Usage: 		"Check app version",
			Action: 	func(c *cli.Context) {
				fmt.Println(app.Version)
			},
		},
		{
			Name: 		"run",
			Aliases: 	[]string{"r"},
			Usage: 		"Run file server",
			Flags: 		[]cli.Flag{
				cli.StringFlag{
					Name: "file, f",
					Usage: "Input file to run on server",
				},
			},
			Action: 	func(c *cli.Context) error {
				runServer(c.String("file"))
				return nil
			},
		},
		
	}
}


func main(){
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}