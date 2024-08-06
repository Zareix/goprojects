/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	currency "zareix/goprojects/05-currency/internal"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type Currency struct {
	Name string
	Id   string
}

var CURRENCIES = [...]Currency{
	{"$ United States Dollar", "USD"},
	{"€ Euro", "EUR"},
	{"£ Pound Sterling", "GBP"},
	{"¥ Japanese Yen", "JPY"},
	{"$ Canadian Dollar", "CAD"},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "05-currency",
	Short: "A simple currency converter",
	Long:  `A simple currency converter that can convert between different currencies.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currencyFrom := cmd.Flags().Lookup("from").Value.String()
		currencyTo := cmd.Flags().Lookup("to").Value.String()
		if currencyFrom == "" || currencyTo == "" {
			var options []huh.Option[string]
			for _, currency := range CURRENCIES {
				options = append(options, huh.NewOption(currency.Name, currency.Id))
			}
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Currency to convert from").
						Options(
							options...,
						).
						Value(&currencyFrom),

					huh.NewSelect[string]().
						Title("Currency to convert to").
						Options(
							options...,
						).
						Value(&currencyTo),
				),
			)
			err := form.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("Invalid amount")
			os.Exit(1)
		}

		usdAmount, err := currency.Convert(amount, currencyFrom, currencyTo)
		if err != nil {
			fmt.Println("Error converting currency")
			os.Exit(1)
		}
		fmt.Println(amount, currencyFrom, "is equal to", usdAmount, currencyTo)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.05-currency.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("to", "t", "", "Currency to convert to")
	rootCmd.Flags().StringP("from", "f", "", "Currency to convert from")
}
