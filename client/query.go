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

const queryFileList = `
query {
  fileList {
    id
    name
    size
    mime
    key
  }
}
`

const queryFile = `
query($fileID: String!) {
  file(id: $fileID){
    id
    name
    size
    mime
    key
  }
}
`

const mutationFileTransfer = `
mutation($fileID: String!, $connectionKey: String!) {
  fileTransfer(id: $fileID, connection: $connectionKey) {
    id
    name
    size
    mime
    key
  }
}
`

const mutationFileRename = `
mutation($fileID: String!, $fileName: String!) {
  fileRename(id: $fileID, name: $fileName) {
    id
    name
    size
    mime
    key
  }
}
`

const mutationFileRemove = `
mutation($fileID: String!) {
  fileRemove(id: $fileID) {
    id
    name
    size
    mime
    key
  }
}
`
