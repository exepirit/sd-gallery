package server

import "github.com/gin-gonic/gin"

// Bindable is a object, that can bind to router.
type Bindable interface {
	Bind(r gin.IRoutes)
}
