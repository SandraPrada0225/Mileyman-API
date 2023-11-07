package routes

import (
	getcarritobyidhandler "Mileyman-API/cmd/server/handlers/get_carrito_by_id"
	getdulcebycodehandler "Mileyman-API/cmd/server/handlers/get_dulce_by_code"
	getfiltroshandler "Mileyman-API/cmd/server/handlers/get_filtros"
	"Mileyman-API/cmd/server/handlers/ping"
	purchasecarritohandler "Mileyman-API/cmd/server/handlers/purchase_carrito"
	updatecarritohandler "Mileyman-API/cmd/server/handlers/update_carrito"
	"Mileyman-API/internal/repositories/carritos"
	"Mileyman-API/internal/repositories/categorias"
	"Mileyman-API/internal/repositories/dulces"
	"Mileyman-API/internal/repositories/marcas"
	"Mileyman-API/internal/repositories/presentaciones"
	"Mileyman-API/internal/repositories/usuarios"
	"Mileyman-API/internal/repositories/ventas"
	getcarritobyidusecase "Mileyman-API/internal/use_case/get_carrito_by_id"
	getdulcebycodeusecase "Mileyman-API/internal/use_case/get_dulce_by_code"
	getfiltrosusecase "Mileyman-API/internal/use_case/get_filtros"
	purchasecarritousecase "Mileyman-API/internal/use_case/purchase_carrito"
	updatecarritousecase "Mileyman-API/internal/use_case/update_carrito"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct{}

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *gorm.DB
}

func NewRouter(eng *gin.Engine, db *gorm.DB) Router {
	return &router{
		eng: eng,
		db:  db,
	}
}

func (r router) MapRoutes() {
	r.rg = r.eng.Group("/api")

	pingHandler := ping.Ping{}
	r.rg.GET("/ping", pingHandler.Handle())

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

	ventasProvider := ventas.Repository{
		DB: r.db,
	}

	usuariosProvider := usuarios.Repository{
		DB: r.db,
	}

	getDulceByCodeUseCase := getdulcebycodeusecase.Implementation{
		DulcesProvider:     dulcesProvider,
		CategoriasProvider: categoriasProvider,
	}

	getFiltros := getfiltrosusecase.Implementation{
		CategoriasProvider:     categoriasProvider,
		MarcasProvider:         marcaProvider,
		PresentacionesProvider: presentacionProvider,
	}

	updatecarrito := updatecarritousecase.Implementation{
		CarritosProvider: carritosProvider,
		DulcesProvider:   dulcesProvider,
	}

	getCarritoByIDUseCase := getcarritobyidusecase.Implementation{
		CarritoProvider:    carritosProvider,
		DulcesProvider:     dulcesProvider,
		CategoriasProvider: categoriasProvider,
	}

	purchaseCarritoUseCase := purchasecarritousecase.Implementation{
		CarritosProvider: carritosProvider,
		UsuariosProvider: usuariosProvider,
		VentasProvider:   ventasProvider,
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

	purchaseCarritoHandler := purchasecarritohandler.PurchaseCarrito{
		UseCase: purchaseCarritoUseCase,
	}

	updateCarritoHandler := updatecarritohandler.UpdateCarrito{
		UseCase: updatecarrito,
	}

	// endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	dulcesGrupo.GET("/:codigo", getDulceByCodeHandler.Handle())

	filtrosGrupo := r.rg.Group("/filtros")
	filtrosGrupo.GET("/", getFiltrosHandler.Handle())

	carritosGroup := r.rg.Group("/carritos")
	carritosGroup.GET(":id", getCarritoByIDHandler.Handle())
	carritosGroup.PUT(":id/comprar", purchaseCarritoHandler.Handle())
	carritosGroup.PUT("/:id", updateCarritoHandler.Handle())
}
