package routes

import (
	"kantin_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/users", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetUsers)
	r.GET("/users/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetUserByID)
	r.PUT("/users/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeleteUser)
	r.PUT("/users/:id/change-password", controllers.AuthMiddleware(), controllers.UpdatePassword)

	r.POST("/bagian", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreateBagian)
	r.GET("/bagian", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetBagians)
	r.GET("/bagian/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetBagianByID)
	r.PUT("/bagian/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdateBagian)
	r.DELETE("/bagian/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeleteBagian)

	r.POST("/jabatan", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreateJabatan)
	r.GET("/jabatan", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetJabatans)
	r.GET("/jabatan/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetJabatanByID)
	r.PUT("/jabatan/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdateJabatan)
	r.DELETE("/jabatan/:id", controllers.AuthMiddleware(), controllers.DeleteJabatan)

	r.POST("/penempatan", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreatePenempatan)
	r.GET("/penempatan", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetPenempatans)
	r.GET("/penempatan/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetPenempatanByID)
	r.PUT("/penempatan/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdatePenempatan)
	r.DELETE("/penempatan/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeletePenempatan)

	r.GET("/makanans", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetMakanans)
	r.POST("/makanans", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreateMakanan)
	r.GET("/makanans/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetMakananByID)
	r.PUT("/makanans/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdateMakanan)
	r.DELETE("/makanans/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeleteMakanan)

	r.GET("/order_kupons", controllers.AuthMiddleware(), controllers.GetOrderKupons)
	r.POST("/order_kupons", controllers.AuthMiddleware(), controllers.CreateOrderKupon)
	r.GET("/order_kupons/:id", controllers.AuthMiddleware(), controllers.GetOrderKuponByID)
	r.PUT("/order_kupons/:id", controllers.AuthMiddleware(), controllers.UpdateOrderKupon)
	r.DELETE("/order_kupons/:id", controllers.AuthMiddleware(), controllers.DeleteOrderKupon)

	r.GET("/kupons", controllers.AuthMiddleware(), controllers.GetKupons)
	r.POST("/kupons", controllers.AuthMiddleware(), controllers.CreateKupon)
	r.GET("/kupons/:id", controllers.AuthMiddleware(), controllers.GetKuponByID)
	r.PUT("/kupons/:id", controllers.AuthMiddleware(), controllers.UpdateKupon)
	r.DELETE("/kupons/:id", controllers.AuthMiddleware(), controllers.DeleteKupon)
	r.GET("/kupons/status/:status", controllers.AuthMiddleware(), controllers.GetKuponsByStatus)

	r.GET("/order_kupon_items", controllers.AuthMiddleware(), controllers.GetOrderKuponItems)
	r.POST("/order_kupon_items", controllers.AuthMiddleware(), controllers.CreateOrderKuponItem)
	r.GET("/order_kupon_items/:id", controllers.AuthMiddleware(), controllers.GetOrderKuponItemByID)
	r.PUT("/order_kupon_items/:id", controllers.AuthMiddleware(), controllers.UpdateOrderKuponItem)
	r.DELETE("/order_kupon_items/:id", controllers.AuthMiddleware(), controllers.DeleteOrderKuponItem)

	r.GET("/order_regulers", controllers.AuthMiddleware(), controllers.GetOrderRegulers)
	r.POST("/order_regulers", controllers.AuthMiddleware(), controllers.CreateOrderReguler)
	r.GET("/order_regulers/:id", controllers.AuthMiddleware(), controllers.GetOrderRegulerByID)
	r.PUT("/order_regulers/:id", controllers.AuthMiddleware(), controllers.UpdateOrderReguler)
	r.DELETE("/order_regulers/:id", controllers.AuthMiddleware(), controllers.DeleteOrderReguler)

	r.GET("/order_reguler_items", controllers.AuthMiddleware(), controllers.GetOrderRegulerItems)
	r.POST("/order_reguler_items", controllers.AuthMiddleware(), controllers.CreateOrderRegulerItem)
	r.GET("/order_reguler_items/:id", controllers.AuthMiddleware(), controllers.GetOrderRegulerItemByID)
	r.PUT("/order_reguler_items/:id", controllers.AuthMiddleware(), controllers.UpdateOrderRegulerItem)
	r.DELETE("/order_reguler_items/:id", controllers.AuthMiddleware(), controllers.DeleteOrderRegulerItem)

	r.GET("/pembelian_kupons", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetPembelianKupons)
	r.POST("/pembelian_kupons", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreatePembelianKupon)
	r.GET("/pembelian_kupons/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetPembelianKuponByID)
	r.PUT("/pembelian_kupons/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdatePembelianKupon)
	r.DELETE("/pembelian_kupons/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeletePembelianKupon)

	r.GET("/transaksi_regulers", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetTransaksiRegulers)
	r.POST("/transaksi_regulers", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreateTransaksiReguler)
	r.GET("/transaksi_regulers/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetTransaksiRegulerByID)
	r.PUT("/transaksi_regulers/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdateTransaksiReguler)
	r.DELETE("/transaksi_regulers/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeleteTransaksiReguler)

	r.GET("/transaksi_reguler_details", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetTransaksiRegulerDetails)
	r.POST("/transaksi_reguler_details", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreateTransaksiRegulerDetail)
	r.GET("/transaksi_reguler_details/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetTransaksiRegulerDetailByID)
	r.PUT("/transaksi_reguler_details/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdateTransaksiRegulerDetail)
	r.DELETE("/transaksi_reguler_details/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeleteTransaksiRegulerDetail)

	r.GET("/transaksi_specials", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetTransaksiSpecials)
	r.POST("/transaksi_specials", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.CreateTransaksiSpecial)
	r.GET("/transaksi_specials/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.GetTransaksiSpecialByID)
	r.PUT("/transaksi_specials/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.UpdateTransaksiSpecial)
	r.DELETE("/transaksi_specials/:id", controllers.AuthMiddleware(), controllers.AdminMiddleware(), controllers.DeleteTransaksiSpecial)
}
