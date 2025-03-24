package model

type HTTPMiddlewareManagerRequest struct {
	Method            string
	Path              string
	UseAuthentication bool
}
