package v1

import (
	// "net/http"

	"github.com/labstack/echo/v4"
	//	"github.com/rs/zerolog/log"
	// "gitlab.com/infra.run/public/b3scale/pkg/store"
)

// BackendsList will list all frontends known
// to the cluster or within the user scope.
func BackendsList(c echo.Context) error {
	return nil
}

// BackendCreate will add a new frontend to the cluster.
func BackendCreate(c echo.Context) error {
	return nil
}

// BackendRetrieve will retrieve a single frontend
// identified by ID.
func BackendRetrieve(c echo.Context) error {
	return nil
}

// BackendDestroy will remove a frontend from the cluster.
// The frontend is identified by ID.
func BackendDestroy(c echo.Context) error {
	return nil
}

// BackendUpdate will update the frontend with values
// provided by the request. Only keys provided will
// be updated.
func BackendUpdate(c echo.Context) error {
	return nil
}
