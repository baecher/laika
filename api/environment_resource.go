package api

import (
	"net/url"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/laika/store"
)

type EnvironmentResource struct {
	store store.Store
	stats *statsd.Client
}

func NewEnvironmentResource(store store.Store, stats *statsd.Client) *EnvironmentResource {
	return &EnvironmentResource{store, stats}
}

func (r *EnvironmentResource) List(c echo.Context) error {
	s, err := r.store.State()
	if err != nil {
		return InternalServerError(c, err)
	}

	return OK(c, s.Environments)
}

func (r *EnvironmentResource) GetFeatures(c echo.Context) error {
	s, err := r.store.State()
	if err != nil {
		return InternalServerError(c, err)
	}

	env, err := url.QueryUnescape(c.Param("name"))
	if err != nil {
		return BadRequest(c, "Bad environment name")
	}

	features := s.GetFeatures(env)
	if features == nil {
		return NotFound(c)
	}

	return OK(c, features)
}
