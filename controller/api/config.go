package api

var Api ApiInfo

type ApiInfo struct {
	Key string `json:"key"`
}

func Configure(a ApiInfo) {
	Api = a
}
