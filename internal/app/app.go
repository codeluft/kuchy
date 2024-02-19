package app

import (
	"context"
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	otelsdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"kuchy/internal/app/config"
	"kuchy/pages"
	"kuchy/pages/layout"
	"kuchy/pkg/dic"
	"net/http"
	"os"
	"time"
)

type app struct {
	serverMux *http.ServeMux
	container dic.Container
}

// New creates a new app
func New(staticFs *embed.FS) *app {
	mux := http.NewServeMux()
	ctx := context.Background()
	container := dic.New()
	_ = godotenv.Load(".env")

	registerRequiredDependencies(ctx, container, staticFs)

	config.RegisterDependencies(ctx, container)
	config.RegisterRoutes(mux, container)

	return &app{
		serverMux: mux,
		container: container,
	}
}

// ServeHTTP serves HTTP
func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	span := a.container.Get("tracer").(func() trace.Span)()

	traceId := r.Header.Get("X-Trace-ID")
	if traceId == "" {
		traceId = span.SpanContext().TraceID().String()
	}

	spanId := r.Header.Get("X-Span-ID")
	if spanId == "" {
		spanId = span.SpanContext().SpanID().String()
	}

	logger := a.container.Get("logger").(*zap.Logger).
		With(zap.String("traceId", traceId), zap.String("spanId", spanId))

	logger.Info("Request received", zap.String("method", r.Method), zap.String("url", r.URL.String()))
	a.serverMux.ServeHTTP(w, r)
}

// Run application server
func (a *app) Run(addr string) {
	logger := a.container.Get("logger").(*zap.Logger)
	logger.Info(fmt.Sprintf("Server is running on %s", addr))
	logger.Error("Shutting down server", zap.Error(http.ListenAndServe(addr, a)))
}

func registerRequiredDependencies(ctx context.Context, c dic.Container, staticFs *embed.FS) {
	c.Register("tracer", func(c dic.Container) interface{} {
		return func() oteltrace.Span {
			_, span := otelsdktrace.NewTracerProvider().Tracer("codeluft/kuchy").Start(ctx, "kuchy-app")
			defer span.End()

			return span
		}
	})

	c.Register("logger", func(c dic.Container) interface{} {
		var logger *zap.Logger

		if env := os.Getenv("ENV"); env == "development" {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction()
		}

		return logger.WithLazy(zap.String("datetime", time.Now().UTC().String()))
	})

	c.Register("staticFs", func(c dic.Container) interface{} {
		return staticFs
	})

	c.Register("translator", func(c dic.Container) interface{} {
		return translator("en")
	})

	c.Register("layout", func(c dic.Container) interface{} {
		return layout.New(
			c.Get("translator").(interface{ Translate(string) string }),
			c.Get("staticFs").(*embed.FS),
		)
	})

	c.Register("pages", func(c dic.Container) interface{} {
		return pages.New(
			c.Get("translator").(interface{ Translate(string) string }),
			c.Get("staticFs").(*embed.FS),
			c.Get("layout").(*layout.Layout),
		)
	})
}

// temporary translator (mock)
type translator string

func (t translator) Translate(v string) string {
	return fmt.Sprintf("%s:%s:%s", t, v, t)
}
