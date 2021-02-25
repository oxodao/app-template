package services

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/oxodao/app-template/config"
	"github.com/oxodao/app-template/dal"
)

type Provider struct {
	Config *config.Config
	DB     *sqlx.DB
	Dal    *dal.Dal
}

func NewProvider(cfg *config.Config) (*Provider, error) {
	db, err := sqlx.Open(config.DbProvider, cfg.Database.GetDSN())
	if err != nil {
		return nil, err
	}

	return &Provider{
		Config: cfg,
		DB:     db,
		Dal:    dal.NewDal(db),
	}, nil
}

func (prv *Provider) Shutdown() {
	prv.DB.Close()
}
