package main

import (
	"fmt"

	"github.com/planet-platform/i2i-sdk-go/app"
	"github.com/planet-platform/i2i-sdk-go/manager"
	"github.com/spf13/cobra"
)

func clientList(cmd *cobra.Command, args []string) {
	mng, err := activeManager()
	if err != nil {
		fail(err)
	}

	managerClient := manager.NewFeClient(manager.FeOpt{
		Address:  mng.Address,
		Login:    mng.User,
		Password: mng.Password,
		Jwt:      mng.Token,
	})

	clients, err := managerClient.ClientList()
	if err != nil {
		fail(err)
	}

	port, err := cmd.Flags().GetInt(flagPort)
	if err != nil {
		fail(err)
	}

	if port == 0 {
		printResult(clients)
	}

	for i := range clients {
		if clients[i].Instance == nil {
			continue
		}

		if clients[i].Instance.HostNetworkPort == port {
			printResult(clients[i])
			return
		}

		if clients[i].Instance.HostAPIPort == port {
			printResult(clients[i])
			return
		}
	}

}

func clientUpdate(cmd *cobra.Command, args []string) {
	mng, err := activeManager()
	if err != nil {
		fail(err)
	}

	managerClient := manager.NewFeClient(manager.FeOpt{
		Address:  mng.Address,
		Login:    mng.User,
		Password: mng.Password,
		Jwt:      mng.Token,
	})

	if err := managerClient.ClientUpdate(args[0]); err != nil {
		fail(err)
	}
}

func versionShow(cmd *cobra.Command, args []string) {
	mng, err := activeManager()
	if err != nil {
		fail(err)
	}

	managerClient := manager.NewFeClient(manager.FeOpt{
		Address:  mng.Address,
		Login:    mng.User,
		Password: mng.Password,
		Jwt:      mng.Token,
	})

	version, err := managerClient.VersionShow()
	if err != nil {
		fail(err)
	}

	printResult(&manager.FeVersion{Version: version})
}

func versionSet(cmd *cobra.Command, args []string) {
	mng, err := activeManager()
	if err != nil {
		fail(err)
	}

	managerClient := manager.NewFeClient(manager.FeOpt{
		Address:  mng.Address,
		Login:    mng.User,
		Password: mng.Password,
		Jwt:      mng.Token,
	})

	if err := managerClient.VersionSet(args[0]); err != nil {
		fail(err)
	}
}

func managerLogin(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	managerClient := manager.NewFeClient(manager.FeOpt{
		Address:  args[0],
		Login:    args[1],
		Password: args[2],
	})

	if err := managerClient.Login(); err != nil {
		fail(err)
	}

	if err := appHandler.ManagerUpdate(&app.Manager{
		Address:  args[0],
		Token:    managerClient.Jwt(),
		User:     args[1],
		Password: args[2],
	}); err != nil {
		fail(err)
	}

}

func managerList(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	cfg := appHandler.Config()

	fmt.Println("managers:")
	for k, mng := range cfg.Managers {
		if k == cfg.SelectedManager {
			fmt.Printf(" * %s\n", managerString(mng))
		} else {
			fmt.Printf("   %s\n", managerString(mng))
		}
	}

}

func managerString(n *app.Manager) string {
	return fmt.Sprintf("%s", n.Address)
}

func activeManager() (*app.Manager, error) {
	appHandler, err := app.NewApp()
	if err != nil {
		return nil, err
	}

	if err := appHandler.LoadConfig(); err != nil {
		return nil, err
	}

	mng, err := appHandler.DefaultManager()
	if err != nil {
		return nil, err
	}

	return mng, nil
}

func planAdd(cmd *cobra.Command, args []string) {
	planDescription, err := cmd.Flags().GetString(flagDescription)
	if err != nil {
		fail(err)
	}

	planDuration, err := cmd.Flags().GetInt64(flagDuration)
	if err != nil {
		fail(err)
	}

	mng, err := activeManager()
	if err != nil {
		fail(err)
	}

	managerClient := manager.NewFeClient(manager.FeOpt{
		Address:  mng.Address,
		Login:    mng.User,
		Password: mng.Password,
		Jwt:      mng.Token,
	})

	err = managerClient.PlanAdd(&manager.FePlan{
		ID:          args[0],
		Name:        args[0],
		Description: planDescription,
		Currency:    "USD",
		Price:       0,
		Duration:    planDuration * 3600,
		StorageSize: 1024,
	})

	if err != nil {
		fail(err)
	}
}

func planList(cmd *cobra.Command, args []string) {
	mng, err := activeManager()
	if err != nil {
		fail(err)
	}

	managerClient := manager.NewFeClient(manager.FeOpt{
		Address:  mng.Address,
		Login:    mng.User,
		Password: mng.Password,
		Jwt:      mng.Token,
	})

	plans, err := managerClient.PlanList()
	if err != nil {
		fail(err)
	}

	printResult(plans)
}
