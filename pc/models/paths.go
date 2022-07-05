package models

const (
	/*
	   common handlers
	*/
	PathRoot         = "/"
	PathMetrics      = "/metrics"
	PathPing         = "/ping"
	PathCapabilities = "/capabilities"
	/*
	   auth handlers
	*/
	PathRegister     = "/auth/register"
	PathLogin        = "/auth/login"
	PathSecureRandom = "/auth/secure-random"
	/*
	   user handlers
	*/
	PathUserRemove      = "/user/remove"
	PathUserInfo        = "/user/info"
	PathUserExchangeKey = "/user/exchange-key"
	/*
	   data paths
	*/
	PathDataAdd    = "/data/add"
	PathDataUpdate = "/data/update"
	PathDataGet    = "/data/get"
	PathDataDelete = "/data/delete"
	PathDataList   = "/data/list"
	PathTableList  = "/table/list"
	/*
	   enclave paths
	*/

	PathEnclaveRoot   = "/"
	PathEnclaveAdd    = "/entry/add"
	PathEnclaveGet    = "/entry/get"
	PathEnclaveDelete = "/entry/delete"

	/*
		Request paths
	*/
	PathRequestSend         = "/request/send"
	PathRequestIncomingGet  = "/request/incoming/get"
	PathRequestIncomingList = "/request/incoming/list"
	PathRequestOutgoingList = "/request/outgoing/list"

	PathShareSend     = "/share/send"
	PathShareAccept   = "/share/accept"
	PathDataSharedGet = "/data/shared/get"
	PathDataSharedPut = "/data/shared/put"
)
