package api

// Structures for the https://api.qubole.com/api/latest/scheduler json responses
// Built from v1.2 of the API

type Scheduler struct {
	PagingInfo PagingInfo `json:"paging_info"`
	Schedules  []Schedule `json:"schedules"`
}

type Schedule struct {
	Bitmap    int  `json:"bitmap"`
	CanNotify bool `json:"can_notify"`
	Command   struct {
		ApproxAggregations bool        `json:"approx_aggregations"`
		ApproxMode         bool        `json:"approx_mode"`
		LoaderStable       interface{} `json:"loader_stable"`
		LoaderTableName    interface{} `json:"loader_table_name"`
		MdCmd              interface{} `json:"md_cmd"`
		Query              string      `json:"query"`
		Sample             bool        `json:"sample"`
		ScriptLocation     interface{} `json:"script_location"`
	} `json:"command"`
	CommandType    string `json:"command_type"`
	Concurrency    uint   `json:"concurrency"`
	DependencyInfo struct {
		HiveTables []struct {
			Columns struct {
				StockExchange []string `json:"stock_exchange"`
				StockSymbol   []string `json:"stock_symbol"`
			} `json:"columns"`
			InitialInstance string `json:"initial_instance"`
			Interval        struct {
				Days string `json:"days"`
			} `json:"interval"`
			Name        string `json:"name"`
			TimeZone    string `json:"time_zone"`
			WindowEnd   string `json:"window_end"`
			WindowStart string `json:"window_start"`
		} `json:"hive_tables"`
	} `json:"dependency_info"`
	DigestTimeHour   int      `json:"digest_time_hour"`
	DigestTimeMinute int      `json:"digest_time_minute"`
	EmailList        string   `json:"email_list"`
	EndTime          string   `json:"end_time"`
	Frequency        int      `json:"frequency"`
	ID               int      `json:"id"`
	Incremental      struct{} `json:"incremental"`
	IsDigest         bool     `json:"is_digest"`
	Label            string   `json:"label"`
	Macros           []struct {
		FormattedDate string `json:"formatted_date"`
	} `json:"macros"`
	Name                 string `json:"name"`
	NextMaterializedTime string `json:"next_materialized_time"`
	NoCatchUp            bool   `json:"no_catch_up"`
	StartTime            string `json:"start_time"`
	Status               string `json:"status"`
	Template             string `json:"template"`
	TimeOut              int    `json:"time_out"`
	TimeUnit             string `json:"time_unit"`
	TimeZone             string `json:"time_zone"`
	UserID               uint   `json:"user_id"`
}
