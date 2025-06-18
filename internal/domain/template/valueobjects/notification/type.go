package notification

type Type interface {
	~string
	Validate() error
}
