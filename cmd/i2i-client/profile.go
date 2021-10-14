package main

import (
	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func profileList(cmd *cobra.Command, args []string) {
	node, err := activeNode()
	if err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      node.Meta.APIToken,
		Keychain: node.Keychain,
	})

	list, err := i2iClient.ProfileList()
	if err != nil {
		fail(err)
	}

	printResult(list)
}

func profileAdd(cmd *cobra.Command, args []string) {
	node, err := activeNode()
	if err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      node.Meta.APIToken,
		Keychain: node.Keychain,
	})

	profile, err := i2iClient.DMeProfileAdd(&client.DMeProfileInput{ProfileName: args[0]})
	if err != nil {
		fail(err)
	}

	printResult(profile)
}

func profileUpdate(cmd *cobra.Command, args []string) {
	avatarURL, err := cmd.Flags().GetString(flagAvatarUrl)
	if err != nil {
		fail(err)
	}

	fileID, err := cmd.Flags().GetString(flagFileID)
	if err != nil {
		fail(err)
	}

	bio, err := cmd.Flags().GetString(flagBio)
	if err != nil {
		fail(err)
	}
	pseudonym, err := cmd.Flags().GetString(flagPseudonym)
	if err != nil {
		fail(err)
	}

	hideFirstName, err := cmd.Flags().GetBool(flagHideFirstName)
	if err != nil {
		fail(err)
	}

	hideSurname, err := cmd.Flags().GetBool(flagHideSurname)
	if err != nil {
		fail(err)
	}

	node, err := activeNode()
	if err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      node.Meta.APIToken,
		Keychain: node.Keychain,
	})

	input := &client.DMeProfileInput{
		ProfileName: args[0],
	}

	if hideSurname == true {
		input.HideSurname = &hideSurname
	}

	if hideFirstName == true {
		input.HideFirstName = &hideFirstName
	}

	if bio != "" {
		input.Bio = &bio
	}

	if avatarURL != "" {
		input.AvatarURL = &avatarURL
	}

	if fileID != "" {
		input.AvatarFileID = &fileID
	}

	if pseudonym != "" {
		input.Pseudonym = &pseudonym
	}

	profile, err := i2iClient.DMeProfileUpdate(input, "")
	if err != nil {
		fail(err)
	}

	printResult(profile)
}
