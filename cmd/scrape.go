/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmaitlandsoto/luna-cli/internal/file"
	"github.com/jmaitlandsoto/luna-cli/internal/scrape"
	"github.com/spf13/cobra"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape the text from a URL",
	Long:  `1`,
	Run:   scrapeRun,
}

func scrapeRun(cmd *cobra.Command, args []string) {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	for i, url := range args {
		text, err := scrapeAndParse(url)
		if err != nil {
			fmt.Println("Ran into an error")
			break
		}

		err = file.CreateTxt(strconv.Itoa(i), cwd, text)
		if err != nil {
			fmt.Printf("Error creating Txt file: %v\n", err)
			break
		}
	}
	fmt.Println("\n---------- scrape complete ----------")
}

func scrapeAndParse(url string) (string, error) {

	siteContent, err := scrape.ScrapeHTML(url)
	if err != nil {
		return "", err
	}

	text, err := scrape.ParseHTML(siteContent)
	if err != nil {
		return "", err
	}

	fmt.Println(text)
	return text, nil
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
