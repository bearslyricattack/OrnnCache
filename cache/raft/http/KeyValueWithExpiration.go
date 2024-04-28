package http

type KeyValueWithExpiration struct {
	Value   interface{} `json:"value,omitempty"`
	Expired int64       `json:"expired,omitempty"`
}
