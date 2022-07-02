package services

import (
	"sgcu65/models"
	"sgcu65/repository"
)

type teamService struct {
	teamRepository repository.TeamRepository
}

type TeamService interface {
	AddTeam(team models.Team) (models.Team, error)
	GetTeam(teamId string) (models.Team, error)
	DeleteTeam(teamId string) error
}

func NewTeamService(r repository.TeamRepository) TeamService {
	return teamService{
		teamRepository: r,
	}
}

func (t teamService) AddTeam(team models.Team) (models.Team, error) {
	return t.teamRepository.AddTeam(team)
}
func (t teamService) GetTeam(teamId string) (models.Team, error) {
	return t.teamRepository.GetTeam(teamId)
}
func (t teamService) DeleteTeam(teamId string) error {
	return t.teamRepository.DeleteTeam(teamId)
}
