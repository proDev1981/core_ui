package core

// struct reciver sms to  client
type reciverSms struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Event Event  `json:"event"`
	State Event  `json:"state"`
}
