package core

type PlayerRepository interface {
	Get(player Player) error
	GetById(player Player) error
	Create(player Player) error
	Update(player Player) error
	Delete(player Player) error
}
