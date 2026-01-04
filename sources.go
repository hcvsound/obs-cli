package main

import (
	"errors"
	"fmt"

	"github.com/andreykaipov/goobs/api/requests/inputs"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var (
	sourceCmd = &cobra.Command{
		Use:   "source",
		Short: "manage sources",
		Long:  `The source command manages sources`,
		RunE:  nil,
	}

	listSourcesCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listSources()
		},
	}

	toggleMuteCmd = &cobra.Command{
		Use:   "toggle-mute",
		Short: "Toggles mute",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("toggle-mute requires a source name as argument")
			}
			return toggleMute(args[0])
		},
	}
)

func listSources() error { //TODO: this should be "listInputs()"
	ilResp, err := client.Inputs.GetInputList()
	if err != nil {
		return err
	}

	fmt.Println("Sources\n=======")
	for _, v := range ilResp.Inputs {
		spew.Dump(v)
	}
	fmt.Println()

	spResp, err := client.Inputs.GetSpecialInputs()
	if err != nil {
		return err
	}

	fmt.Println("Special Sources")
	fmt.Println("===============")
	fmt.Printf("Desktop1: %s\n", spResp.Desktop1)
	fmt.Printf("Desktop2: %s\n", spResp.Desktop2)
	fmt.Printf("Mic1: %s\n", spResp.Mic1)
	fmt.Printf("Mic2: %s\n", spResp.Mic2)
	fmt.Printf("Mic3: %s\n", spResp.Mic3)

	return nil
}

func toggleMute(inputName string) error {
	par := &inputs.ToggleInputMuteParams{
		InputName: &inputName,
	}

	_, err := client.Inputs.ToggleInputMute(par)
	return err
}

func init() {
	sourceCmd.AddCommand(listSourcesCmd)
	sourceCmd.AddCommand(toggleMuteCmd)
	rootCmd.AddCommand(sourceCmd)
}
