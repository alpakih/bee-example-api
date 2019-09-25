package response

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data, omitempty"`
}
