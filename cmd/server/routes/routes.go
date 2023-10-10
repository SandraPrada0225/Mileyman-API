package routes

import (
	getcarritobyidhandler "Mileyman-API/cmd/server/handlers/get_carrito_by_id"
	getdulcebycodehandler "Mileyman-API/cmd/server/handlers/get_dulce_by_code"
	getfiltroshandler "Mileyman-API/cmd/server/handlers/get_filtros"
	"Mileyman-API/cmd/server/handlers/ping"
	"Mileyman-API/internal/repositories/carritos"
	"Mileyman-API/internal/repositories/categorias"
	"Mileyman-API/internal/repositories/dulces"
	"Mileyman-API/internal/repositories/marcas"
	"Mileyman-API/internal/repositories/presentaciones"
	getcarritobyidusecase "Mileyman-API/internal/use_case/get_carrito_by_id"
	getdulcebycodeusecase "Mileyman-API/internal/use_case/get_dulce_by_code"
	getfiltrosusecase "Mileyman-API/internal/use_case/get_filtros"

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
	dulcesProvider := dulces.Repository{
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

	carritosProvider := carritos.Repository{
		DB: r.db,
	}

	// UseCase
	getDulceByCodeUseCase := getdulcebycodeusecase.Implementation{
		DulcesProvider:     dulcesProvider,
		CategoriasProvider: categoriasProvider,
	}

	getFiltros := getfiltrosusecase.Implementation{
		CategoriasProvider:     categoriasProvider,
		MarcasProvider:         marcaProvider,
		PresentacionesProvider: presentacionProvider,
	}

	getCarritoByIDUseCase := getcarritobyidusecase.Implementation{
		CarritoProvider:    carritosProvider,
		DulcesProvider:     dulcesProvider,
		CategoriasProvider: categoriasProvider,
	}

	// Handlers
	getDulceByCodeHandler := getdulcebycodehandler.GetDulceByCode{
		UseCase: getDulceByCodeUseCase,
	}

	getFiltrosHandler := getfiltroshandler.GetFiltros{
		UseCase: getFiltros,
	}

	getCarritoByIDHandler := getcarritobyidhandler.GetDetalleCarritoById{
		UseCase: getCarritoByIDUseCase,
	}

	// endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	dulcesGrupo.GET("/:codigo", getDulceByCodeHandler.Handle())

	filtrosGrupo := r.rg.Group("/filtros")
	filtrosGrupo.GET("/", getFiltrosHandler.Handle())

	carritosGroup := r.rg.Group("/carritos")
	carritosGroup.GET(":id", getCarritoByIDHandler.Handle())
}
