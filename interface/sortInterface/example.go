package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title 	string
	Artist 	string
	Album  	string
	Year  	int
	Length 	time.Duration
}

var tracks = []*Track{
	{"go", "wby", "fuck sky", 2019, length("5ms") },
	{"go Ahead", "wyr", "up up up", 2018, length("1499m49s")},
	{"forget", "zxl", "down down down", 2015, length("100000m59s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2,' ',0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Leangth")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track
func(x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	value := []int{1,10,4,7,9}
	sort.Ints(value)
	fmt.Println(value)
	a :=[]string {"dhadfsd", "a", "g", "c"}
	sort.Strings(a)
	//fmt.Println(a)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

}