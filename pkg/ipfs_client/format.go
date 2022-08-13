package ipfs_client

func BoolString(v bool) string {
	if v {
		return "true"
	}

	return "false"
}
