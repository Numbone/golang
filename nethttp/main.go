package nethttp

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (rt *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(rt.logger, "%s %s %s\n", time.Now().Format(time.ANSIC), req.Method, req.URL)
	return rt.next.RoundTrip(req)
}

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	//jar.SetCookies(url.URL{
	//	Host: "",
	//
	//}("http://localhost:8000"),[]*http.Cookie{})
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.StatusCode)
			fmt.Println("REDIRECT")
			return nil
		},
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Jar: jar,
	}
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
