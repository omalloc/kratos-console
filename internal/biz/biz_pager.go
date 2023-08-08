package biz

type QueryPager struct {
	PageSize int32 `json:"page_size"`
	Current  int32 `json:"current"`
}

type QueryKVFilter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
