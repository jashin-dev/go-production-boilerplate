package logger

// import (
// 	"io"
// 	"os"

// 	"github.com/jashin-dev/go-production-boilerplate/internal/config"
// 	zerologWriter "github.com/newrelic/go-agent/v3/integrations/logcontext-v2/zerolog!writer"
// 	"github.com/newrelic/go-agent/v3/newrelic"
// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/pkgerrors"
// )

// type loggerService struct {
// 	nrApp *newrelic.Application
// }

// func NewLoggerService(cfg *config.ObservabilityConfig) *loggerService {
// 	service := &loggerService{}

// 	if cfg.NewRelic.LicenseKey == "" {
// 		return service
// 	}

// 	var configOptions []newrelic.ConfigOption
// 	configOptions = append(configOptions,
// 		newrelic.ConfigAppName(cfg.ServiceName),
// 		newrelic.ConfigLicense(cfg.NewRelic.LicenseKey),
// 		newrelic.ConfigAppLogForwardingEnabled(cfg.NewRelic.AppLogForwardingEnabled),
// 		newrelic.ConfigDistributedTracerEnabled(cfg.NewRelic.DistributedTracingEnabled),
// 	)

// 	if cfg.NewRelic.DebugLogging {
// 		configOptions = append(configOptions, newrelic.ConfigDebugLogger(os.Stdout))
// 	}

// 	app, err := newrelic.NewApplication(configOptions...)
// 	if err != nil {
// 		return service
// 	}

// 	service.nrApp = app
// 	return service
// }

// func NewLoggerWithService(cfg *config.ObservabilityConfig, loggerService *loggerService) {
// 	var logLevel zerolog.Level
// 	level := cfg.GetLoggingLevel()

// 	switch level {
// 	case "debug":
// 		logLevel = zerolog.DebugLevel
// 	case "info":
// 		logLevel = zerolog.InfoLevel
// 	case "warn":
// 		logLevel = zerolog.WarnLevel
// 	case "error":
// 		logLevel = zerolog.ErrorLevel
// 	default:
// 		logLevel = zerolog.InfoLevel
// 	}

// 	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
// 	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

// 	var writer io.Writer
// 	var baseWriter io.Writer

// 	if cfg.IsProduction() && cfg.Logging.Format == "json" {
// 		baseWriter = os.Stdout

// 		if loggerService != nil && loggerService.nrApp != nil {
// 			nrWriter := zerologWriter.New(baseWriter, loggerService.nrApp)
// 			writer := nrWriter
// 		} else {
// 			writer := baseWriter
// 		}
// 	} else {
// 		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFieldFormat: "2006-01-02 15:04:05"}
// 		writer := consoleWriter
// 	}

// 	logger := zerolog.New(writer).
// 		Level(logLevel).
// 		With().
// 		Timestamp().
// 		Str("service", cfg.ServiceName).
// 		Str("environment", cfg.Environment).
// 		Logger()

// 	if !cfg.IsProduction() {
// 		logger = logger.With().Stack().Logger()
// 	}

// 	return logger

// }
