package config

const (
	// Service ports
	// ItemSaverPort = 1234
	// WorkerPort0   = 9000

	// ElasticSearch
	ElasticIndex = "dating_profile"
	// RPC Endpoints
	ItemSaverRpc      = "ItemSaverService.Save"
	CrawlerServiceRpc = "CrawlerService.Process"

	// Parser names
	ParseCity      = "ParseCity"
	ParserCityList = "ParserCityList"
	ParseProfile   = "ParseProfile"
	NilParser      = "NilParser"

	// Queries Per Second
	Qps = 5
)
