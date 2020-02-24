package devices

//Repository interface
type Repository interface {
	Save(d *Device) error
}
