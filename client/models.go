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
