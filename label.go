package main

import (
	"errors"

	"github.com/andreykaipov/goobs/api/requests/inputs"
	"github.com/spf13/cobra"
)

var (
	labelCmd = &cobra.Command{
		Use:   "label",
		Short: "manage text labels",
		Long:  `The label command manages text labels`,
		RunE:  nil,
	}

	textCmd = &cobra.Command{
		Use:   "text",
		Short: "Changes a text label",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("text requires a source and the new text")
			}
			return changeLabel(args[0], args[1])
		},
	}
)

func changeLabel(source string, text string) error {
	overlay := true
	p := &inputs.SetInputSettingsParams{
		InputName: &source,
		InputSettings: map[string]any{
			"Text": text,
		},
		Overlay: &overlay,
	}

	//TODO: first check if the new text is applicable at all
	_, err := client.Inputs.SetInputSettings(p)
	return err
}

func init() {
	labelCmd.AddCommand(textCmd)
	rootCmd.AddCommand(labelCmd)
}
