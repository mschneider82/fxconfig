# fxconfig - A Dependency Injection Wrapper for Dynamic Configuration

`fxconfig` is a lightweight wrapper library designed to simplify the integration of dynamic configuration management into applications using the `go.uber.org/fx` dependency injection framework. It leverages the `schneider.vip/config` package to provide a seamless way to load and manage configuration sections dynamically.

## Features

* *Dynamic Configuration Loading*: Load configuration sections dynamically and watch for changes.

* *Dependency Injection Integration*: Easily inject configuration sections into your application using fx.

* *Customizable*: Supports custom Viper instances, environment variable binding, and more.

## Installation
To use `fxconfig`, you need to have Go installed. Then, you can install the package using:

```bash
go get schneider.vip/fxconfig
```

# Usage
## Basic Example
Here’s a basic example of how to use fxconfig with fx to load and inject configuration sections:

```go
package main

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
		fmt.Println("Service Config:", cfg.URL, cfg.True)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fx.New(
		fx.Provide(
			// Provide a dynamic configuration loader for the ConfigSection.
			fxconfig.New(config.WithSubSection[ConfigSection]("ServiceConfig")),
		),
		fx.Invoke(
			// Invoke the NewService function with the loaded configuration.
			NewService,
		),
	).Run()
}
```

## Advanced Example

In this example, we demonstrate how to use multiple configuration sections and inject them into different parts of the application:

```go
package main

import (
	"fmt"
	"time"

	"schneider.vip/config"
	"schneider.vip/fxconfig"

	"go.uber.org/fx"
)

// DatabaseConfig represents the database configuration section.
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// APIConfig represents the API configuration section.
type APIConfig struct {
	Endpoint string
	Timeout  time.Duration
}

// NewDatabaseService is a constructor that uses the database configuration.
func NewDatabaseService(loader config.Dynamic[DatabaseConfig]) {
	for {
		cfg := loader.Load()
		fmt.Println("Database Config:", cfg.Host, cfg.Port)
		time.Sleep(1 * time.Second)
	}
}

// NewAPIService is a constructor that uses the API configuration.
func NewAPIService(loader config.Dynamic[APIConfig]) {
	for {
		cfg := loader.Load()
		fmt.Println("API Config:", cfg.Endpoint, cfg.Timeout)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fx.New(
		fx.Provide(
			// Provide dynamic configuration loaders for different sections.
			fxconfig.New(config.WithSubSection[DatabaseConfig]("Database")),
			fxconfig.New(config.WithSubSection[APIConfig]("API")),
		),
		fx.Invoke(
			// Invoke the service constructors with the loaded configurations.
			NewDatabaseService,
			NewAPIService,
		),
	).Run()
}
```

## Configuration File Example
Assuming you have a config.yml file with the following content:

```yaml
Database:
  Host: "localhost"
  Port: 5432
  Username: "user"
  Password: "pass"

API:
  Endpoint: "https://api.example.com"
  Timeout: "5s"
```

The above code will load the `Database` and `API` sections from the `config.yml` file and inject them into the respective services.

## Example: Using a Global Configuration Without a Dynamic Config Loader

In this example, we demonstrate how to use `fxconfig` to load a global configuration and inject it into different parts of the application. Unlike the dynamic config loader, this approach loads the configuration statically, meaning it is loaded once at application startup and remains unchanged.

```go
package main

import (
	"fmt"
	"time"

	"schneider.vip/fxconfig"

	"go.uber.org/fx"
)

// Config represents the global configuration structure.
type Config struct {
	fx.Out
	Database DatabaseConfig
	API      APIConfig
}

// DatabaseConfig represents the database configuration section.
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// APIConfig represents the API configuration section.
type APIConfig struct {
	Endpoint string
	Timeout  time.Duration
}

// NewDatabaseService is a constructor that uses the database configuration.
func NewDatabaseService(cfg DatabaseConfig) {
	for {
		fmt.Println("Database Config:", cfg.Host, cfg.Port)
		time.Sleep(1 * time.Second)
	}
}

// NewAPIService is a constructor that uses the API configuration.
func NewAPIService(cfg APIConfig) {
	for {
		fmt.Println("API Config:", cfg.Endpoint, cfg.Timeout)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fx.New(
		fx.Provide(
			// Provide static configuration for different sections.
			fxconfig.New[Config](),
		),
		fx.Invoke(
			// Invoke the service constructors with the loaded configurations.
			NewDatabaseService,
			NewAPIService,
		),
	).Run()
}
```

## API Reference

## `fxconfig.New`

```go
func New[T any](opts ...config.Option[T]) func() (config.Dynamic[T], T)
```

`New` returns a constructor function that creates a `config.Dynamic[T]` loader and a parsed configuration of type `T`. This function is designed to be used with `fx.Provide`.

## `config.Dynamic[T]`

```go
type Dynamic[T any] interface {
	Load() T
}
```

`Dynamic[T]` is an interface that provides a method to load the latest configuration of type `T`.


