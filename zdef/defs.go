package zdef

type SessionConf struct {
	ServiceName string `json:"ServiceName,omitempty"` // game gate ...
	ID          int64  `json:"ID,omitempty"`
	Host        string `json:"Host,omitempty"`
}
