package main

import (
	"errors"
	"fmt"

	"github.com/andreykaipov/goobs/api/requests/sceneitems"
	"github.com/andreykaipov/goobs/api/typedefs"
	"github.com/spf13/cobra"
)

var (
	sceneItemCmd = &cobra.Command{
		Use:   "sceneitem",
		Short: "manage scene items",
		Long:  `The sceneitem command manages a scene's items`,
		RunE:  nil,
	}

	listSceneItemsCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all items of a scene",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("list requires a scene")
			}
			return listSceneItems(args[0])
		},
	}

	toggleSceneItemCmd = &cobra.Command{
		Use:   "toggle",
		Short: "Toggles visibility of a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("toggle requires a scene and scene-item")
			}
			return toggleSceneItem(args[0], args[1:]...)
		},
	}

	showSceneItemCmd = &cobra.Command{
		Use:   "show",
		Short: "Makes a scene-item visible",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("show requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(true, args[0], args[1:]...)
		},
	}

	hideSceneItemCmd = &cobra.Command{
		Use:   "hide",
		Short: "Hides a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("hide requires a scene and scene-item(s)")
			}
			return setSceneItemVisible(false, args[0], args[1:]...)
		},
	}

	getSceneItemVisibilityCmd = &cobra.Command{
		Use:   "visible",
		Short: "Show visibility status of a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("visible requires a scene and scene-item")
			}
			return getSceneItemVisibility(args[0], args[1:]...)
		},
	}

	centerSceneItemCmd = &cobra.Command{
		Use:   "center",
		Short: "Horizontally centers a scene-item",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("center requires a scene and scene-item")
			}
			return centerSceneItem(args[0], args[1:]...)
		},
	}
)

func listSceneItems(sceneName string) error {
	par := &sceneitems.GetSceneItemListParams{
		SceneName: &sceneName,
	}

	resp, err := client.SceneItems.GetSceneItemList(par)
	if err != nil {
		return err
	}

	for _, item := range resp.SceneItems {
		fmt.Println(item.SourceName)
	}

	return nil
}

func setSceneItemVisible(visible bool, scene string, items ...string) error {
	if len(items) == 0 {
		return nil
	}

	itemsPar := &sceneitems.GetSceneItemListParams{
		SceneName: &scene,
	}

	itemsResp, err := client.SceneItems.GetSceneItemList(itemsPar)
	if err != nil {
		return err
	}

	itemsMap := make(map[string]*typedefs.SceneItem, len(itemsResp.SceneItems))
	for _, item := range itemsResp.SceneItems {
		itemsMap[item.SourceName] = item
	}

	for _, itemName := range items { //TODO: eliminate this loop as soon as API allows
		item, ok := itemsMap[itemName]
		if !ok || item.SceneItemEnabled == visible {
			continue
		}

		p := &sceneitems.SetSceneItemEnabledParams{
			SceneItemEnabled: &visible,
			SceneItemId:      &item.SceneItemID,
			SceneName:        &scene,
		}

		_, err := client.SceneItems.SetSceneItemEnabled(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func toggleSceneItem(scene string, items ...string) error {
	if len(items) == 0 {
		return nil
	}

	itemsPar := &sceneitems.GetSceneItemListParams{
		SceneName: &scene,
	}

	itemsResp, err := client.SceneItems.GetSceneItemList(itemsPar)
	if err != nil {
		return err
	}

	itemsMap := make(map[string]*typedefs.SceneItem, len(itemsResp.SceneItems))
	for _, item := range itemsResp.SceneItems {
		itemsMap[item.SourceName] = item
	}

	for _, itemName := range items { //TODO: eliminate this loop as soon as API allows
		item, ok := itemsMap[itemName]
		if !ok {
			continue
		}

		enabled := !item.SceneItemEnabled
		p := &sceneitems.SetSceneItemEnabledParams{
			SceneItemEnabled: &enabled,
			SceneItemId:      &item.SceneItemID,
			SceneName:        &scene,
		}

		_, err := client.SceneItems.SetSceneItemEnabled(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func getSceneItemVisibility(scene string, items ...string) error {
	if len(items) == 0 {
		return nil
	}

	itemsMap := make(map[string]bool, len(items))
	for _, itemName := range items {
		itemsMap[itemName] = true
	}

	itemsPar := &sceneitems.GetSceneItemListParams{
		SceneName: &scene,
	}

	itemsResp, err := client.SceneItems.GetSceneItemList(itemsPar)
	if err != nil {
		return err
	}

	for _, item := range itemsResp.SceneItems {
		if itemsMap[item.SourceName] {
			fmt.Printf("%s: %t\n", item.SourceName, item.SceneItemEnabled)

		}
	}

	return nil
}

func centerSceneItem(scene string, items ...string) error {
	if len(items) == 0 {
		return nil
	}

	vidSet, err := client.Config.GetVideoSettings()
	if err != nil {
		return err
	}

	posX := vidSet.BaseWidth / 2
	posY := vidSet.BaseHeight / 2

	itemsPar := &sceneitems.GetSceneItemListParams{
		SceneName: &scene,
	}

	itemsResp, err := client.SceneItems.GetSceneItemList(itemsPar)
	if err != nil {
		return err
	}

	itemsMap := make(map[string]*typedefs.SceneItem, len(itemsResp.SceneItems))
	for _, item := range itemsResp.SceneItems {
		itemsMap[item.SourceName] = item
	}

	for _, itemName := range items {
		item, ok := itemsMap[itemName]
		if !ok {
			continue
		}

		item.SceneItemTransform.PositionX = posX
		item.SceneItemTransform.PositionY = posY

		transPar := &sceneitems.SetSceneItemTransformParams{
			SceneItemId:        &item.SceneItemID,
			SceneItemTransform: &item.SceneItemTransform,
			SceneName:          &scene,
		}

		_, err = client.SceneItems.SetSceneItemTransform(transPar)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	sceneItemCmd.AddCommand(centerSceneItemCmd)
	sceneItemCmd.AddCommand(toggleSceneItemCmd)
	sceneItemCmd.AddCommand(showSceneItemCmd)
	sceneItemCmd.AddCommand(hideSceneItemCmd)
	sceneItemCmd.AddCommand(getSceneItemVisibilityCmd)
	sceneItemCmd.AddCommand(listSceneItemsCmd)
	rootCmd.AddCommand(sceneItemCmd)
}
