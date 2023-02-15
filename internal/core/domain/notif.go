package domain

type PushNotif struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Type    int    `json:"type"`
}

type ShowNotif struct {
	Lists []PushNotif
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewNotif(id string, message string, tipe int) PushNotif {
	return PushNotif{
		ID:      id,
		Message: message,
		Type:    tipe,
	}
}
