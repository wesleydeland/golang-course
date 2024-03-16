package main

import "fmt"

func main() {
	urls := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(urls)
	fmt.Println(urls["Google"])
	urls["LinkedIn"] = "https://linkedin.com"
	fmt.Println(urls)

	delete(urls, "Amazon Web Services")
	fmt.Println(urls)
}
