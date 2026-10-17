/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"jam/audio"

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
	songs := audio.LoadLibrary()

	for i, song := range songs {
		fmt.Printf("%d. %s - %s - %s\n", i+1, song.Name, song.Artist, song.Album)
	}
}
