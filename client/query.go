package client

const queryInfo = `
query {
  info {
    version
    network{
      running
      ip
      network_port
      api_port
      vpn_port
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

const mutationDMeUpdate = `
mutation($input:DMeInput!){
	dMeUpdate(input: $input){
		first_name
	}
}
`

const mutationAclRemove = `
mutation ($id: ID!) {
    aclRemove(id: $id)
  }
`

const queryAclList = `
query {
  aclList {
    id
    application_name
    application_version
    os_version
    location {
      latitude
      longitude
    }
    uuid
    device_token
    notification_provider
    device_name
    private_pl_scope_name
    permissions
    created_at
    last_usage
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

const queryPlDataRead = `
query($input: PlDataReadInput!) {
  plDataRead(input: $input)
}
`

const mutationAddConnection = `
mutation ($input: ConnectionInput!){
	connectionAdd(input:$input)
}
`

const interactiveActionsFriendRequestQry = `
query {
	interactiveActions {
		... on FriendRequest {
			id
			source
			time
			full_name
		}
	}
}
`

const interactiveActionsQry = `
query {
	interactiveActions {
		id
		source
		time
	}
}
`

const mutationInteractiveAction = `
mutation($input: NotificationAction!) {
  interactiveActionUpdate(action: $input)
}
`

const mutationDMeProfileAdd = `
mutation($input:DMeProfileInput!){
	dMeProfileAdd(input: $input) {
		avatar_url
    	profile_name
	}
}
`

const queryConnectionList = `
query($profile: ID!) {
  connectionList(profile: $profile) {
    avatar_url
    id
    profile
    public_key
    signature_key
    display_name
    name
    surname
    transactions
  }
}
`

const queryProfileList = `
query {
  profileList {
    avatar_url
    profile_name
  }
}
`

const mutationDMeProfileUpdate = `
mutation($input:DMeProfileInput!){
	dMeProfileUpdate(input: $input) {
		avatar_url
    	profile_name
	}
}
`

const mutationVnfWireguardCreate = `
mutation($input:WireguardConfigInput!) {
  vnfWireguardCreate(input: $input) {
    running
    name
    address
    listen_port
    private_key
    public_key
    post_up
    post_down
    dns
  }
}
`

const mutationVnfWireguardStop = `
mutation($input:String!) {
  vnfWireguardStop(name: $input) {
    running
    name
    address
    listen_port
    private_key
    public_key
    post_up
    post_down
    dns
  }
}
`

const mutationVnfWireguardStart = `
mutation($input:String!) {
  vnfWireguardStart(name: $input) {
    running
    name
    address
    listen_port
    private_key
    public_key
    post_up
    post_down
    dns
  }
}
`

const mutationVnfWireguardCreatePeerConfig = `
mutation($input: WireguardPeerInput!) {
  vnfWireguardCreatePeerConfig(input: $input ) {
    name
    address
    endpoint
    private_key
    peer_public_key
    dns
    allowed_ips
  }
}
`

const queryDMeInfo = `
query {
  dMeInfo {
    first_name
    surname
    public_key
    signature_key
  }
}
`

const mutationReset = `
mutation {
  reset
}
`

const queryPlScopeList = `
query {
  plScopeList
}
`

const queryPlVerify = `
query {
  plVerify {
    scope
    object
    name
    message
  }
}
`

const queryPlInstances = `
query($filter: InstanceFilterInput ) {
  plInstances(filter: $filter) {
    ID
    as
    nbuckets
    characteristics {
      name
      value
    }
  }
}
`

const queryPlRelations = `
query($filter: RelationFilterInput) {
  plRelations(filter: $filter) {
    ID
    as
    characteristics {
      name
      value
    }
    relatives {
      name
      ID
    }
    nbuckets
  }
}
`
