package routes

import (
	getcarritobyidhandler "Mileyman-API/cmd/server/handlers/getcarritobyid"
	getdulcebycodehandler "Mileyman-API/cmd/server/handlers/getdulcebycode"
	getfiltroshandler "Mileyman-API/cmd/server/handlers/getfiltros"
	getpurchaselisthandler "Mileyman-API/cmd/server/handlers/getpurchaselistbyuserid"
	"Mileyman-API/cmd/server/handlers/ping"
	purchasecarritohandler "Mileyman-API/cmd/server/handlers/purchasecarrito"
	updatecarritohandler "Mileyman-API/cmd/server/handlers/updatecarrito"
	"Mileyman-API/internal/repositories/carritos"
	"Mileyman-API/internal/repositories/categorias"
	"Mileyman-API/internal/repositories/dulces"
	"Mileyman-API/internal/repositories/marcas"
	"Mileyman-API/internal/repositories/presentaciones"
	"Mileyman-API/internal/repositories/usuarios"
	"Mileyman-API/internal/repositories/ventas"
	getcarritobyidusecase "Mileyman-API/internal/usecase/getcarritobyid"
	getdulcebycodeusecase "Mileyman-API/internal/usecase/getdulcebycode"
	getfiltrosusecase "Mileyman-API/internal/usecase/getfiltros"
	getpurchaselistusecase "Mileyman-API/internal/usecase/getpurchaselistbyuserid"
	purchasecarritousecase "Mileyman-API/internal/usecase/purchasecarrito"
	updatecarritousecase "Mileyman-API/internal/usecase/updatecarrito"

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

	getPurchaseListByUserIDUseCase := getpurchaselistusecase.Implementation{
		VentasProvider: ventasProvider,
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

	getPurchaseListByUserIDHandler := getpurchaselisthandler.GetPurchaseListByUserID{
		UseCase: getPurchaseListByUserIDUseCase,
	}

	// endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	dulcesGrupo.GET("/:codigo", getDulceByCodeHandler.Handle())

	filtrosGrupo := r.rg.Group("/filtros")
	filtrosGrupo.GET("/", getFiltrosHandler.Handle())

	carritosGroup := r.rg.Group("/carritos")
	carritosGroup.GET(":id", getCarritoByIDHandler.Handle())
	carritosGroup.PUT(":id/comprar", purchaseCarritoHandler.Handle())
	carritosGroup.PUT(":id/comprar", purchaseCarritoHandler.Handle())
	carritosGroup.PUT(":id", updateCarritoHandler.Handle())

	usersGroup := r.rg.Group("/users")
	usersGroup.GET(":id/compras", getPurchaseListByUserIDHandler.Handle())
}
