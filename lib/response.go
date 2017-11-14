package lib

type Response struct {
	ResponseId  string  `json:"id" xml:"id,attr"`
	ElapsedTime float64 `json:"elapsedtime" xml:"elapsedtime,omitempty"`
	Error       string  `json:"error" xml:"error,omitempty"`
}
