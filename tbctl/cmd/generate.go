/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/hanapedia/the-bench/tbctl/pkg/kubernetes/generate"
	"github.com/spf13/cobra"
)

var inputPath string
var outputPath string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate Kubernetes manifests for given service unit configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(inputPath) == "" {
			fmt.Println("Error: Missing -f flag or empty file path")
			return
		}

		fileInfo, err := os.Stat(inputPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if fileInfo.IsDir() {
			generate.GenerateFromDirectory(inputPath, outputPath)
		} else {
			manifestGenerator := generate.NewManifestGenerator(inputPath, outputPath)
			manifestGenerator.GenerateFromFile()
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVarP(&inputPath, "file", "f", "", "YAML file or directory to validate")
	generateCmd.PersistentFlags().StringVarP(&outputPath, "out", "o", "", "output directory for generated files")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
