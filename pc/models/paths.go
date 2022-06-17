package models

const (
	/*
	   common handlers
	*/
	PathRoot    = "/"
	PathMetrics = "/metrics"
	/*
	   auth handlers
	*/
	PathRegister     = "/auth/register"
	PathLogin        = "/auth/login"
	PathSecureRandom = "/auth/secure-random"
	/*
	   user handlers
	*/
	PathUserRemove = "/user/remove"
	/*
	   data paths
	*/
	PathDataAdd    = "/data/add"
	PathDataUpdate = "/data/update"
	PathDataGet    = "/data/get"
	PathDataDelete = "/data/delete"
	PathDataList   = "/data/list"
	PathTableList  = "/table/list"
)
