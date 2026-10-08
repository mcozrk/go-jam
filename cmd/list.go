/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"jam/audio"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"go.senan.xyz/taglib"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all songs within the music library",
	Long:  `Lists all songs and their name, artist, album, and duration`,
	Run:   listSongs,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listSongs(cmd *cobra.Command, args []string) {
	files, err := os.ReadDir(filepath.Join("music"))

	if err != nil {
		fmt.Printf("error reading music directory: %v\n", err)
		return
	}

	songs := []audio.Song{}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		tags, err := taglib.ReadTags(filepath.Join("music", f.Name()))

		if err != nil {
			fmt.Printf("error reading tags for %s: %v\n", f.Name(), err)
			continue
		}

		song := audio.Song{
			Name:   getTag(tags, taglib.Title),
			Artist: getTag(tags, taglib.Artist),
			Album:  getTag(tags, taglib.Album),
		}

		songs = append(songs, song)
	}

	for i, song := range songs {
		fmt.Printf("%d. %s - %s - %s\n", i+1, song.Name, song.Artist, song.Album)
	}
}

func getTag(tags map[string][]string, key string) string {
	values, ok := tags[key]
	if !ok || len(values) == 0 {
		return ""
	}

	return values[0]
}
