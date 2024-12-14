package utils

import "net/http"

// Create http response with no content (HTTP code 204).
func NoContentResponseBuilder(httpResponseWriter http.ResponseWriter) {
	httpResponseWriter.WriteHeader(http.StatusNoContent)
	httpResponseWriter.Write(nil)
}
