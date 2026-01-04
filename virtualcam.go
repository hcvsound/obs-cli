package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	virtualCamCmd = &cobra.Command{
		Use:   "virtualcam",
		Short: "manage virtual camera",
		Long:  `The virtualcam command manages the virtual camera`,
		RunE:  nil,
	}

	startStopVirtualCamCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggle virtual camera status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return starStopVirtualCam()
		},
	}

	startVirtualCamCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts virtual camera",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startVirtualCam()
		},
	}

	stopVirtualCamCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stops virtual camera",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopVirtualCam()
		},
	}

	virtualCamStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Reports virtual camera status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return virtualCamStatus()
		},
	}
)

func starStopVirtualCam() error {
	_, err := client.Outputs.ToggleVirtualCam()
	return err
}

func startVirtualCam() error {
	_, err := client.Outputs.StartVirtualCam()
	return err
}

func stopVirtualCam() error {
	_, err := client.Outputs.StopVirtualCam()
	return err
}

func virtualCamStatus() error {
	r, err := client.Outputs.GetVirtualCamStatus()
	if err != nil {
		return err
	}

	fmt.Print("Virtual camera: ")
	if !r.OutputActive {
		fmt.Print("not ")
	}
	fmt.Println("active")

	// fmt.Printf("Timecode: %s\n", r.VirtualCamTimecode)	TODO: actually obtain this
	return nil
}

func init() {
	virtualCamCmd.AddCommand(startStopVirtualCamCmd)
	virtualCamCmd.AddCommand(startVirtualCamCmd)
	virtualCamCmd.AddCommand(stopVirtualCamCmd)
	virtualCamCmd.AddCommand(virtualCamStatusCmd)
	rootCmd.AddCommand(virtualCamCmd)
}
