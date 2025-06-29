package main

import (
	"busManager/controller"
	"busManager/repository"
	"busManager/service"
	"embed"
	"fmt"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	//// Create an instance of the app structure
	//app := NewApp()
	//
	//// Create application with options
	//err := wails.Run(&options.App{
	//	Title:  "busManager",
	//	Width:  1024,
	//	Height: 768,
	//	AssetServer: &assetserver.Options{
	//		Assets: assets,
	//	},
	//	BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	//	OnStartup:        app.startup,
	//	Bind: []interface{}{
	//		app,
	//	},
	//})
	//
	//if err != nil {
	//	println("Error:", err.Error())
	//}
	r, err := repository.NewSqliteBusRepository("db.db")

	if err != nil {
		fmt.Println(err)
	}
	driverRepo, err := repository.NewSqliteDriverRepository("db.db")
	if err != nil {
		fmt.Println(err)
	}
	busService := service.NewBusService(r)
	driverService := service.NewDriverService(driverRepo)
	busController := controller.NewBusController(*busService)
	driverController := controller.NewDriverController(*driverService)
	jsonStr := `
	{
		"Id": "",
		"RegisterNumber": "TEST1488",
		"Brand":          "Scania PIZDEC",
		"BusModel":       "123",
		"AssemblyDate":   "2012-03-03T11:11:11Z",
		"LastRepairDate": "2012-03-03T11:11:11Z"
	}
	`
	fmt.Println(busController.Add(jsonStr))

	fmt.Println(busController.DeleteById("235235"))
	fmt.Println(busController.DeleteById("740764a5-69f2-42d0-af93-5ab5d73bbeae"))
	fmt.Println(busController.GetAll())
	jsonStr = `
	{
		"Id": "2",
		"RegisterNumber": "TEST666",
		"Brand":          "Sca2nia MANDEC",
		"BusModel":       "66622",
		"AssemblyDate":   "2022-03-03T11:11:11Z",
		"LastRepairDate": "2032-03-03T11:11:11Z"
	}`
	fmt.Println(busController.UpdateById(jsonStr))
	jsonStr = `{
		
		"Name":           "John",
		"Surname":        "Doe",
		"Patronymic":     "Ivanovich",
		"BirthDate":      "2032-03-03T11:11:11Z",
		"PassportSeries": "AB123456",
		"Snils":          "123-456-789 00",
		"LicenseSeries":  "CD789012"
	}`
	fmt.Println(driverController.Add(jsonStr))
	//jsonStr = `{
	//
	//	"Lat":           "32.21",
	//	"Long":        "33.11",
	//	"Name":     "Село пердяино",
	//	}`
	//fmt.Println(driverController.Add(jsonStr))
}
