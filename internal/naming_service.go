package internal

type NamingService struct {
}

func NewNamingService() *NamingService {
	return &NamingService{}
}

func (n NamingService) GetNewFilename(uploadedFilename string) string {
	//TODO implement me
	panic("implement me")
}

func (n NamingService) GetFullFSPath(filename string) string {
	//TODO implement me
	panic("implement me")
}

func (n NamingService) GetURL(filename string) string {
	//TODO implement me
	panic("implement me")
}
