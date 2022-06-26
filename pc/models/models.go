package models

type LoginResponse struct {
	Authorization string   `json:"authorization"`
	SecureRandom  [32]byte `json:"secure_random"`
}

type LoginRequest struct {
	Login            string   `json:"login"`
	AuthorizationKey [32]byte `json:"secret"`
}

type SecureRandomRequest struct {
	AuthorizationKey [32]byte `json:"secret"`
}

type SecureRandomResponse struct {
	SecureRandom [32]byte `json:"secure_random"`
}

type RegisterRequest struct {
	// Login is raw login, needs to be done that way, because it is used for verification
	Login            string   `json:"login"`
	AuthorizationKey [32]byte `json:"authorization_key"`
	// name of verification method which will be used for verifying user login
	VerificationMethod          string   `json:"verification_method"`
	SecureRandom                [32]byte `json:"secure_random"`
	ExchangePublicKey           [32]byte `json:"share_public_key"`
	EncryptedExchangePrivateKey []byte   `json:"encrypted_exchange_private_key"`
}

type DataGetRequest struct {
	// Table is name of the table in which the data is stored
	Table []byte `json:"table"`
	// Key is data unique key per table data is stored in
	Key []byte `json:"key"`
}

type DataListRequest struct {
	// Tables is name of the tables from which tables should be listed, if nil, all data is returned
	Tables [][]byte `json:"table"`
}

type DataAddRequest struct {
	// Table is name of the table in which the data is stored
	Table []byte `json:"table"`
	// Key is data unique key per table data is stored in
	Key []byte `json:"key"`
	// Value is data value
	Value []byte `json:"value"`
	// EncryptedValueKey is key used for encrypting value. The key is encrypted with
	EncryptedValueKey []byte `json:"encrypted_value_key"`
}

type DataUpdateRequest struct {
	// Table is name of the table in which the data is stored
	Table []byte `json:"table"`
	// Key is data unique key per table data is stored in
	Key []byte `json:"key"`
	// Value is data value
	Value []byte `json:"value"`
}

type DataDeleteRequest struct {
	// Table is name of the table in which the data is stored
	Table []byte `json:"table"`
	// Key is data unique key per table data is stored in
	Key []byte `json:"key"`
}

type DataResponse struct {
	// Table is name of the table in which the data is stored
	Table []byte `json:"table"`
	// Key is data unique key per table data is stored in
	Key []byte `json:"key"`
	// Value is data value
	Value []byte `json:"value"`
	// CreatedAt is epoch time when data was added
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is epoch time when data was last time updated
	ModifiedAt int64 `json:"modified_at"`
	// EncryptedValueKey is key used for encrypting value. The key is encrypted with
	EncryptedValueKey []byte `json:"encrypted_value_key"`
}

type TableListResponse struct {
	// Tables is set of names of user tables
	Tables [][]byte `json:"tables"`
}

type UserInfoResponse struct {
	ID                          string   `json:"id"`
	VerificationMethod          string   `json:"verification_method"`
	CreatedAt                   int64    `json:"created_at"`
	ExchangePublicKey           [32]byte `json:"exchange_public_key"`
	EncryptedExchangePrivateKey []byte   `json:"encrypted_exchange_private_key"`
}

type CapabilitiesResponse struct {
	Version             string   `json:"version"`
	VerificationMethods []string `json:"verification_methods"`
}

type DataRequestPack struct {
	Incoming DataRequestSendRequest `json:"incoming"`
	Outgoing DataRequestSendRequest `json:"outgoing"`
}

type DataRequestSendRequest struct {
	// ToUserID is user identifier who is recipient of the request
	ToUserID [32]byte `json:"to_user_id"`
	// Until is epoch seconds determining when data request will expire
	Until int64 `json:"until"`
	// EncryptedFrom is encrypted login with master or exchange key, this is used for storing info about sent requests
	EncryptedFrom []byte `json:"encrypted_from"`
	// EncryptedTo is encrypted login with master key or exchange, this is used for storing info about sent requests
	EncryptedTo []byte `json:"encrypted_to"`
	// EncryptedDescription is encrypted description, this is used for storing info about sent requests
	EncryptedDescription []byte `json:"encrypted_description"`
	// Items are collection of data which are being requested from destination
	Items []*DataRequestItem `json:"items"`
	// Incoming if set to true means that request is incoming one, otherwise outgoing
	Incoming bool `json:"incoming"`
	// Accepted if set to true, means that recipient has agreed to share the data
	Accepted bool `json:"accepted"`
}

type DataRequestItem struct {
	// EncryptedTable is name of the table in which the data is stored
	EncryptedTable []byte `json:"encrypted_table"`
	// EncryptedKey is data unique key per table data is stored in
	EncryptedKey []byte `json:"encrypted_key"`
}

type GetPublicExchangeKeyRequest struct {
	// User is user identifier which public key is requested
	User [32]byte `json:"user"`
}

type GetPublicExchangeKeyResponse struct {
	// ExchangePublicKey is requested user public key
	ExchangePublicKey [32]byte `json:"exchange_public_key"`
}
