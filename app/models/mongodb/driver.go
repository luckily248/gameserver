package mongodb

var MaxPool int
var PATH string
var DBNAME string

func CheckAndInitServiceConnection() {
	if service.baseSession == nil {
		println("PATH:" + PATH)
		service.URL = PATH
		err := service.New()
		if err != nil {
			panic(err)
		}
	}
}
