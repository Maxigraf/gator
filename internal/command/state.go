package command

import (
	"github.com/maxigraf/gator/internal/config"
	"github.com/maxigraf/gator/internal/database"
)

type State struct {
	Database *database.Queries
	Config   *config.Config
}
