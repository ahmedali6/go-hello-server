package main

import (
	"fmt"
	"time"

	verklgoruntime "verkill-go-runtime"
)

func handler(req verklgoruntime.Request) verklgoruntime.Response {
	return verklgoruntime.Response{
		StatusCode: 200,
		Headers:    map[string]string{"content-type": "application/json"},
		Body: fmt.Sprintf(
			`{"message":"Hello from Go!","path":%q,"method":%q,"timestamp":%q}`,
			req.Path, req.Method, time.Now().UTC().Format(time.RFC3339),
		),
	}
}

func main() {
	verklgoruntime.Serve(handler)
}
