package driver

import "go-rest/util/runner"

func PrepareAll() {
	runner.PrepareRuntime()

	conf, _ := Config()
	Database(conf.Databases)
	Validator()
	NewRBAC()
}
