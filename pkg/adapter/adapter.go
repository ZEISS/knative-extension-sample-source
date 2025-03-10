// Package adapter implements a sample receive adapter that generates events
// at a regular interval.
package adapter

import (
	"context"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"

	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"
)

type envConfig struct {
	// Include the standard adapter.EnvConfig used by all adapters.
	adapter.EnvConfig

	// Interval between events, for example "5s", "100ms"
	Interval time.Duration `envconfig:"INTERVAL" required:"true"`
}

func NewEnvConfig() adapter.EnvConfigAccessor { return &envConfig{} }

// Adapter generates events at a regular interval.
type Adapter struct {
	client   cloudevents.Client
	interval time.Duration
	logger   *zap.SugaredLogger

	nextID int
}

type dataExample struct {
	Sequence  int    `json:"sequence"`
	Heartbeat string `json:"heartbeat"`
}

func (a *Adapter) newEvent() cloudevents.Event {
	event := cloudevents.NewEvent()
	event.SetType("com.zeiss.eventing.sources.sample")
	event.SetSource("sample.knative.dev/heartbeat-source")

	if err := event.SetData(cloudevents.ApplicationJSON, &dataExample{
		Sequence:  a.nextID,
		Heartbeat: a.interval.String(),
	}); err != nil {
		a.logger.Errorw("failed to set data")
	}
	a.nextID++
	return event
}

// Start runs the adapter.
// Returns if ctx is cancelled or Send() returns an error.
func (a *Adapter) Start(ctx context.Context) error {
	a.logger.Infow("Starting heartbeat", zap.String("interval", a.interval.String()))
	for {
		select {
		case <-time.After(a.interval):
			event := a.newEvent()

			a.logger.Infow("Sending new event", zap.String("event", event.String()))
			if result := a.client.Send(ctx, event); !cloudevents.IsACK(result) {
				a.logger.Infow("failed to send event", zap.String("event", event.String()), zap.Error(result))
				// We got an error but it could be transient, try again next interval.
				continue
			}
		case <-ctx.Done():
			a.logger.Info("Shutting down...")
			return nil
		}
	}
}

func NewAdapter(ctx context.Context, aEnv adapter.EnvConfigAccessor, ceClient cloudevents.Client) adapter.Adapter {
	env := aEnv.(*envConfig) // Will always be our own envConfig type
	logger := logging.FromContext(ctx)
	logger.Infow("Heartbeat example", zap.Duration("interval", env.Interval))

	return &Adapter{
		interval: env.Interval,
		client:   ceClient,
		logger:   logger,
	}
}
