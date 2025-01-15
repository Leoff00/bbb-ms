package voto

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type VotoRepository struct {
	Db *sql.DB
}

func NewVotoRepository(db *sql.DB) *VotoRepository {
	return &VotoRepository{
		Db: db,
	}
}

func (vr *VotoRepository) Save(voto int) error {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := vr.Db.PrepareContext(ctx, "INSERT INTO voto (voto) VALUES ($1) RETURNING id;")
	if err != nil {
		logger.Err(err).Msg("Failed to prepare query")
		return err
	}

	_, err = stmt.ExecContext(ctx, voto)

	if err != nil {
		logger.Err(err).Msg("Failed to exec query")
		return err
	}

	logger.Info().Msg("Data persisted successfully")
	return nil
}
