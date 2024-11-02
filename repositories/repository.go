package repositories

type Repostories interface {
	Add(interface{}) error
	Update(interface{}) error
	Delete(interface{}) error
	GetByID(interface{}) (interface{}, error)
	GetAll() ([]interface{}, error)
	Search(interface{}) ([]interface{}, error)
}
