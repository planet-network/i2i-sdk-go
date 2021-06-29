package main

import (
	"github.com/spf13/cobra"
)

const (
	flagInitializeAs = "initialize-as"
	flagHosting      = "hosting"
	flagName         = "name"
	flagPlan         = "plan"
	flagPassword     = "password"
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
	rootCmd.AddCommand(createStateCommand())

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

	quickOrderCmd := &cobra.Command{
		Use:   "quick-order",
		Short: "orders and initializes i2i",
		Long:  `orders and initializes i2i`,
		Run:   managerQuickOrder,
	}

	quickOrderCmd.Flags().String(flagInitializeAs, "DME", "initialize ordered i2i as [DME|DORG]")
	quickOrderCmd.Flags().String(flagName, "", "local name of ordered i2i instance")
	quickOrderCmd.MarkFlagRequired(flagName)

	quickOrderCmd.Flags().String(flagHosting, "", "hosting provider address")
	quickOrderCmd.MarkFlagRequired(flagHosting)

	quickOrderCmd.Flags().String(flagPlan, "", "hosting plan to use")
	quickOrderCmd.MarkFlagRequired(flagPlan)

	quickOrderCmd.Flags().String(flagPassword, "password_0123456789", "client password")

	managerCmd.AddCommand(quickOrderCmd)

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

	setActive := &cobra.Command{
		Use:   "set-active [node]",
		Short: "set active node",
		Long:  `set active node`,
		Args:  cobra.ExactArgs(1),
		Run:   cfgSetActive,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list configured nodes",
		Long:  `list configured nodes`,
		Run:   cfgList,
	}

	showCmd := &cobra.Command{
		Use:   "show [name]",
		Short: "show details about node",
		Long:  `show details about node`,
		Run:   cfgShow,
	}

	cfgCmd.AddCommand(initCmd)
	cfgCmd.AddCommand(setActive)
	cfgCmd.AddCommand(listCmd)
	cfgCmd.AddCommand(showCmd)

	return cfgCmd
}

func createStateCommand() *cobra.Command {
	stateCmd := &cobra.Command{
		Use:   "state",
		Short: "show i2i state",
		Long:  `show i2i state`,
		Run:   state,
	}

	return stateCmd
}
