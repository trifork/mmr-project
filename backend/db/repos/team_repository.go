package repos

import (
	"example.com/m/v2/db/models"
	"gorm.io/gorm"
)

type ITeamRepository interface {
    CreateTeam(team *models.Team) error
    ListTeams() ([]*models.Team, error)
}

type TeamRepository struct {
    db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) ITeamRepository {
    return &TeamRepository{db}
}

func (tr *TeamRepository) CreateTeam(team *models.Team) error {
    return tr.db.Create(team).Error
}

func (tr *TeamRepository) ListTeams() ([]*models.Team, error) {
    var teams []*models.Team
    if err := tr.db.Find(&teams).Error; err != nil {
        return nil, err
    }
    return teams, nil
}
