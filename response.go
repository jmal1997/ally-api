package ally_api

type Response struct {
	ResponseId  string  `xml:"id,attr"`
	ElapsedTime float64 `xml:"elapsedtime"`
	Error       string  `xml:"error"`
}
