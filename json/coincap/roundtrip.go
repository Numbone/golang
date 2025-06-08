package coincap

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (rt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(rt.logger, "%s %s %s\n", time.Now().Format(time.ANSIC), req.Method, req.URL)
	return rt.next.RoundTrip(req)
}
