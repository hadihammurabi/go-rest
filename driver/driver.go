package driver

func Init() {
	initRuntime()
	initConfig()
	initDatabase()
	initValidator()
	initRBAC()
}
