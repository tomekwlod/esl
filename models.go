package esl

type Log struct {
	Project string      `json:"project"`
	Domain  string      `json:"domain"`
	Command string      `json:"command"`
	Flag    string      `json:"flag"`
	Data    interface{} `json:"data"`
	Created string      `json:"created"`
}

type Client struct {
	Url      string
	Port     int
	User     string
	Password string
	Project  string
}

type JSONResp struct {
	Error  string `json:"error"`
	Index  string `json:"_index"`
	Type   string `json:"_type"`
	ID     string `json:"_id"`
	Result string `json:"result"`
}
