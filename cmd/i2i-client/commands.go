package main

import (
	"github.com/spf13/cobra"
)

const (
	flagInitializeAs = "initialize-as"
	flagHosting      = "hosting"
	flagName         = "name"
	flagDescription  = "description"
	flagDuration     = "duration"
	flagPort         = "port"
	flagI2iPath      = "i2i-path"
	flagPlan         = "plan"
	flagPrivateScope = "private-scope"
	flagPassword     = "password"
	flagProfile      = "profile"
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
	rootCmd.AddCommand(createInfoCommand())
	rootCmd.AddCommand(createAclCommand())
	rootCmd.AddCommand(createExecCommand())
	rootCmd.AddCommand(createInitializeCommand())
	rootCmd.AddCommand(createFileCommand())
	rootCmd.AddCommand(createUnlockCommand())
	rootCmd.AddCommand(createConnectionCommand())
	rootCmd.AddCommand(createProfileCommand())

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

func createUnlockCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "unlock",
		Short: "unlocks i2i",
		Long:  `unlocks i2i`,
		Run:   unlock,
	}
}

func createFileCommand() *cobra.Command {
	fileCmd := &cobra.Command{
		Use:   "file",
		Short: "file management command",
		Long:  `file management command`,
		Run:   nil,
	}

	fileUploadCmd := &cobra.Command{
		Use:   "up [local-file]",
		Args:  cobra.ExactArgs(1),
		Short: "upload local file to i2i",
		Long:  `upload local file to i2i`,
		Run:   fileUpload,
	}

	fileDownloadCmd := &cobra.Command{
		Use:   "dl [id] [directory]",
		Args:  cobra.ExactArgs(2),
		Short: "download file from i2i to local directory",
		Long:  `download file from i2i to local directory`,
		Run:   fileDownload,
	}

	fileListCmd := &cobra.Command{
		Use:   "ls",
		Short: "list files hosted on i2i",
		Long:  `list files hosted on i2i`,
		Run:   fileList,
	}

	fileRemoveCmd := &cobra.Command{
		Use:   "rm [id]",
		Args:  cobra.ExactArgs(1),
		Short: "remove file from i2i node",
		Long:  `remove file form i2i node`,
		Run:   fileRemove,
	}

	fileStatCmd := &cobra.Command{
		Use:   "stat [id]",
		Args:  cobra.ExactArgs(1),
		Short: "show information about file",
		Long:  `show information about file`,
		Run:   fileStat,
	}

	fileTransferCmd := &cobra.Command{
		Use:   "transfer [id] [connection]",
		Args:  cobra.ExactArgs(2),
		Short: "transfer file to connection",
		Long:  `transfer file to connection`,
		Run:   fileTransfer,
	}

	fileRenameCmd := &cobra.Command{
		Use:   "rename [id] [name]",
		Args:  cobra.ExactArgs(2),
		Short: "rename remote file",
		Long:  `rename remote file`,
		Run:   fileRename,
	}

	fileCmd.AddCommand(fileUploadCmd)
	fileCmd.AddCommand(fileStatCmd)
	fileCmd.AddCommand(fileDownloadCmd)
	fileCmd.AddCommand(fileListCmd)
	fileCmd.AddCommand(fileRemoveCmd)
	fileCmd.AddCommand(fileTransferCmd)
	fileCmd.AddCommand(fileRenameCmd)

	return fileCmd
}

func createManagerCommand() *cobra.Command {
	managerCmd := &cobra.Command{
		Use:   "manager",
		Short: "interact with i2i manager",
		Long:  `interact with i2i manager`,
		Run:   nil,
	}

	nodeUpdateCmd := &cobra.Command{
		Use:   "node-update",
		Short: "update remote node",
		Long:  `update remote node`,
		Run:   managerNodeUpdate,
	}

	quickOrderCmd := &cobra.Command{
		Use:   "quick-order",
		Short: "orders and initializes i2i",
		Long:  `orders and initializes i2i`,
		Run:   managerQuickOrder,
	}

	versionGetCmd := &cobra.Command{
		Use:   "version",
		Short: "show current version of hosted i2i",
		Long:  `show current version of hosted i2i`,
		Run:   versionShow,
	}

	versionSetCmd := &cobra.Command{
		Use:   "set",
		Short: "orders and initializes i2i",
		Long:  `orders and initializes i2i`,
		Run:   versionSet,
	}
	versionGetCmd.AddCommand(versionSetCmd)

	loginCmd := &cobra.Command{
		Use:   "login [address] [user] [pass]",
		Args:  cobra.ExactArgs(3),
		Short: "login to i2i-manager",
		Long:  `login to i2i-manager`,
		Run:   managerLogin,
	}

	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "manage i2i-manager clients",
		Long:  `manage i2i-manager clients`,
		Run:   nil,
	}

	clientUpdateCmd := &cobra.Command{
		Use:   "update [id]",
		Args:  cobra.ExactArgs(1),
		Short: "update client i2i",
		Long:  `update client i2i`,
		Run:   clientUpdate,
	}

	clientListCmd := &cobra.Command{
		Use:   "list",
		Short: "list clients",
		Long:  `list clients`,
		Run:   clientList,
	}
	clientListCmd.Flags().Int(flagPort, 0, "show only client with given port")

	planCmd := &cobra.Command{
		Use:   "plan",
		Short: "manage i2i-manager plans",
		Long:  `manage i2i-manager plans`,
		Run:   nil,
	}

	planAddCmd := &cobra.Command{
		Use:   "add [name]",
		Args:  cobra.ExactArgs(1),
		Short: "create new plan",
		Long:  `create new plan`,
		Run:   planAdd,
	}

	planAddCmd.Flags().String(flagDescription, "", "plan description")
	planAddCmd.Flags().Int64(flagDuration, 0, "duration in hours")
	planAddCmd.MarkFlagRequired(flagDuration)

	planListCmd := &cobra.Command{
		Use:   "list",
		Short: "list plans",
		Long:  `list plans`,
		Run:   planList,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list configured managers",
		Long:  `list configured managers`,
		Run:   managerList,
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

	clientCmd.AddCommand(clientUpdateCmd)
	clientCmd.AddCommand(clientListCmd)

	planCmd.AddCommand(planListCmd)
	planCmd.AddCommand(planAddCmd)
	managerCmd.AddCommand(clientCmd)
	managerCmd.AddCommand(loginCmd)
	managerCmd.AddCommand(listCmd)
	managerCmd.AddCommand(versionGetCmd)
	managerCmd.AddCommand(planCmd)
	managerCmd.AddCommand(nodeUpdateCmd)

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

	deleteCmd := &cobra.Command{
		Use:   "delete [name]",
		Short: "remove node from configuration",
		Long:  `remove node from configuration`,
		Run:   cfgDelete,
	}

	showCmd := &cobra.Command{
		Use:   "show [name]",
		Short: "show details about node",
		Long:  `show details about node`,
		Run:   cfgShow,
	}

	addCmd := &cobra.Command{
		Use:   "add [name] [ip:port]",
		Args:  cobra.ExactArgs(2),
		Short: "add new node",
		Long:  `add new node`,
		Run:   cfgAdd,
	}

	cfgCmd.AddCommand(addCmd)
	cfgCmd.AddCommand(initCmd)
	cfgCmd.AddCommand(setActive)
	cfgCmd.AddCommand(listCmd)
	cfgCmd.AddCommand(showCmd)
	cfgCmd.AddCommand(deleteCmd)

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

func createInfoCommand() *cobra.Command {
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "show i2i internal details",
		Long:  `show i2i internal details`,
		Run:   info,
	}

	return infoCmd
}

func createAclCommand() *cobra.Command {
	aclCmd := &cobra.Command{
		Use:   "acl",
		Short: "manage i2i acl",
		Long:  `manage i2i acl`,
		Run:   nil,
	}

	aclAddCmd := &cobra.Command{
		Use:   "add",
		Short: "create new acl",
		Long:  `create new acl`,
		Run:   aclAdd,
	}

	aclAddCmd.Flags().String(flagPrivateScope, "", "name of private scope to use")
	aclAddCmd.Flags().String(flagName, "i2i-sdk-go", "name of the acl")

	aclListCmd := &cobra.Command{
		Use:   "list",
		Short: "manage i2i acl",
		Long:  `manage i2i acl`,
		Run:   aclList,
	}

	aclDeleteCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "remove acl",
		Long:  `remove alc`,
		Run:   aclRemove,
	}

	aclCmd.AddCommand(aclListCmd)
	aclCmd.AddCommand(aclAddCmd)
	aclCmd.AddCommand(aclDeleteCmd)

	return aclCmd
}

func createExecCommand() *cobra.Command {
	execCmd := &cobra.Command{
		Use:   "exec",
		Short: "run i2i on local machine",
		Long:  `run i2i on local machine`,
		Run:   execute,
	}

	execCmd.Flags().String(flagName, "", "name of the local i2i")
	execCmd.Flags().String(flagI2iPath, "", "path to i2i executable")
	execCmd.Flags().Int(flagPort, 9090, "graphql listener port")

	return execCmd
}

func createInitializeCommand() *cobra.Command {
	initializeCmd := &cobra.Command{
		Use:   "initialize [type]",
		Args:  cobra.ExactArgs(1),
		Short: "initialize node as DME|DORG|SUPERNODE",
		Long:  `run i2i on local machine`,
		Run:   initialize,
	}

	return initializeCmd
}

func createConnectionCommand() *cobra.Command {
	connectionCmd := &cobra.Command{
		Use:   "connection",
		Short: "manage i2i connections (contacts)",
		Long:  `manage i2i connections (contacts)`,
		Run:   nil,
	}

	connectionAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add new i2i connection (contact)",
		Long:  `add new i2i connection (contact)`,
		Run:   connectionAdd,
	}

	connectionAddLocalCmd := &cobra.Command{
		Use:   "add-by-cfg [name1] [name2]",
		Args:  cobra.ExactArgs(2),
		Short: "connects two i2i from local cfg",
		Long:  `connects two i2i from local cfg`,
		Run:   connectionAddLocal,
	}

	connectionListCmd := &cobra.Command{
		Use:   "list",
		Short: "list i2i connections (contacts)",
		Long:  `list i2i connections (contacts)`,
		Run:   connectionList,
	}
	connectionListCmd.Flags().String(flagProfile, "", "name of profile to display connections from")

	connectionCmd.AddCommand(connectionAddCmd)
	connectionCmd.AddCommand(connectionListCmd)
	connectionCmd.AddCommand(connectionAddLocalCmd)

	return connectionCmd
}

func createProfileCommand() *cobra.Command {
	profileCmd := &cobra.Command{
		Use:   "profile",
		Short: "manage i2i profiles",
		Long:  `manage i2i profiles`,
		Run:   nil,
	}

	profileAddCmd := &cobra.Command{
		Use:   "add [name]",
		Args:  cobra.ExactArgs(1),
		Short: "add new profile",
		Long:  `add new profile`,
		Run:   profileAdd,
	}

	profileListCmd := &cobra.Command{
		Use:   "list",
		Short: "list added profiles",
		Long:  `list added profiles`,
		Run:   profileList,
	}

	profileCmd.AddCommand(profileListCmd)
	profileCmd.AddCommand(profileAddCmd)

	return profileCmd
}
