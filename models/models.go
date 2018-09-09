package models

type Sigure struct {
	Sig     string `json:"sig"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"admin"`
}

type ApiToken struct {
	Sig  string `json:"sig"`
	Name string `json:"name"`
}

type ReqHistory struct {
	Step      int      `json:"step"`
	StartTime int64    `json:"start_time"`
	HostNames []string `json:"hostnames"`
	EndTime   int64    `json:"end_time"`
	Counters  []string `json:"counters"`
	ConsolFun string   `json:"consol_fun"`
}

type Resp struct {
	Endpoint string `json:"endpoint"`
	Counter  string `json:"counter"`
	Dstype   string `json:"dstype"`
	Step     int    `json: "step"`
	Values   []V    `json:"Values"`
}

type V struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type DashBordScreen struct {
	Id   int64  `json:"id"`
	Pid  int64  `json:"pid"`
	Name string `json:"name"`
}

type GraphInfo struct {
	Counters   []string `json:"counters"`
	Endpoints  []string `json:"endpoints"`
	FalconTags string   `json:"falcon_tags"`
	GraphId    int64    `json:"graph_id"`
	GraphType  string   `json:"graph_type"`
	Method     string   `json:"method"`
	Position   int64    `json:"position"`
	ScreenId   int64    `json:"screen_id"`
	Timespan   int64    `json:"timespan"`
	Title      string   `json:"title"`
}

type CheckTable struct {
	ParScreen   string  `json:"parScreen"`
	ChildScreen string  `json:"childScreen"`
	Metric      string  `json:"metric"`
	ComMode     string  `json:"comMode"`
	JudgeSymbol string  `json:"judgeSymbol"`
	Threshold   float64 `json:"threshold"`
	SpanTime    int     `json:"spanTime"`
	IsAbnormal  bool    `json:"isAbnormal"`
	Desc        string  `json:"desc"`
}

type Cfg struct {
	Debug            bool   `json:"debug"`
	ImportExcelPath  string `json:"importExcelPath"`
	ExportExecelPath string `json:"exportExecelPath"`
	ApiAddr          string `json:"apiAddr"`
	MailServer       string `json:"mailServer"`
	MailServerPort   int    `json:"mailServerPort"`
	User             string `json:"user"`
	Passwd           string `json:"passwd"`
	From             string `json:"from"`
	Tos              string `json:"tos"`
}
