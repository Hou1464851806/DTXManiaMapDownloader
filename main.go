package main

import (
	"DTXMapDownload/app/client"
	"DTXMapDownload/app/config"
	"DTXMapDownload/pkg/global"
	"DTXMapDownload/pkg/utils"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	global.Settings = config.NewConfig()
	err := global.Settings.Load()
	if err != nil {
		fmt.Printf("Load settings fail, use default config: %v", err)
	}
	log.Printf("%#v\n", global.Settings)
	c := client.NewCollector(global.Settings.SourceURL)
	rootCMD := &cobra.Command{
		Short: "DTXMania Map Download Command",
		Args:  cobra.NoArgs,
	}
	//listCMD := &cobra.Command{
	//	Use:   "list",
	//	Short: "getSongsInfo available maps",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		c.getSongsInfo(constants.SourceURL)
	//	},
	//}
	downloadCMD := &cobra.Command{
		Use:   "download [index]",
		Args:  cobra.ExactArgs(1),
		Short: "Download the specified map from source",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Println("You download", name)
			c.Download(name)
		},
	}
	searchCMD := &cobra.Command{
		Use:   "search [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Search song's map from source",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Println("You search", name)
			c.Search(name)
		},
	}

	configCMD := &cobra.Command{
		Use:   "config",
		Short: "Manage downloader configs",
	}

	configSetCMD := &cobra.Command{
		Use: "set [key] [value]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("you need to use 2 exact args, you only use %d args now", len(args))
			}
			validArgs := []string{"game", "source"}
			if !utils.ContainsString(args[0], validArgs) {
				return fmt.Errorf("you can only use these args: %s", strings.Join(cmd.ValidArgs, " "))
			}
			return nil
		},
		Short: "Set downloader config",
		Run: func(cmd *cobra.Command, args []string) {
			key, value := args[0], args[1]
			fmt.Printf("You set config [%s] to [%s]\n", key, value)
			c.SetConfig(key, value)
		},
	}
	configListCMD := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List downloader config",
		Run: func(cmd *cobra.Command, args []string) {
			global.Settings.List()
		},
	}
	//rootCMD.AddCommand(listCMD)
	rootCMD.AddCommand(downloadCMD)
	rootCMD.AddCommand(searchCMD)
	configCMD.AddCommand(configSetCMD, configListCMD)
	rootCMD.AddCommand(configCMD)
	err = rootCMD.Execute()
	if err != nil {
		os.Exit(1)
	}
}
