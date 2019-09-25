package mapper

type CustomerMapper struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}