package main

import (
	"github.com/spf13/cobra"
)

const (
	flagInitializeAs    = "initialize-as"
	flagHosting         = "hosting"
	flagName            = "name"
	flagFirstName       = "first-name"
	flagSurname         = "surname"
	flagDescription     = "description"
	flagDuration        = "duration"
	flagPort            = "port"
	flagI2iPath         = "i2i-path"
	flagPlan            = "plan"
	flagPrivateScope    = "private-scope"
	flagPassword        = "password"
	flagProfile         = "profile"
	flagAvatarUrl       = "avatar-url"
	flagFileID          = "avatar-file-id"
	flagBio             = "bio"
	flagPseudonym       = "pseudonym"
	flagHideFirstName   = "hide-first-name"
	flagHideSurname     = "hide-surname"
	flagWireguardFormat = "wireguard-format"
	flagType            = "type"
	flagAs              = "as"
	flagAddress         = "address"
	flagDockerImage     = "docker-image"
	flagReply           = "reply"
	flagMethod          = "method"
)

func createCommandsStructure() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "i2i-client",
		Short: "i2i-client is client application for i2i",
		Long:  `i2i-client is client application for i2i`,
		Run:   nil,
	}

	rootCmd.AddCommand(createGroupCommand())
	rootCmd.AddCommand(createManagerCommand())
	rootCmd.AddCommand(createCfgCommand())
	rootCmd.AddCommand(createTuiCommand())
	rootCmd.AddCommand(createStateCommand())
	rootCmd.AddCommand(createInfoCommand())
	rootCmd.AddCommand(createAclCommand())
	rootCmd.AddCommand(createInitializeCommand())
	rootCmd.AddCommand(createFileCommand())
	rootCmd.AddCommand(createUnlockCommand())
	rootCmd.AddCommand(createConnectionCommand())
	rootCmd.AddCommand(createProfileCommand())
	rootCmd.AddCommand(createNodeCommand())
	rootCmd.AddCommand(createVpnCommand())
	rootCmd.AddCommand(createDmeCommand())
	rootCmd.AddCommand(createActionCommand())
	rootCmd.AddCommand(createResetCommand())
	rootCmd.AddCommand(createPlCommand())
	rootCmd.AddCommand(createPingCommand())
	rootCmd.AddCommand(createConversationsCommand())
	rootCmd.AddCommand(createBenchmarkCommand())
	rootCmd.AddCommand(createDirectMessageCommand())
	rootCmd.AddCommand(createMixinCommand())
	rootCmd.AddCommand(createPlanetConnectCommand())

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

func createBenchmarkCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "benchmark",
		Short: "run some benchmarks",
		Long:  `run some benchmarks`,
		Run:   benchmark,
	}
}

func createMixinCommand() *cobra.Command {
	mixinCmd := &cobra.Command{
		Use:   "mixin",
		Short: "manipulate mixin wallet",
		Long:  `manipulate mixin wallet`,
		Run:   nil,
	}

	setIDCmd := &cobra.Command{
		Use:   "set-id [value]",
		Args:  cobra.ExactArgs(1),
		Short: "set mixin id",
		Long:  `set mixin id"`,
		Run:   mixinSetId,
	}

	getIDCmd := &cobra.Command{
		Use:   "id",
		Short: "get mixin id",
		Long:  `get mixin id"`,
		Run:   mixinGetId,
	}

	mixinCmd.AddCommand(setIDCmd)
	mixinCmd.AddCommand(getIDCmd)

	return mixinCmd
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

	clientCmd.AddCommand(clientUpdateCmd)
	clientCmd.AddCommand(clientListCmd)

	planCmd.AddCommand(planListCmd)
	planCmd.AddCommand(planAddCmd)

	configCmd := &cobra.Command{
		Use:   "config",
		Short: "config show",
		Long:  `config show`,
		Run:   managerConfigShow,
	}

	configSetCmd := &cobra.Command{
		Use:   "set",
		Short: "update config value",
		Long:  `update config value`,
		Run:   managerConfigSet,
	}
	configSetCmd.Flags().String(flagAddress, "", "")
	configSetCmd.Flags().String(flagDockerImage, "", "")
	configCmd.AddCommand(configSetCmd)

	diagCmd := &cobra.Command{
		Use:   "diagnostic",
		Short: "run diagnostic",
		Long:  `run diagnostic`,
		Run:   managerDiagnostic,
	}

	managerCmd.AddCommand(diagCmd)
	managerCmd.AddCommand(configCmd)
	managerCmd.AddCommand(clientCmd)
	managerCmd.AddCommand(loginCmd)
	managerCmd.AddCommand(listCmd)
	managerCmd.AddCommand(versionGetCmd)
	managerCmd.AddCommand(planCmd)

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

	completion := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "generate completions for shell",
		Long:  `generate completions for shell`,
		Run:   completion,
	}

	cfgCmd.AddCommand(initCmd)
	cfgCmd.AddCommand(completion)

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

func createPingCommand() *cobra.Command {
	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "ping i2i node",
		Long:  `ping i2i node`,
		Run:   info,
	}

	return pingCmd
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

func createInitializeCommand() *cobra.Command {
	initializeCmd := &cobra.Command{
		Use:       "initialize [type]",
		ValidArgs: []string{"dme", "dorg", "supernode"},
		Args:      cobra.ExactArgs(1),
		Short:     "initialize node as dme|dorg|supernode",
		Long:      `run i2i on local machine`,
		Run:       initialize,
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
		Use:   "add [profile] [public key]",
		Args:  cobra.ExactArgs(2),
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

	connectionListAnonymousCmd := &cobra.Command{
		Use:   "list-anonymous",
		Short: "list i2i anonymous connections (contacts)",
		Long:  `list i2i anonymous connections (contacts)`,
		Run:   connectionListAnonymous,
	}
	connectionListAnonymousCmd.Flags().String(flagProfile, "", "name of profile to display connections from")

	connectionCmd.AddCommand(connectionListAnonymousCmd)
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

	profileUpdateCmd := &cobra.Command{
		Use:   "update [name]",
		Args:  cobra.ExactArgs(1),
		Short: "update profile fields",
		Long:  `update profile fields`,
		Run:   profileUpdate,
	}

	profileUpdateCmd.Flags().String(flagAvatarUrl, "", "url of the avatar to use")
	profileUpdateCmd.Flags().String(flagFileID, "", "id of i2i hosted file to use")
	profileUpdateCmd.Flags().String(flagBio, "", "change bio info")
	profileUpdateCmd.Flags().String(flagPseudonym, "", "custom pseudonym")
	profileUpdateCmd.Flags().Bool(flagHideFirstName, false, "hide first name")
	profileUpdateCmd.Flags().Bool(flagHideSurname, false, "hide surname")

	profileListCmd := &cobra.Command{
		Use:   "list",
		Short: "list added profiles",
		Long:  `list added profiles`,
		Run:   profileList,
	}

	profileCmd.AddCommand(profileUpdateCmd)
	profileCmd.AddCommand(profileListCmd)
	profileCmd.AddCommand(profileAddCmd)

	return profileCmd
}

func createNodeCommand() *cobra.Command {
	nodeCmd := &cobra.Command{
		Use:   "node",
		Short: "manage i2i nodes",
		Long:  `manage i2i nodes`,
		Run:   nodeInfo,
	}

	setActive := &cobra.Command{
		Use:   "set-active [node]",
		Short: "set active node",
		Long:  `set active node`,
		Args:  cobra.ExactArgs(1),
		Run:   nodeSetActive,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list configured nodes",
		Long:  `list configured nodes`,
		Run:   nodeList,
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [name]",
		Args:  cobra.ExactArgs(1),
		Short: "remove node from configuration",
		Long:  `remove node from configuration`,
		Run:   nodeRemove,
	}

	showCmd := &cobra.Command{
		Use:   "show [name]",
		Args:  cobra.ExactArgs(1),
		Short: "show details about node",
		Long:  `show details about node`,
		Run:   nodeShow,
	}

	updateCmd := &cobra.Command{
		Use:   "update [name]",
		Short: "update version of remote node",
		Long:  `update version of remote node`,
		Run:   nodeUpdate,
	}

	addCmd := &cobra.Command{
		Use:   "add [name] [ip:port]",
		Args:  cobra.ExactArgs(2),
		Short: "add new node",
		Long:  `add new node`,
		Run:   nodeAdd,
	}

	execCmd := &cobra.Command{
		Use:   "exec",
		Short: "run i2i node locally",
		Long:  `run i2i node locally`,
		Run:   nodeExec,
	}
	execCmd.Flags().String(flagName, "", "name of the local i2i")
	execCmd.Flags().String(flagI2iPath, "", "path to i2i executable")
	execCmd.Flags().Int(flagPort, 9090, "graphql listener port")

	orderCmd := &cobra.Command{
		Use:   "order",
		Short: "orders and initializes i2i",
		Long:  `orders and initializes i2i`,
		Run:   nodeOrder,
	}
	orderCmd.Flags().String(flagInitializeAs, "DME", "initialize ordered i2i as [DME|DORG]")
	orderCmd.Flags().String(flagName, "", "local name of ordered i2i instance")
	orderCmd.MarkFlagRequired(flagName)
	orderCmd.Flags().String(flagHosting, "", "hosting provider address")
	orderCmd.MarkFlagRequired(flagHosting)
	orderCmd.Flags().String(flagPlan, "", "hosting plan to use")
	orderCmd.MarkFlagRequired(flagPlan)
	orderCmd.Flags().String(flagPassword, "password_0123456789", "client password")

	nodeCmd.AddCommand(setActive)
	nodeCmd.AddCommand(listCmd)
	nodeCmd.AddCommand(showCmd)
	nodeCmd.AddCommand(deleteCmd)
	nodeCmd.AddCommand(addCmd)
	nodeCmd.AddCommand(execCmd)
	nodeCmd.AddCommand(orderCmd)
	nodeCmd.AddCommand(updateCmd)

	return nodeCmd
}

func createVpnCommand() *cobra.Command {
	vpnCmd := &cobra.Command{
		Use:   "vpn",
		Short: "manage vpn provided by i2i",
		Long:  `manage vpn provided by i2i`,
		Run:   nil,
	}

	peerCmd := &cobra.Command{
		Use:   "peer",
		Short: "manage vpn peers",
		Long:  `manage vpn peers`,
		Run:   nil,
	}

	addCmd := &cobra.Command{
		Use:   "add [network] [peer-name]",
		Args:  cobra.ExactArgs(2),
		Short: "add new vpn peer",
		Long:  `add new vpn peer`,
		Run:   vpnCreatePeerConfig,
	}
	addCmd.Flags().Bool(flagWireguardFormat, false, "generate wireguard format config")

	vpnStartCmd := &cobra.Command{
		Use:   "start [network]",
		Args:  cobra.ExactArgs(1),
		Short: "start vpn service",
		Long:  `start vpn service`,
		Run:   vpnStart,
	}

	vpnStopCmd := &cobra.Command{
		Use:   "stop [network]",
		Args:  cobra.ExactArgs(1),
		Short: "stop vpn service",
		Long:  `stop vpn service`,
		Run:   vpnStop,
	}

	vpnCreateCmd := &cobra.Command{
		Use:   "create [network]",
		Args:  cobra.ExactArgs(1),
		Short: "create vpn service",
		Long:  `create vpn service`,
		Run:   vpnCreate,
	}

	peerCmd.AddCommand(addCmd)
	vpnCmd.AddCommand(peerCmd)
	vpnCmd.AddCommand(vpnCreateCmd)
	vpnCmd.AddCommand(vpnStopCmd)
	vpnCmd.AddCommand(vpnStartCmd)

	return vpnCmd
}

func createActionCommand() *cobra.Command {
	actionCmd := &cobra.Command{
		Use:   "action",
		Short: "interactive actions",
		Long:  `interactive actions`,
		Run:   nil,
	}

	actionListCmd := &cobra.Command{
		Use:   "list",
		Short: "list interactive actions",
		Long:  `list interactive actions`,
		Run:   actionList,
	}
	actionListCmd.Flags().String(flagType, "", "action type expand")

	actionUpdateCmd := &cobra.Command{
		Use:   "update [id] [value]",
		Args:  cobra.ExactArgs(2),
		Short: "reply to interactive action",
		Long:  `reply to interactive action`,
		Run:   actionUpdate,
	}

	actionCmd.AddCommand(actionListCmd)
	actionCmd.AddCommand(actionUpdateCmd)

	return actionCmd
}

func createDmeCommand() *cobra.Command {
	dmeCmd := &cobra.Command{
		Use:   "dme",
		Short: "manage dme type of i2i",
		Long:  `manage dme type of i2i`,
		Run:   dmeInfo,
	}

	dmeUpdate := &cobra.Command{
		Use:   "update",
		Short: "update dme variables",
		Long:  `update dme variables`,
		Run:   dmeUpdate,
	}

	dmeUpdate.Flags().String(flagFirstName, "", "set first name")
	dmeUpdate.Flags().String(flagSurname, "", "set surname")

	dmeCmd.AddCommand(dmeUpdate)

	return dmeCmd
}

func createResetCommand() *cobra.Command {
	resetCmd := &cobra.Command{
		Use:   "reset",
		Short: "clean i2i database",
		Long:  `clean i2i database`,
		Run:   reset,
	}

	return resetCmd
}

func createPlCommand() *cobra.Command {
	plCmd := &cobra.Command{
		Use:   "pl",
		Short: "planet language operation",
		Long:  `planet language operation`,
		Run:   nil,
	}

	scopeCmd := &cobra.Command{
		Use:   "scope",
		Short: "scopes manipulation",
		Long:  `scopes manipulation`,
		Run:   nil,
	}

	scopeListCmd := &cobra.Command{
		Use:   "list",
		Short: "list planet language scopes",
		Long:  `list planet language scopes`,
		Run:   scopeList,
	}

	instanceCmd := &cobra.Command{
		Use:   "instance [scope] [id]",
		Short: "read instance",
		Long:  `read instance`,
		Run:   instance,
	}

	instanceListCmd := &cobra.Command{
		Use:   "list [scope]",
		Short: "list planet language instances",
		Long:  `list planet language instances`,
		Run:   instancesList,
	}
	instanceListCmd.Flags().String(flagAs, "", "filter instances by AS metadata")

	relationCmd := &cobra.Command{
		Use:   "relation",
		Short: "relations manipulation",
		Long:  `relations manipulation`,
		Run:   nil,
	}

	relationListCmd := &cobra.Command{
		Use:   "list [scope]",
		Short: "list planet language relations",
		Long:  `list planet language relations`,
		Run:   relationsList,
	}
	relationListCmd.Flags().String(flagAs, "", "filter relations by AS metadata")

	verifyCmd := &cobra.Command{
		Use:   "verify",
		Short: "verify planet language objects",
		Long:  `verify planet language objects`,
		Run:   plVerify,
	}

	instanceCmd.AddCommand(instanceListCmd)
	relationCmd.AddCommand(relationListCmd)

	scopeCmd.AddCommand(scopeListCmd)

	plCmd.AddCommand(instanceCmd)
	plCmd.AddCommand(relationCmd)
	plCmd.AddCommand(scopeCmd)
	plCmd.AddCommand(verifyCmd)

	return plCmd
}

func createConversationsCommand() *cobra.Command {
	conversationsCmd := &cobra.Command{
		Use:   "conversations",
		Short: "show latest conversations",
		Long:  `show latest conversations`,
		Run:   conversations,
	}
	return conversationsCmd
}

func createGroupCommand() *cobra.Command {
	groupCmd := &cobra.Command{
		Use:   "group",
		Short: "group chat manipulation",
		Long:  `group chat manipulation`,
		Run:   nil,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "lists group chats",
		Long:  `lists group chats`,
		Run:   groupList,
	}

	createCmd := &cobra.Command{
		Use:   "create [name]",
		Args:  cobra.ExactArgs(1),
		Short: "create new group chat",
		Long:  `create new group chat`,
		Run:   groupCreate,
	}

	addParticipantCmd := &cobra.Command{
		Use:   "add-participant [group-id] [participant-key]",
		Args:  cobra.ExactArgs(2),
		Short: "add new participant to group",
		Long:  `add new participant to group`,
		Run:   groupAddParticipant,
	}

	leaveCmd := &cobra.Command{
		Use:   "leave [group-id]",
		Args:  cobra.ExactArgs(1),
		Short: "leave group",
		Long:  `leaveGroup`,
		Run:   groupLeave,
	}

	messageCmd := &cobra.Command{
		Use:   "message",
		Short: "message manipulation",
		Long:  `message manipulation`,
		Run:   nil,
	}

	messageShowCmd := &cobra.Command{
		Use:   "show [id]",
		Args:  cobra.ExactArgs(1),
		Short: "show messages from group chat",
		Long:  `show messages from group chat`,
		Run:   groupMessageShow,
	}

	messageSendCmd := &cobra.Command{
		Use:   "send [id] [content]",
		Args:  cobra.ExactArgs(2),
		Short: "send message to group",
		Long:  `send message to group`,
		Run:   groupMessageSend,
	}
	messageCmd.AddCommand(messageShowCmd)
	messageCmd.AddCommand(messageSendCmd)

	groupCmd.AddCommand(listCmd)
	groupCmd.AddCommand(createCmd)
	groupCmd.AddCommand(leaveCmd)
	groupCmd.AddCommand(addParticipantCmd)
	groupCmd.AddCommand(messageCmd)

	return groupCmd
}

func createDirectMessageCommand() *cobra.Command {
	dmCmd := &cobra.Command{
		Use:   "direct-message",
		Short: "direct message manipulation",
		Long:  `direct message manipulation`,
		Run:   nil,
	}

	dmSendCmd := &cobra.Command{
		Use:   "send [destination] [content]",
		Args:  cobra.ExactArgs(2),
		Short: "send direct message",
		Long:  `send direct message`,
		Run:   dmSend,
	}
	dmSendCmd.Flags().String(flagReply, "", "reply to message with id")

	dmViewCmd := &cobra.Command{
		Use:   "view [conversation]",
		Args:  cobra.ExactArgs(1),
		Short: "view direct message",
		Long:  `view direct message`,
		Run:   dmView,
	}

	dmCmd.AddCommand(dmSendCmd)
	dmCmd.AddCommand(dmViewCmd)

	return dmCmd
}

func createPlanetConnectCommand() *cobra.Command {
	pcCmd := &cobra.Command{
		Use:   "pc",
		Short: "planet connect functions",
		Long:  `planet connect functions`,
		Run:   nil,
	}

	pcCmd.PersistentFlags().String(flagAddress, "http://localhost:20222", "address of the planet connect")

	pcCmd.AddCommand(createPcAuthCommand())
	pcCmd.AddCommand(createPcDataCommand())
	pcCmd.AddCommand(createPcTableCommand())
	pcCmd.AddCommand(createPcManagerCommand())
	pcCmd.AddCommand(createPcPingCommand())
	pcCmd.AddCommand(createPcCapabilitiesCommand())
	pcCmd.AddCommand(createUserCommand())

	return pcCmd
}

func createPcPingCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "authorized ping request to planet connect",
		Long:  `authorized ping request to planet connect`,
		Run:   pcPing,
	}
}

func createUserCommand() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "manage user details",
		Long:  `manage user details`,
		Run:   nil,
	}

	userInfoCmd := &cobra.Command{
		Use:   "info",
		Short: "show user info",
		Long:  `show user info`,
		Run:   userInfo,
	}

	userExchangeKeyCmd := &cobra.Command{
		Use:   "exchange-key [login]",
		Args:  cobra.ExactArgs(1),
		Short: "get other user exchange key",
		Long:  `get other user exchange key`,
		Run:   userExchangeKey,
	}

	userCmd.AddCommand(userInfoCmd)
	userCmd.AddCommand(userExchangeKeyCmd)

	return userCmd
}

func createPcCapabilitiesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "capabilities",
		Short: "show planet connect capabilities",
		Long:  `show planet connect capabilities`,
		Run:   pcCapabilities,
	}
}

func createPcManagerCommand() *cobra.Command {
	managerCmd := &cobra.Command{
		Use:   "manager",
		Short: "i2i-manager integration",
		Long:  `i2i-manager integration`,
		Run:   nil,
	}

	orderCmd := &cobra.Command{
		Use:   "order [name]",
		Args:  cobra.ExactArgs(1),
		Short: "order i2i node",
		Long:  `order i2i node`,
		Run:   pcNodeOrder,
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [name]",
		Args:  cobra.ExactArgs(1),
		Short: "delete i2i node",
		Long:  `delete i2i node`,
		Run:   pcNodeDelete,
	}

	showCmd := &cobra.Command{
		Use:   "show [name]",
		Args:  cobra.ExactArgs(1),
		Short: "show i2i node",
		Long:  `show i2i node`,
		Run:   pcNodeShow,
	}

	managerCmd.AddCommand(orderCmd)
	managerCmd.AddCommand(deleteCmd)
	managerCmd.AddCommand(showCmd)

	return managerCmd
}

func createPcAuthCommand() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "manager user authorization",
		Long:  `manager user authorization`,
		Run:   nil,
	}

	registerCmd := &cobra.Command{
		Use:   "register [login] [secret]",
		Args:  cobra.ExactArgs(2),
		Short: "register new user to planet connect",
		Long:  `register new user to planet connect`,
		Run:   pcRegister,
	}
	registerCmd.Flags().String(flagMethod, "email", "verification method to use")

	loginCmd := &cobra.Command{
		Use:   "login [login] [secret]",
		Args:  cobra.ExactArgs(2),
		Short: "login existing user",
		Long:  `login existing user`,
		Run:   pcLogin,
	}

	secureRandomCmd := &cobra.Command{
		Use:   "secure-random [secret]",
		Args:  cobra.ExactArgs(1),
		Short: "obtain secure random",
		Long:  `obtain secure random`,
		Run:   pcSecureRandom,
	}

	verifyCmd := &cobra.Command{
		Use:   "verify",
		Short: "verify authorization",
		Long:  `verify authorization`,
		Run:   pcVerify,
	}

	authCmd.AddCommand(verifyCmd)
	authCmd.AddCommand(registerCmd)
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(secureRandomCmd)

	return authCmd
}

func createPcDataCommand() *cobra.Command {
	dataCmd := &cobra.Command{
		Use:   "data",
		Short: "user data manipulation",
		Long:  `user data manipulation`,
		Run:   nil,
	}

	dataAddCmd := &cobra.Command{
		Use:   "add [table] [key] [value]",
		Args:  cobra.ExactArgs(3),
		Short: "add new data to user database",
		Long:  `add new data to user database`,
		Run:   pcDataAdd,
	}

	dataGetCmd := &cobra.Command{
		Use:   "get [table] [key]",
		Args:  cobra.ExactArgs(2),
		Short: "get value of user data identified by key",
		Long:  `get value of user data identified by key`,
		Run:   pcDataGet,
	}

	dataDeleteCmd := &cobra.Command{
		Use:   "delete [table] [key]",
		Args:  cobra.ExactArgs(2),
		Short: "delete existing user data identified by key",
		Long:  `delete existing user data identified by key`,
		Run:   pcDataDelete,
	}

	dataUpdateCmd := &cobra.Command{
		Use:   "update [table] [key] [value]",
		Args:  cobra.ExactArgs(3),
		Short: "update existing data with new value",
		Long:  `update existing data with new value`,
		Run:   pcDataUpdate,
	}

	dataListCmd := &cobra.Command{
		Use:   "list [table...]",
		Short: "list all data stored by the user in given tables. If no table is selected all data is returned.",
		Long:  `list all data stored by the user in given tables. If no table is selected all data is returned.`,
		Run:   pcDataListCmd,
	}

	dataCmd.AddCommand(dataAddCmd)
	dataCmd.AddCommand(dataGetCmd)
	dataCmd.AddCommand(dataDeleteCmd)
	dataCmd.AddCommand(dataUpdateCmd)
	dataCmd.AddCommand(dataListCmd)

	return dataCmd
}

func createPcTableCommand() *cobra.Command {
	tableCmd := &cobra.Command{
		Use:   "table",
		Short: "user table manipulation",
		Long:  `user table manipulation`,
		Run:   nil,
	}

	tableListCmd := &cobra.Command{
		Use:   "list",
		Short: "list user data tables",
		Long:  `list user data tables`,
		Run:   pcTableList,
	}

	tableCmd.AddCommand(tableListCmd)

	return tableCmd
}
