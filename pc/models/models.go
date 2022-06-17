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
	VerificationMethod       string   `json:"verification_method"`
	SecureRandom             [32]byte `json:"secure_random"`
	ExchangePublicKey        [32]byte `json:"share_public_key"`
	EncryptedSharePrivateKey []byte   `json:"encrypted_share_private_key"`
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
