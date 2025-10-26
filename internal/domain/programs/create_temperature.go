package programs

import "context"

type CreateTemperature struct {
	repo TemperatureRepository
}

func (c CreateTemperature) Create(ctx context.Context, p *TemperatureProgram) error {
	return c.repo.Save(ctx, p)
}

func NewCreateTemperature(repo TemperatureRepository) *CreateTemperature {
	return &CreateTemperature{repo: repo}
}
