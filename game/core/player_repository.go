package core

type PlayerRepository interface {
	GetAll([]Player) error
	GetById(id string) (*Player, error)
	Create(player Player) error
	Update(id string) (*Player, error)
	Delete(id string) (*Player, error)
}
