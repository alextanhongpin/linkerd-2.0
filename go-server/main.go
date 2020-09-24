package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/label"

	"go.opentelemetry.io/otel/exporters/trace/jaeger"
)

// JAEGER_AGENT_HOST=linkerd-jaeger.linkerd:14268/api/traces
var jaegerAgentHost = os.Getenv("JAEGER_AGENT_HOST")

func initTracer() func() {
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(jaegerAgentHost),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "trace-demo",
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.Float64("float", 312.23),
			},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return func() {
		flush()
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	stop := initTracer()
	defer stop()

	log.Println("listening to port *:8080. press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	timeout := r.URL.Query().Get("timeout")
	if timeout == "true" {
		time.Sleep(3 * time.Second)
	}
	retry := r.URL.Query().Get("retry")
	if retry == "true" {
		if rand.Float32() < 0.5 {
			log.Fatal("bad request")
		}
	}
	tracer := global.Tracer("component/indexHandler")
	ctx, span := tracer.Start(r.Context(), "indexHandler")
	defer span.End()
	greet(ctx)

	fmt.Fprint(w, `{"hello": "world"}`)
}

func greet(ctx context.Context) {
	tr := global.Tracer("component/bar")
	_, span := tr.Start(ctx, "bar")
	defer span.End()
}
