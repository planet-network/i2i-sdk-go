package client

type NetworkInfo struct {
	Running bool `json:"running"`
}

// Info contains i2i related information
type Info struct {
	// i2i version
	Version string      `json:"version"`
	Network NetworkInfo `json:"network"`
}

type ACL struct {
	Name                 *string  `json:"name"`
	UUID                 string   `json:"UUID"`
	Permissions          []string `json:"permissions"`
	Authorization        string   `json:"authorization"`
	DeviceToken          *string  `json:"device_token"`
	NotificationProvider *string  `json:"notification_provider"`
	PrivatePlScope       *bool    `json:"private_pl_scope"`
}

type ACLInput struct {
	DeviceToken          *string  `json:"device_token"`
	NotificationProvider *string  `json:"notification_provider"`
	UUID                 string   `json:"uuid"`
	Name                 *string  `json:"name"`
	PrivatePlScope       *bool    `json:"private_pl_scope"`
	Permissions          []string `json:"permissions"`
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
