package main

import (
	"encoding/hex"
	"fmt"
	"github.com/planet-network/i2i-sdk-go/pc"
	"github.com/planet-network/i2i-sdk-go/pc/cryptography"
	"github.com/planet-network/i2i-sdk-go/pc/models"
	"github.com/spf13/cobra"
	"os"
	"time"
)

const (
	authorizationEnv = "PC_CLI_AUTH"
	masterKeyEnv     = "PC_CLI_MASTER_KEY"
)

func createClient(cmd *cobra.Command, loadAuthFromEnv bool) *pc.RestClient {
	addr, err := cmd.Flags().GetString(flagAddress)
	if err != nil {
		fail("failed to read address flag:", err)
	}

	restClient, err := pc.NewRestClient(addr)
	if err != nil {
		fail("failed to create client:", err)
	}

	if loadAuthFromEnv {
		authorization := os.Getenv(authorizationEnv)
		if authorization == "" {
			fail(authorizationEnv, "env not set")
		}

		if err := verifyAuthorization(authorization); err != nil {
			fail("authorization verification failed:", err)
		}

		masterKeyString := os.Getenv(masterKeyEnv)
		if masterKeyString == "" {
			fail(masterKeyEnv, "env not set")
		}

		masterKey, err := cryptography.MasterKeyFromString(masterKeyString)
		if err != nil {
			fail("failed to parse", masterKeyEnv, "err:", err)
		}

		restClient.SetMasterKey(masterKey)
		restClient.SetAuthorization(authorization)
	}

	return restClient
}

func verifyAuthorization(auth string) error {
	return nil
}

func pcLogin(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, false)

	response, err := pcClient.Login(args[0], args[1])
	if err != nil {
		fail("failed:", err)
	}

	printForEval(args[1], response)
}

func pcRegister(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, false)

	method, err := cmd.Flags().GetString(flagMethod)
	if err != nil {
		fail("failed to get method flag:", err)
	}

	if _, err := pcClient.Register(args[0], args[1], method); err != nil {
		fail(err)
	}
}

func pcSecureRandom(cmd *cobra.Command, args []string) {

}

func printForEval(secret string, response *models.LoginResponse) {
	var (
		derivedPassword = cryptography.DerivedPassword([]byte(secret))
		preMasterKey    = cryptography.CalculatePreMasterKey(derivedPassword)
		masterPassword  = cryptography.CalculateMasterKey(preMasterKey, response.SecureRandom)
	)

	fmt.Printf("export %s=%s\n", authorizationEnv, response.Authorization)
	fmt.Printf("export %s=%s\n", masterKeyEnv, masterPassword.String())
}

func pcDataGet(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	data, err := pcClient.DataGet(args[0], args[1])
	if err != nil {
		fail("call failed:", err)
	}

	parsed, err := pcClient.ParseDataResponse(data)
	if err != nil {
		fail("failed to parse response:", err)
	}

	fmt.Println("Table             :", string(parsed.Table))
	fmt.Println("Key               :", string(parsed.Key))
	fmt.Println("Value             :", string(parsed.Value))
	fmt.Println("Creation time     :", time.Unix(parsed.CreatedAt, 0).Format(time.RFC822))
	fmt.Println("Modification time :", time.Unix(parsed.ModifiedAt, 0).Format(time.RFC822))
}

func pcDataAdd(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	err := pcClient.DataAdd(args[0], args[1], args[2])
	if err != nil {
		fail("call failed:", err)
	}
}

func pcDataDelete(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	err := pcClient.DataDelete(args[0], args[1])
	if err != nil {
		fail("call failed:", err)
	}
}

func pcDataUpdate(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	err := pcClient.DataUpdate(args[0], args[1], args[2])
	if err != nil {
		fail("call failed:", err)
	}
}

func pcDataListCmd(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	dataList, err := pcClient.DataList(args)
	if err != nil {
		fail("call failed:", err)
	}

	parsed, err := pcClient.ParseDataListResponse(dataList)
	if err != nil {
		fail("failed to parse response:", err)
	}

	fmt.Println("   |      Table       |     Key      |          Value           |        Created      |      Modified")
	for i, data := range parsed {
		fmt.Printf("%2d | %-16s | %-12s | %-24s | %-16s | %-16s\n",
			i, byteCut(data.Table, 16), byteCut(data.Key, 12), byteCut(data.Value, 24),
			time.Unix(data.CreatedAt, 0).Format("2006-01-02 15:04:05"),
			time.Unix(data.CreatedAt, 0).Format("2006-01-02 15:04:05"))
	}
}

func byteCut(buff []byte, length int) string {
	output := make([]byte, length)
	n := copy(output, buff)
	if len(buff) > length {
		output[length-1] = '.'
		output[length-2] = '.'
		output[length-3] = '.'
	}
	return string(output[:n])
}

func pcTableList(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	tableList, err := pcClient.TableList()
	if err != nil {
		fail("call failed:", err)
	}

	parsed, err := pcClient.ParseTableListResponse(tableList)
	if err != nil {
		fail("failed to parse response:", err)
	}

	for i := range parsed.Tables {
		fmt.Println(string(parsed.Tables[i]))
	}
}

func pcNodeOrder(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	manager := pc.NewManager(pcClient)

	if err := manager.NodeOrder(args[0]); err != nil {
		fail(err)
	}
}

func pcPing(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	if err := pcClient.Ping(); err != nil {
		fail(err)
	}
}

func pcCapabilities(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	capabilities, err := pcClient.Capabilities()
	if err != nil {
		fail(err)
	}

	fmt.Println("Verification methods :", capabilities.VerificationMethods)
	fmt.Println("Version              :", capabilities.Version)
}

func userInfo(cmd *cobra.Command, args []string) {
	pcClient := createClient(cmd, true)

	userInfo, err := pcClient.UserInfo()
	if err != nil {
		fail(err)
	}

	createdAt := time.Unix(userInfo.CreatedAt, 0)

	fmt.Println("ID                  :", userInfo.ID)
	fmt.Println("Created at          :", createdAt.Format(createdAt.Format(time.RFC822)))
	fmt.Println("Verification method :", userInfo.VerificationMethod)
	fmt.Println("Exchange public key :", hex.EncodeToString(userInfo.ExchangePublicKey[:]))
}
