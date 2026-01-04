package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	streamCmd = &cobra.Command{
		Use:   "stream",
		Short: "manage streams",
		Long:  `The stream command manages streams`,
		RunE:  nil,
	}

	startStopStreamCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggle streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startStopStream()
		},
	}

	startStreamCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startStream()
		},
	}

	stopStreamCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stops streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopStream()
		},
	}

	streamStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Reports streaming status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return streamStatus()
		},
	}
)

func startStopStream() error {
	_, err := client.Stream.ToggleStream()
	return err
}

func startStream() error {
	_, err := client.Stream.StartStream()
	return err
}

func stopStream() error {
	_, err := client.Stream.StopStream()
	return err
}

func streamStatus() error {
	statResp, err := client.Stream.GetStreamStatus()
	if err != nil {
		return err
	}

	fmt.Printf("Streaming: %s\n", strconv.FormatBool(statResp.OutputActive))
	if !statResp.OutputActive {
		return nil
	}

	fmt.Printf("Timecode: %s\n", statResp.OutputTimecode)

	servResp, err := client.Config.GetStreamServiceSettings()
	if err != nil {
		return err
	}

	fmt.Printf("URL: %s\n", servResp.StreamServiceSettings.Server)
	return nil
}

func init() {
	streamCmd.AddCommand(startStopStreamCmd)
	streamCmd.AddCommand(startStreamCmd)
	streamCmd.AddCommand(stopStreamCmd)
	streamCmd.AddCommand(streamStatusCmd)
	rootCmd.AddCommand(streamCmd)
}
