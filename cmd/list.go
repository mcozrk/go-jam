/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"jam/audio"
	"time"

	"github.com/spf13/cobra"
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

	//TODO: remove this placeholder and actuall set song library here
	// **********************
	song := audio.Song{
		Name:     "test",
		Artist:   "artist",
		Album:    "album",
		Duration: time.Now()} // initializes songs var

	song2 := audio.Song{
		Name:     "test2",
		Artist:   "artist2",
		Album:    "album2",
		Duration: time.Now()} // initializes songs var

	songs := []audio.Song{song, song2}
	//************************

	for _, song := range songs {
		fmt.Printf("%s %s %s %s\n", song.Name, song.Artist, song.Album, song.Duration)
	}
}
