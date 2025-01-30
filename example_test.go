package fxconfig_test

import (
	"fmt"
	"time"

	"schneider.vip/config"
	"schneider.vip/fxconfig"

	"go.uber.org/fx"
)

// ConfigSection represents a configuration section.
type ConfigSection struct {
	URL  string
	True bool `mapstructure:",omitempty"`
}

// NewService is a constructor that uses the dynamic configuration.
func NewService(loader config.Dynamic[ConfigSection]) {
	for {
		cfg := loader.Load()
		fmt.Printf("Service Config: URL=%s, True=%v\n", cfg.URL, cfg.True)
		time.Sleep(1 * time.Second)
	}
}

// Example_fxconfig demonstrates how to use fxconfig with fx to load and inject configuration sections.
func Example_fxconfig() {
	// Create a new fx App with fxconfig.
	app := fx.New(
		fx.Provide(
			// Provide a dynamic configuration loader for the ConfigSection.
			fxconfig.New(config.WithSubSection[ConfigSection]("ServiceConfig")),
		),
		fx.Invoke(
			// Invoke the NewService function with the loaded configuration.
			NewService,
		),
	)

	// Run the application.
	app.Run()

	// Output:
	// Service Config: URL=example.com, True=true
}
