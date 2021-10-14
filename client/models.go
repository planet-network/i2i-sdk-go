package client

import "time"

type NetworkInfo struct {
	Running     bool    `json:"running"`
	IP          *string `json:"ip"`
	NetworkPort *int    `json:"network_port"`
	APIPort     *int    `json:"api_port"`
	VpnPort     *int    `json:"vpn_port"`
}

// Info contains i2i related intofmation
type Info struct {
	// i2i version
	Version string       `json:"version"`
	Network *NetworkInfo `json:"network"`
}

type ACLInput struct {
	//  id is id of the token to be updates, the value is not meant to be used for token creation
	ID *string `json:"id"`
	//  name of application which will use the token
	ApplicationName *string `json:"application_name"`
	//  version of the application
	ApplicationVersion *string `json:"application_version"`
	//  name of the operating system
	OsName *string `json:"os_name"`
	//  version of the operating system
	OsVersion *string `json:"os_version"`
	//  current location of the device
	Location *LocationInput `json:"location"`
	//  hardware unique identifier
	UUID *string `json:"uuid"`
	//  device token for notification provider
	DeviceToken *string `json:"device_token"`
	//  notification provider is determine where notification will be sent
	NotificationProvider *string `json:"notification_provider"`
	//  name of the device, eg Apple iPhone 8
	DeviceName *string `json:"device_name"`
	//  name of the private scope to create, preferably name of the app
	PrivatePlScopeName *string  `json:"private_pl_scope_name"`
	Permissions        []string `json:"permissions"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LocationInput struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ACL struct {
	ID string `json:"id"`
	//  name of application which will use the token
	ApplicationName *string `json:"application_name"`
	//  version of the application
	ApplicationVersion *string `json:"application_version"`
	//  name of the operating system
	OsName *string `json:"os_name"`
	//  version of the operating system
	OsVersion *string `json:"os_version"`
	//  current location of the device
	Location *Location `json:"location"`
	//  hardware unique identifier
	UUID *string `json:"uuid"`
	//  device token for notification provider
	DeviceToken          *string `json:"device_token"`
	NotificationProvider *string `json:"notification_provider"`
	//  name of the device, eg Apple iPhone 8
	DeviceName *string `json:"device_name"`
	//  name of the private scope to create, preferably name of the app
	PrivatePlScopeName *string  `json:"private_pl_scope_name"`
	Permissions        []string `json:"permissions"`
	//  value of the token
	Authorization string    `json:"authorization"`
	CreatedAt     time.Time `json:"created_at"`
	LastUsage     time.Time `json:"last_usage"`
}

type File struct {
	// id is the database id of the object
	ID string `json:"id"`
	// name contains original file name
	Name string `json:"name"`
	// size is files size in bytes after encryption
	Size int `json:"size"`
	// mime contains MIME information
	Mime string `json:"mime"`
	Key  string `json:"key"`
	// location of file on the i2i node
	Path string `json:"path"`
}

type PlDataReadInput struct {
	Dst            string `json:"dst"`
	Instance       string `json:"instance"`
	Characteristic string `json:"characteristic"`
}

// Connection represents other entity in the network which interacts
type Connection struct {
	AvatarURL string `json:"avatar_url"`
	// id is the database id of the object
	ID string `json:"id"`
	// profile contains id of the profiles to which connection is related
	Profile      []string `json:"profile"`
	PublicKey    string   `json:"public_key"`
	SignatureKey string   `json:"signature_key"`
	// display_name is string value under which contact is displayed in contact list
	DisplayName string `json:"display_name"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	// transactions contains list of transactions related to given contact
	Transactions []string `json:"transactions"`
}

type Profile struct {
	AvatarURL   string `json:"avatar_url"`
	ProfileName string `json:"profile_name"`
}

type DMeProfileInput struct {
	ProfileName   string  `json:"profile_name"`
	AvatarURL     *string `json:"avatar_url"`
	HideFirstName *bool   `json:"hide_first_name"`
	HideSurname   *bool   `json:"hide_surname"`
	Pseudonym     *string `json:"pseudonym"`
	Bio           *string `json:"bio"`
}
