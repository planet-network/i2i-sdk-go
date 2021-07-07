package client

const queryInfo = `
query {
  info {
    version
    network {
      running
    }
  }
}
`

const mutationAclAdd = `
mutation($input:AclInput!){
	aclAdd(input:{data: $input}){
		authorization
		device_token
		notification_provider
	}
}
`

const queryAclList = `
query {
  aclList {
    name
    UUID
    permissions
    authorization
    device_token
    notification_provider
    private_pl_scope
  }
}
`

const mutationInitialize = `
mutation ($type: I2IType!) {
    initialize(input:{type: $type})
  }
`
