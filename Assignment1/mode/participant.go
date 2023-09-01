package mode

type Data struct {
	Id         string `json:"id"`
	Code       string `json:"student_code"`
	Name       string `json:"student_name"`
	Address    string `json:"student_address"`
	Occupation string `json:"student_occupation"`
	Reason     string `json:"joining_reason"`
}

type Participants struct {
	Participants []Data `json:"participants"`
}
