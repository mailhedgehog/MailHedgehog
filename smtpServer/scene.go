package smtpServer

type Scene interface {
	Init(receivedLine string, protocol *Protocol) *Reply
	HandleLine(receivedLine string) *Reply
}
