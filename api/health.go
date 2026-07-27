package main

import verklgoruntime "verkill-go-runtime"

func handler(req verklgoruntime.Request) verklgoruntime.Response {
	return verklgoruntime.Response{
		StatusCode: 200,
		Headers:    map[string]string{"content-type": "application/json"},
		Body:       `{"status":"ok"}`,
	}
}

func main() {
	verklgoruntime.Serve(handler)
}
