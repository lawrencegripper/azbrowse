package tracing

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/lawrencegripper/azbrowse/internal/pkg/errorhandling"
	opentracing "github.com/opentracing/opentracing-go"
	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
)

// StartTracing starts an OpenTracing UI on localhost:8700
// inspired by: https://medium.com/opentracing/distributed-tracing-in-10-minutes-51b378ee40f1
func StartTracing() func(opentracing.Span) string {

	store := appdash.NewMemoryStore()

	// Listen on any available TCP port locally.
	l, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0})
	if err != nil {
		log.Fatal(err)
	}
	collectorPort := l.Addr().(*net.TCPAddr).Port
	collectorAdd := fmt.Sprintf(":%d", collectorPort)

	// Start an Appdash collection server that will listen for spans and
	// annotations and add them to the local collector (stored in-memory).
	cs := appdash.NewServer(l, appdash.NewLocalCollector(store))
	go cs.Start()

	tracer := appdashot.NewTracer(appdash.NewRemoteCollector(collectorAdd))
	opentracing.InitGlobalTracer(tracer)

	getURLFunc := func(s opentracing.Span) string {
		// Start the webui for viewing traces
		appdashPort := 8700
		appdashURLStr := fmt.Sprintf("http://localhost:%d", appdashPort)
		appdashURL, err := url.Parse(appdashURLStr)
		if err != nil {
			log.Fatalf("Error parsing %s: %s", appdashURLStr, err)
		}

		// Start the web UI in a separate goroutine.
		tapp, err := traceapp.New(nil, appdashURL)
		if err != nil {
			log.Fatal(err)
		}
		tapp.Store = store
		tapp.Queryer = store
		go func() {
			// recover from panic, if one occurrs, and leave terminal usable
			defer errorhandling.RecoveryWithCleanup()

			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", appdashPort), tapp))
		}()
		traces, err := tapp.Queryer.Traces(appdash.TracesOpts{})
		if err != nil {
			log.Panic(err)
		}
		if len(traces) < 1 {
			return appdashURLStr
		}
		return appdashURLStr + "/traces/" + traces[0].ID.Trace.String()
	}

	return getURLFunc
}
