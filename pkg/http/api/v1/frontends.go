package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"gitlab.com/infra.run/public/b3scale/pkg/store"
)

// FrontendsList will list all frontends known
// to the cluster or within the user scope.
func FrontendsList(c echo.Context) error {
	ctx := c.(*APIContext)
	ref := ctx.FilterAccountRef()
	reqCtx := ctx.Ctx()

	q := store.Q()
	if ref != nil {
		q.Where("account_ref = ?", *ref)
	}
	tx, err := store.ConnectionFromContext(reqCtx).Begin(reqCtx)
	if err != nil {
		return err
	}
	defer tx.Rollback(reqCtx)
	frontends, err := store.GetFrontendStates(reqCtx, tx, q)
	return c.JSON(http.StatusOK, frontends)
}

// FrontendCreate will add a new frontend to the cluster.
func FrontendCreate(c echo.Context) error {
	ctx := c.(*APIContext)
	cctx := ctx.Ctx()
	accountRef := ctx.AccountRef()
	isAdmin := ctx.HasScope(ScopeAdmin)

	f := &store.FrontendState{}
	if err := c.Bind(f); err != nil {
		return err
	}

	frontend := store.InitFrontendState(&store.FrontendState{
		Frontend: f.Frontend,
		Settings: f.Settings,
		Active:   f.Active,
	})

	if err := frontend.Validate(); err != nil {
		return err
	}

	tx, err := store.ConnectionFromContext(cctx).Begin(cctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(cctx)

	return nil
}

// FrontendRetrieve will retrieve a single frontend
// identified by ID.
func FrontendRetrieve(c echo.Context) error {
	return nil
}

// FrontendDestroy will remove a frontend from the cluster.
// The frontend is identified by ID.
func FrontendDestroy(c echo.Context) error {
	return nil
}

// FrontendUpdate will update the frontend with values
// provided by the request. Only keys provided will
// be updated.
func FrontendUpdate(c echo.Context) error {
	return nil
}
