package ports

type BroadCastHandler interface {
	BroadCast(data []map[string]interface{})
}
