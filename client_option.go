package vmess

type ClientOption func(*Client)

func ClientWithGlobalPadding() ClientOption {
	return func(client *Client) {
		client.globalPadding = true
	}
}

func ClientWithAuthenticatedLength() ClientOption {
	return func(client *Client) {
		client.authenticatedLength = true
	}
}
