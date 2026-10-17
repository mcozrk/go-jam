/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a specific song within the music library",
	Long:  `Play a song from within the music library, to specify use --song flag`,
	Args:  cobra.NoArgs,
	RunE:  playMusic,
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	playCmd.Flags().StringP("song", "s", "", "specify a song to play from music library")
	if err := playCmd.MarkFlagRequired("song"); err != nil {
		panic(err)
	}

	// TODO: add ability to autocomplete selections based off /music contents
}

func playMusic(cmd *cobra.Command, args []string) error {
	songSelected := cmd.Flag("song").Value.String()

	file, err := os.Open("music/" + songSelected + ".mp3")
	if err != nil {
		return fmt.Errorf("Error opening song %q: %w", songSelected, err)
	}

	// decoding mp3 file
	stream, format, err := mp3.Decode(file)
	if err != nil {
		file.Close()
		return fmt.Errorf("Error decoding song %q: %w", songSelected, err)
	}
	defer stream.Close()

	// activate speaker
	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return fmt.Errorf("Error initializing speaker: %w", err)
	}

	defer speaker.Close()

	speaker.PlayAndWait(stream)

	return nil
}
