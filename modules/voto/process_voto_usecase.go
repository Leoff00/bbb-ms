package voto

import (
	"os"

	"github.com/rs/zerolog"
)

type VotoUseCase struct {
	vp *VotoProducer
}

func NewVotoUseCase(vp *VotoProducer) *VotoUseCase {
	return &VotoUseCase{
		vp: vp,
	}
}

func (vc *VotoUseCase) processVote(voto int) error {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	err := vc.vp.VotoProducer(voto)

	if err != nil {
		logger.Err(err).Msg("Failed to produce the data in queue")
		return err
	}

	return nil
}
