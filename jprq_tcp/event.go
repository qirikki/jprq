package jprq_tcp

type TunnelStartedEvent struct {
	PublicServerPort  int `json:"public_server_port"`
	PrivateServerPort int `json:"private_server_port"`
}

type ConnectionReceivedEvent struct {
	PublicClientPort int `json:"public_client_port"`
}

type ConnectionPairedEvent struct {
	PublicClientPort  int `json:"public_client_port"`
	PrivateClientPort int `json:"private_client_port"`
}
