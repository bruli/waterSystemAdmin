package programs

import "context"

type RemoveTemperature struct {
	repo TemperatureRepository
}

func (r RemoveTemperature) Remove(ctx context.Context, temperature int) error {
	return r.repo.Remove(ctx, temperature)
}

func NewRemoveTemperature(repo TemperatureRepository) *RemoveTemperature {
	return &RemoveTemperature{repo: repo}
}
