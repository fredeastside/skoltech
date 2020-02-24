package devices

//IService interface
type IService interface {
	CreateDevice() *Device
	Save(d *Device) error
}

//Service layer for managing devices
type Service struct {
	repository Repository
}

//NewService constructor
func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

//CreateDevice return new device structure
func (s *Service) CreateDevice() *Device {
	return NewDevice()
}

//Save device in stirage
func (s *Service) Save(d *Device) error {
	return s.repository.Save(d)
}
