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
