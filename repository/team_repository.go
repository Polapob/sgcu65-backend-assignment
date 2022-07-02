package repository

import (
	"sgcu65/models"

	"gorm.io/gorm"
)

type teamRepository struct {
	DB *gorm.DB
}

type TeamRepository interface {
	AddTeam(team models.Team) (models.Team, error)
	GetTeam(teamId string) (models.Team, error)
	DeleteTeam(teamId string) error
	Migrate() error
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return teamRepository{
		DB: db,
	}
}

func (t teamRepository) Migrate() error {
	return t.DB.AutoMigrate(&models.Team{})
}

func (t teamRepository) AddTeam(team models.Team) (models.Team, error) {
	err := t.DB.Create(&team).Error
	return team, err
}

func (t teamRepository) GetTeam(teamId string) (models.Team, error) {
	var team models.Team
	err := t.DB.Where("id = ?", teamId).First(&team).Error
	return team, err
}

func (t teamRepository) DeleteTeam(teamId string) error {
	err := t.DB.Delete(&models.Team{}, teamId).Error
	return err
}
