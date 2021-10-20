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
	AvatarFileID  *string `json:"avatar_file_id"`
	ProfileName   string  `json:"profile_name"`
	AvatarURL     *string `json:"avatar_url"`
	HideFirstName *bool   `json:"hide_first_name"`
	HideSurname   *bool   `json:"hide_surname"`
	Pseudonym     *string `json:"pseudonym"`
	Bio           *string `json:"bio"`
}

type WireguardConfig struct {
	Running    bool   `json:"running"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	ListenPort int    `json:"listen_port"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	PostUp     string `json:"post_up"`
	PostDown   string `json:"post_down"`
	DNS        string `json:"dns"`
}

type WireguardConfigInput struct {
	Name       string  `json:"name"`
	PrivateKey *string `json:"privateKey"`
	Address    *string `json:"address"`
}

type WireguardPeerConfig struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	Endpoint      string `json:"endpoint"`
	PrivateKey    string `json:"private_key"`
	PeerPublicKey string `json:"peer_public_key"`
	DNS           string `json:"dns"`
	AllowedIps    string `json:"allowed_ips"`
}

type WireguardPeerInput struct {
	NetworkName string  `json:"network_name"`
	PeerName    string  `json:"peer_name"`
	PrivateKey  *string `json:"privateKey"`
	Address     *string `json:"address"`
}

type ConnectionInput struct {
	//  ID is the database ID of the connection
	ID string `json:"ID"`
	// profile contains id of the profiles to which connection is related
	Profile []string `json:"profile"`
	// public_key if set will change public key of the connection
	PublicKey string `json:"public_key"`
	// signature_key if set will change signature key of the connection
	SignatureKey string `json:"signature_key,omitempty"`
	// display_name if set will change display name of the connection
	DisplayName string `json:"display_name"`
	// name if set will change name of the connection
	Name string `json:"name"`
	// surname if set will change surname of the connection
	Surname string `json:"surname"`
	// country if set will change country of the connection
	Country string `json:"country"`
	//  for internal usage
	Transactions string `json:"transactions"`
}

type FriendRequest struct {
	// id is the database id of the object
	ID string `json:"id"`
	// id is the database id of the object
	Source string `json:"source"`
	// full_name is complete set of entity name making request
	FullName string `json:"full_name"`
	// time is
	Time time.Time `json:"time"`
}

type InteractiveAction struct {
	// id is the database id of the object
	ID string `json:"id"`
	// id is the database id of the object
	Source string `json:"source"`
	// time is
	Time time.Time `json:"time"`
}

// NotificationAction is action done on notification which require manual input from owner
type NotificationAction struct {
	ID     string `json:"id"`
	Action string `json:"action"`
}
