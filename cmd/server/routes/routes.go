package routes

import (
	"Mileyman-API/cmd/server/handlers"
	"Mileyman-API/internal/repositories/categorias"
	"Mileyman-API/internal/repositories/dulces"
	"Mileyman-API/internal/repositories/marcas"
	"Mileyman-API/internal/repositories/presentaciones"
	getdulcebycode "Mileyman-API/internal/use_case/get_dulce_by_code"
	getfiltros "Mileyman-API/internal/use_case/get_filtros"

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

	// ping
	ping := handlers.Ping{}
	r.rg.GET("/ping", ping.Handle())

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

	categoriaProvider := categorias.Repository{
		DB: r.db,
	}

	// UseCase
	getDulceByCodeUseCase := getdulcebycode.Implementation{
		DulcesProvider: dulceProvider,
	}

	getFitros := getfiltros.Implementation{
		CategoriasProvider:     categoriaProvider,
		MarcasProvider:         marcaProvider,
		PresentacionesProvider: presentacionProvider,
	}

	// Handlers
	getDulceByCodeHandler := handlers.GetDulceByCode{
		UseCase: getDulceByCodeUseCase,
	}

	getFiltrosHandler := handlers.GetFiltros{
		UseCase: getFitros,
	}

	// endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	dulcesGrupo.GET("/:codigo", getDulceByCodeHandler.Handle())

	filtrosGrupo := r.rg.Group("/filtros")
	filtrosGrupo.GET("/", getFiltrosHandler.Handle())
}
