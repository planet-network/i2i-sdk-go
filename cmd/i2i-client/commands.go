package main

import (
	"github.com/spf13/cobra"
)

const (
	flagSkipInitialize = "skip-initialize"
	flagInitializeAs   = "initialize-as"
	flagHosting        = "hosting"
)

func createCommandsStructure() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "i2i-client",
		Short: "i2i-client is client application for i2i",
		Long:  `i2i-client is client application for i2i`,
		Run:   nil,
	}

	rootCmd.AddCommand(createManagerCommand())
	rootCmd.AddCommand(createCfgCommand())
	rootCmd.AddCommand(createTuiCommand())

	return rootCmd
}

func createTuiCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tui",
		Short: "run terminal ui",
		Long:  `run terminal ui`,
		Run:   tui,
	}
}

func createManagerCommand() *cobra.Command {
	managerCmd := &cobra.Command{
		Use:   "manager",
		Short: "interact with i2i manager",
		Long:  `interact with i2i manager`,
		Run:   nil,
	}

	orderCmd := &cobra.Command{
		Use:   "order [name]",
		Short: "orders new i2i",
		Long:  `orders new i2i`,
		Run:   order,
	}

	orderCmd.Flags().Bool(flagSkipInitialize, false, "skip i2i initialization, order only")
	orderCmd.Flags().String(flagInitializeAs, "DME", "initialize ordered i2i as [DME|DORG]")
	orderCmd.Flags().String(flagHosting, "", "hosting provider address")

	plansCmd := &cobra.Command{
		Use:   "plans",
		Short: "show available plans",
		Long:  `show available plans`,
		Run:   managerPlans,
	}

	plansCmd.Flags().String(flagHosting, "", "hosting provider address")
	plansCmd.MarkFlagRequired(flagHosting)

	managerCmd.AddCommand(orderCmd)
	managerCmd.AddCommand(plansCmd)

	return managerCmd
}

func createCfgCommand() *cobra.Command {
	cfgCmd := &cobra.Command{
		Use:   "cfg",
		Short: "manage local configuration",
		Long:  `manage local configuration`,
		Run:   nil,
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "initialize local configuration",
		Long:  `initialize local configuration`,
		Run:   cfgInit,
	}

	cfgCmd.AddCommand(initCmd)

	return cfgCmd
}
