package routes

import (
	getdulcebycodehandler "Mileyman-API/cmd/server/handlers/get_dulce_by_code"
	updatecarritohandler "Mileyman-API/cmd/server/handlers/update_carrito"
	getfiltroshandler "Mileyman-API/cmd/server/handlers/get_filtros"
	"Mileyman-API/cmd/server/handlers/ping"
	carritos "Mileyman-API/internal/repositories/Carritos"
	"Mileyman-API/internal/repositories/categorias"
	"Mileyman-API/internal/repositories/dulces"
	"Mileyman-API/internal/repositories/marcas"
	"Mileyman-API/internal/repositories/presentaciones"
	getdulcebycodeusecase "Mileyman-API/internal/use_case/get_dulce_by_code"
	getfiltrosusecase "Mileyman-API/internal/use_case/get_filtros"
	updatecarrito "Mileyman-API/internal/use_case/update_carrito"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct{}

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine      // crea
	rg  *gin.RouterGroup // exponer los handlers en la url//guarda
	db  *gorm.DB         // inyectar la base de datos
}

func NewRouter(eng *gin.Engine, db *gorm.DB) Router {
	return &router{
		eng: eng,
		db:  db,
	}
}

func (r router) MapRoutes() {
	r.rg = r.eng.Group("/api")

	// pingHandler
	pingHandler := ping.Ping{}
	r.rg.GET("/ping", pingHandler.Handle())

	// providers
	dulceProvider := dulces.Repository{
		DB: r.db,
	}

	presentacionProvider := presentaciones.Repository{
		DB: r.db,
	}

	marcaProvider := marcas.Repository{
		DB: r.db,
	}

	categoriasProvider := categorias.Repository{
		DB: r.db,
	}
	carritosdulcesProvider := carritos.Repository{
		DB: r.db,
	}

	// UseCase
	getDulceByCodeUseCase := getdulcebycodeusecase.Implementation{
		DulcesProvider:     dulceProvider,
		CategoriasProvider: &categoriasProvider,
	}

	getFitros := getfiltrosusecase.Implementation{
		CategoriasProvider:     categoriasProvider,
		MarcasProvider:         marcaProvider,
		PresentacionesProvider: presentacionProvider,
	}

	updatecarrito := updatecarrito.Implementation{
		CarritosDulcesProvider: carritosdulcesProvider,
		DulcesProvider: dulceProvider,

	} 

	// Handlers
	getDulceByCodeHandler := getdulcebycodehandler.GetDulceByCode{
		UseCase: getDulceByCodeUseCase,
	}

	getFiltrosHandler := getfiltroshandler.GetFiltros{
		UseCase: getFitros,
	}

	updateCarritoHandler := updatecarritohandler.UpdateCarrito{
		UseCase: updatecarrito,
	}

	// endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	dulcesGrupo.GET("/:codigo", getDulceByCodeHandler.Handle())

	filtrosGrupo := r.rg.Group("/filtros")
	filtrosGrupo.GET("/", getFiltrosHandler.Handle())

	carritos := r.rg.Group("/carritos")
	carritos.PUT("/:id", updateCarritoHandler.Handle())
}
