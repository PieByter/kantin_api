package routes

import (
	"kantin_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PUT("/users/:id/change-password", controllers.UpdatePassword)

	r.POST("/bagian", controllers.CreateBagian)
	r.GET("/bagian", controllers.GetBagians)
	r.GET("/bagian/:id", controllers.GetBagianByID)
	r.PUT("/bagian/:id", controllers.UpdateBagian)
	r.DELETE("/bagian/:id", controllers.DeleteBagian)

	r.POST("/jabatan", controllers.CreateJabatan)
	r.GET("/jabatan", controllers.GetJabatans)
	r.GET("/jabatan/:id", controllers.GetJabatanByID)
	r.PUT("/jabatan/:id", controllers.UpdateJabatan)
	r.DELETE("/jabatan/:id", controllers.DeleteJabatan)

	r.POST("/penempatan", controllers.CreatePenempatan)
	r.GET("/penempatan", controllers.GetPenempatans)
	r.GET("/penempatan/:id", controllers.GetPenempatanByID)
	r.PUT("/penempatan/:id", controllers.UpdatePenempatan)
	r.DELETE("/penempatan/:id", controllers.DeletePenempatan)

	r.GET("/makanans", controllers.GetMakanans)
	r.POST("/makanans", controllers.CreateMakanan)
	r.GET("/makanans/:id", controllers.GetMakananByID)
	r.PUT("/makanans/:id", controllers.UpdateMakanan)
	r.DELETE("/makanans/:id", controllers.DeleteMakanan)

	r.GET("/order_kupons", controllers.GetOrderKupons)
	r.POST("/order_kupons", controllers.CreateOrderKupon)
	r.GET("/order_kupons/:id", controllers.GetOrderKuponByID)
	r.PUT("/order_kupons/:id", controllers.UpdateOrderKupon)
	r.DELETE("/order_kupons/:id", controllers.DeleteOrderKupon)

	r.GET("/kupons", controllers.GetKupons)
	r.POST("/kupons", controllers.CreateKupon)
	r.GET("/kupons/:id", controllers.GetKuponByID)
	r.PUT("/kupons/:id", controllers.UpdateKupon)
	r.DELETE("/kupons/:id", controllers.DeleteKupon)
	r.GET("/kupons/status/:status", controllers.GetKuponsByStatus)

	r.GET("/order_kupon_items", controllers.GetOrderKuponItems)
	r.POST("/order_kupon_items", controllers.CreateOrderKuponItem)
	r.GET("/order_kupon_items/:id", controllers.GetOrderKuponItemByID)
	r.PUT("/order_kupon_items/:id", controllers.UpdateOrderKuponItem)
	r.DELETE("/order_kupon_items/:id", controllers.DeleteOrderKuponItem)

	r.GET("/order_regulers", controllers.GetOrderRegulers)
	r.POST("/order_regulers", controllers.CreateOrderReguler)
	r.GET("/order_regulers/:id", controllers.GetOrderRegulerByID)
	r.PUT("/order_regulers/:id", controllers.UpdateOrderReguler)
	r.DELETE("/order_regulers/:id", controllers.DeleteOrderReguler)

	r.GET("/order_reguler_items", controllers.GetOrderRegulerItems)
	r.POST("/order_reguler_items", controllers.CreateOrderRegulerItem)
	r.GET("/order_reguler_items/:id", controllers.GetOrderRegulerItemByID)
	r.PUT("/order_reguler_items/:id", controllers.UpdateOrderRegulerItem)
	r.DELETE("/order_reguler_items/:id", controllers.DeleteOrderRegulerItem)

	r.GET("/pembelian_kupons", controllers.GetPembelianKupons)
	r.POST("/pembelian_kupons", controllers.CreatePembelianKupon)
	r.GET("/pembelian_kupons/:id", controllers.GetPembelianKuponByID)
	r.PUT("/pembelian_kupons/:id", controllers.UpdatePembelianKupon)
	r.DELETE("/pembelian_kupons/:id", controllers.DeletePembelianKupon)

	r.GET("/transaksi_regulers", controllers.GetTransaksiRegulers)
	r.POST("/transaksi_regulers", controllers.CreateTransaksiReguler)
	r.GET("/transaksi_regulers/:id", controllers.GetTransaksiRegulerByID)
	r.PUT("/transaksi_regulers/:id", controllers.UpdateTransaksiReguler)
	r.DELETE("/transaksi_regulers/:id", controllers.DeleteTransaksiReguler)

	r.GET("/transaksi_reguler_details", controllers.GetTransaksiRegulerDetails)
	r.POST("/transaksi_reguler_details", controllers.CreateTransaksiRegulerDetail)
	r.GET("/transaksi_reguler_details/:id", controllers.GetTransaksiRegulerDetailByID)
	r.PUT("/transaksi_reguler_details/:id", controllers.UpdateTransaksiRegulerDetail)
	r.DELETE("/transaksi_reguler_details/:id", controllers.DeleteTransaksiRegulerDetail)

	r.GET("/transaksi_specials", controllers.GetTransaksiSpecials)
	r.POST("/transaksi_specials", controllers.CreateTransaksiSpecial)
	r.GET("/transaksi_specials/:id", controllers.GetTransaksiSpecialByID)
	r.PUT("/transaksi_specials/:id", controllers.UpdateTransaksiSpecial)
	r.DELETE("/transaksi_specials/:id", controllers.DeleteTransaksiSpecial)
}
