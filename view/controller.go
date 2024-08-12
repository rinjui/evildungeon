package view

type Controller interface {
	Redirect(id ViewID) error
}
