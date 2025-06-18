package repository

type NotificationTemplate struct {
	ID          int64
	Name        string
	Description string
	ChannelType string
	Template    string
}
