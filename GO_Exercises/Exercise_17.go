package main
import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// web application - rest api with go routines example
var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}


func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return "", err
	}
	wg.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage Endpoint Hit")
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		//we have to add and done to the original data so pointer varibale is used
		go fetch(url, &wg)
	}

	wg.Wait()

	//if we dont have the wait group it will return the response before performing the operations
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "Responses")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}