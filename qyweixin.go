package DailyPrice

type PushMsgText struct {
	Content string `json:"content"`
}

type PushMsg struct {
	MsgType string      `json:"msgtype"`
	Text    PushMsgText `json:"text"`
}
