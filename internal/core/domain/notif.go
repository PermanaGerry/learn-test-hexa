package domain

type PushNotif struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Type    int    `json:"type"`
}

func NewNotif(id string, message string, tipe int) PushNotif {
	return PushNotif{
		ID:      id,
		Message: message,
		Type:    tipe,
	}
}
