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
	// decrypted_size is files size in bytes before encryption
	DecryptedSize int `json:"decrypted_size"`
	// encrypted_size is files size in bytes after encryption
	EncryptedSize int `json:"encrypted_size"`
	// mime contains MIME information
	Mime string `json:"mime"`
	Key  string `json:"key"`
	// location of file on the i2i node
	Path string `json:"path"`
}

// FileRest is designed for for /fu endpoint
type FileRest struct {
	// id is the database id of the object
	ID string `json:"id"`
	// name contains original file name
	Name string `json:"name"`
	// decrypted_size is files size in bytes before encryption
	DecryptedSize int `json:"decrypted_size"`
	// encrypted_size is files size in bytes after encryption
	EncryptedSize int `json:"encrypted_size"`
	// mime contains MIME information
	Mime string      `json:"mime"`
	Key  interface{} `json:"key"`
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

type DMeInfo struct {
	FirstName    string `json:"first_name"`
	Surname      string `json:"surname"`
	PublicKey    string `json:"public_key"`
	SignatureKey string `json:"signature_key"`
}

type Relation struct {
	ID              string            `json:"ID"`
	As              string            `json:"as"`
	Characteristics []*Characteristic `json:"characteristics"`
	Relatives       []*Relative       `json:"relatives"`
	Nbuckets        []string          `json:"nbuckets"`
}

type Instance struct {
	ID              string            `json:"ID"`
	As              string            `json:"as"`
	Nbuckets        []string          `json:"nbuckets"`
	Characteristics []*Characteristic `json:"characteristics"`
}

type InstanceFilterInput struct {
	Scope *string `json:"scope"`
	As    *string `json:"as"`
}

type Characteristic struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

type Relative struct {
	Name string `json:"name"`
	ID   string `json:"ID"`
}

type RelationFilterInput struct {
	Scope *string `json:"scope"`
	As    *string `json:"as"`
}

type PlReport struct {
	Scope   string `json:"scope"`
	Object  string `json:"object"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type InstanceInput struct {
	Scope           string                 `json:"scope"`
	As              *string                `json:"as"`
	ID              *string                `json:"id"`
	Characteristics []*CharacteristicInput `json:"characteristics"`
}

type CharacteristicInput struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

type DirectMessageInput struct {
	ID *string `json:"id"`
	// destination is recipent public key
	Destination string `json:"destination"`
	// content is content of message to be delivered to recipent
	Content string `json:"content"`
	//  reply if set is reply to message with ID
	Reply       *string            `json:"reply"`
	Attachments []*AttachmentInput `json:"attachments"`
}

type AttachmentInput struct {
	//  attachment_type is type of upload can be: video, image, gif, file
	AttachmentType string `json:"attachment_type"`
	// if uploaded from local device upload goes here
	//Upload *graphql.Upload `json:"upload"`
	//  if used from external source, url is provided
	URL *string `json:"url"`
	//  raw payload is used for case of some eg. JSON formatted data
	RawPayload *string `json:"raw_payload"`
}

type DirectMessage struct {
	// id is the database id of the object
	ID string `json:"id"`
	// source is public key of sender
	Source string `json:"source"`
	// destination is the public key of recipient
	Destination string `json:"destination"`
	// content is read content of the message
	Content string `json:"content"`
	// read is boolean value defining whether message was read or not
	Read bool `json:"read"`
	// time is time when message was send or received
	Time time.Time `json:"time"`
	// star is custom attribute set by user to mark message
	Star bool `json:"star"`
	// reply if set is ID of message it's reply to
	Reply *string `json:"reply"`
	//  incoming if set to true would mean that last message is incoming, otherwise outcoming
	Incoming bool `json:"incoming"`
}

type DirectMessagePage struct {
	TotalCount  int              `json:"totalCount"`
	HasNextPage bool             `json:"has_next_page"`
	Messages    []*DirectMessage `json:"messages"`
}

type MessageViewInput struct {
	Conversation string  `json:"conversation"`
	After        *string `json:"after"`
	Before       *string `json:"before"`
	Count        int     `json:"count"`
}

type Conversation struct {
	// id is the database id of the message, in case of friend request it's transaction ID
	ID string `json:"id"`
	// source is public key of sender
	Source string `json:"source"`
	// destination is the public key of recipient
	Destination string `json:"destination"`
	// time is time when message was send or received
	Time time.Time `json:"time"`
	// content is read content of the message
	Content string `json:"content"`
	// url of the file containing avatar of
	AvatarURL string `json:"avatar_url"`
	// message_type is type of
	MessageType string `json:"message_type"`
	// display_name is string value under which contact is displayed in contact list
	DisplayName string `json:"display_name"`
	//  unread_count is number of unread messages in conversation
	UnreadCount int `json:"unread_count"`
	//  incoming if set to true would mean that last message is incoming, otherwise outcoming
	Incoming bool `json:"incoming"`
	//  action_id defines action
	ActionID string `json:"action_id"`
	//  conversation_id is id related to message conversation
	ConversationID   string `json:"conversation_id"`
	GroupDisplayName string `json:"group_display_name"`
	GroupAvatarURL   string `json:"group_avatar_url"`
}

// GroupChat defines group chat
type GroupChat struct {
	// id is the database id of the object
	ID string `json:"id"`
	// public_key is identifier of the group chat
	PublicKey string `json:"public_key"`
	// private_key is per chat private key
	PrivateKey string `json:"private_key"`
	// name is human readable name of the group chat
	GroupDisplayName string `json:"group_display_name"`
	// participants is list of prublic key who participate in group chat
	Participants []*ChatParticipant `json:"participants"`
	// admin is public key of group chat admin
	Admin *ChatParticipant `json:"admin"`
	//  avatar of the group chat
	GroupAvatarURL string `json:"group_avatar_url"`
	// created_at is time when group chat was created
	CreatedAt time.Time `json:"created_at"`
	// joined_at is time when user joined group chat
	JoinedAt time.Time `json:"joined_at"`
	// left_at is time when user left group chat
	LeftAt *time.Time `json:"left_at"`
	// left defines if user has left group chat
	Left bool `json:"left"`
}

type ChatParticipant struct {
	AvatarURL    string `json:"avatar_url"`
	PublicKey    string `json:"public_key"`
	SignatureKey string `json:"signature_key"`
	DisplayName  string `json:"display_name"`
}

// GroupchatAddUser is used as argument for mutations adding user to group chat.
// The caller responsibility is to be sure that he has sufficient permissions to do so.
// Otherwise error is returned
type GroupchatAddUser struct {
	// chat_public_key is public key of the chat
	ChatPublicKey string `json:"chat_public_key"`
	// user_signature_key is signature key of user which is added to group chat
	UserSignatureKey []*string `json:"user_signature_key"`
	// user_signature_key is network key of user which is added to group chat, it will be automatically set by i2i
	UserNetworkKey []*string `json:"user_network_key"`
}

type GroupchatInput struct {
	Profile          *string   `json:"profile"`
	GroupDisplayName *string   `json:"group_display_name"`
	ChatPublicKey    *string   `json:"chat_public_key"`
	ChatPrivateKey   *string   `json:"chat_private_key"`
	GroupAvatarURL   *string   `json:"group_avatar_url"`
	Participants     []*string `json:"participants"`
}
