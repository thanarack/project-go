package main

import (
	"github.com/thanarack/project-go/movie"
	"github.com/thanarack/project-go/ticket"
)

func main() {
	movieName := movie.Find("A00001")
	movie.Review(movieName, 4.1)
	ticket.Buy(movieName, 5)
}
