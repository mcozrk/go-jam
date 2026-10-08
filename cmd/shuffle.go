/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shuffleCmd represents the shuffle command
var shuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Shuffle all songs within the music library",
	Long:  `Shuffle all songs within the music library`,
	Run:   shuffleMusic,
}

func init() {
	rootCmd.AddCommand(shuffleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shuffleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shuffleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func shuffleMusic(cmd *cobra.Command, args []string) {
	fmt.Println("shuffle called")
}
