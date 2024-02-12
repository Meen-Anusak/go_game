package core

type PlayerService interface {
	CreateNewPlayer(player Player) error
	GetPlayer(player Player) error
	GetPlayerById(player Player) error
	UpdatePlayer(player Player) error
	DeletePlayer(player Player) error
}

type playerServiceImpl struct {
	repo PlayerRepository
}

// DeletePlayer implements PlayerService.
func (*playerServiceImpl) DeletePlayer(player Player) error {
	panic("unimplemented")
}

// GetPlayer implements PlayerService.
func (*playerServiceImpl) GetPlayer(player Player) error {
	panic("unimplemented")
}

// GetPlayerById implements PlayerService.
func (*playerServiceImpl) GetPlayerById(player Player) error {
	panic("unimplemented")
}

// UpdatePlayer implements PlayerService.
func (*playerServiceImpl) UpdatePlayer(player Player) error {
	panic("unimplemented")
}

func NewPlayerService(repo PlayerRepository) PlayerService {
	return &playerServiceImpl{repo: repo}
}

func (s *playerServiceImpl) CreateNewPlayer(player Player) error {
	if err := s.repo.Create(player); err != nil {
		return err
	}
	return nil
}
