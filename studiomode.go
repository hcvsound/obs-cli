package main

import (
	"fmt"
	"strconv"

	"github.com/andreykaipov/goobs/api/requests/ui"
	"github.com/spf13/cobra"
)

var (
	studioModeCmd = &cobra.Command{
		Use:   "studiomode",
		Short: "manage studio mode",
		Long:  `The studiomode command manages the studio mode`,
		RunE:  nil,
	}

	disableStudioModeCmd = &cobra.Command{
		Use:   "disable",
		Short: "Disables the studio mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			return disableStudioMode()
		},
	}

	enableStudioModeCmd = &cobra.Command{
		Use:   "enable",
		Short: "Enables the studio mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			return enableStudioMode()
		},
	}

	studioModeStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Reports studio mode status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return studioModeStatus()
		},
	}

	toggleStudioModeCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggles the studio mode (enable/disable)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return toggleStudioMode()
		},
	}

	transitionToProgramCmd = &cobra.Command{
		Use:   "transition",
		Short: "Transition to program",
		RunE: func(cmd *cobra.Command, args []string) error {
			return transitionToProgram()
		},
	}
)

func SetStudioModeEnabled(enabled bool) error {
	par := ui.NewSetStudioModeEnabledParams().WithStudioModeEnabled(enabled)
	_, err := client.Ui.SetStudioModeEnabled(par)
	return err
}

func disableStudioMode() error {
	return SetStudioModeEnabled(false)
}

func enableStudioMode() error {
	return SetStudioModeEnabled(true)
}

// Determine if the studio mode is currently enabled in OBS.
func IsStudioModeEnabled() (bool, error) {
	r, err := client.Ui.GetStudioModeEnabled()
	return r.StudioModeEnabled, err
}

func studioModeStatus() error {
	isStudioModeEnabled, err := IsStudioModeEnabled()
	if err != nil {
		return err
	}

	fmt.Printf("Studio Mode: %s\n", strconv.FormatBool(isStudioModeEnabled))
	return nil
}

func toggleStudioMode() error {
	enabled, err := IsStudioModeEnabled()
	if err != nil {
		return err
	}

	return SetStudioModeEnabled(!enabled)
}

func transitionToProgram() error {
	_, err := client.Transitions.TriggerStudioModeTransition()
	return err
}

func init() {
	studioModeCmd.AddCommand(disableStudioModeCmd)
	studioModeCmd.AddCommand(enableStudioModeCmd)
	studioModeCmd.AddCommand(studioModeStatusCmd)
	studioModeCmd.AddCommand(toggleStudioModeCmd)
	studioModeCmd.AddCommand(transitionToProgramCmd)
	rootCmd.AddCommand(studioModeCmd)
}
