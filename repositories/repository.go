package repositories

type Repostories interface {
	Add(interface{}) error
	Update(interface{}) error
	Delete(interface{}) error
	GetByID(interface{}) (interface{}, error)
	GetAll(limit, offset int) ([]interface{}, error)
	GetAllWithFilter(interface{}) ([]interface{}, error)
}
