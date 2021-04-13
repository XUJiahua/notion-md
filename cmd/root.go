/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xujiahua/notion-md/pkg/notion"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var rootPageID string
var token string
var outputDir string
var hugoImagePrefix string
var supportListView bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "notion-md",
	Short: "convert notion pages into markdowns",
	Run: func(cmd *cobra.Command, args []string) {
		notion.New(token, rootPageID, outputDir, hugoImagePrefix).Do(supportListView)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.notion-md.yaml)")
	rootCmd.PersistentFlags().StringVarP(&rootPageID, "id", "i", "", "id of root page which contains subpages")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "notion token")
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "./output", "output directory of markdowns and images")
	rootCmd.PersistentFlags().StringVarP(&hugoImagePrefix, "prefix", "p", "", "hugo markdown image prefix (relative path to image folder)")
	rootCmd.PersistentFlags().BoolVarP(&supportListView, "listview", "v", false, "use listview hold blogs, contain category, tags")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".notion-md" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".notion-md")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
