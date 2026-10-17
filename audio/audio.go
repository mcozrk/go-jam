package audio

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.senan.xyz/taglib"
)

type Song struct {
	Name        string
	Artist      string
	Album       string
	Duration    time.Time
	LibraryPath string
}

func LoadLibrary() []Song {
	files, err := os.ReadDir(filepath.Join("music"))

	if err != nil {
		fmt.Printf("error reading music directory: %v\n", err)
		return nil
	}

	songs := []Song{}

	for _, f := range files {
		tags, err := taglib.ReadTags(filepath.Join("music", f.Name()))

		libraryPath := filepath.Join("music", f.Name())

		if err != nil {
			fmt.Printf("error reading tags for %s: %v\n", f.Name(), err)
			continue
		}

		song := Song{
			Name:        getTag(tags, taglib.Title),
			Artist:      getTag(tags, taglib.Artist),
			Album:       getTag(tags, taglib.Album),
			LibraryPath: libraryPath,
		}

		songs = append(songs, song)
	}

	return songs
}

func getTag(tags map[string][]string, key string) string {
	values, ok := tags[key]
	if !ok || len(values) == 0 {
		return ""
	}

	return values[0]
}
