package main

import "fmt"

// Entry - Declare an interface for any item that needs to be displayed
// Title returns details of the item
// Has to be a function
type Entry interface {
	Title() string
}

type Song struct {
	Name string
	Artist string
	Album string
}

func (song Song) Title() string {
	return fmt.Sprintf("Song::\n\tTitle::%s \n\tArtist::%s \n\tAlbum::%s", song.Name, song.Artist, song.Album)
}

type Movie struct {
	Name string
	Director string
	Rating float32
}

func (movie Movie) Title() string {
	return fmt.Sprintf("Movie::\n\tTitle::%s \n\tDirector::%s \n\tRating::%.2f", movie.Name, movie.Director, movie.Rating)
}

func Display(e Entry) {
	fmt.Println(e.Title())
}

func main() {
	movie := Movie{
		Name:     "Forest Gumpp",
		Director: "Joe Ngigi",
		Rating:   9.5,
	}

	song := Song{
		"Tear Drops of Jupiter",
		"Train",
		"Magic",
	}

	Display(movie)
	Display(song)
}
