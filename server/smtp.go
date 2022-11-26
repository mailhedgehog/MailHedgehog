package server

import (
	"fmt"
	"github.com/mailpiggy/MailPiggy/logger"
	"net"
)

func smtpListen(context *Context, exitCh chan int) {
	logManager().Debug(fmt.Sprintf("SMTP Binding to address %s", context.smtpBindAddr()))
	listener, err := net.Listen("tcp", context.smtpBindAddr())
	logger.PanicIfError(err)
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			logManager().Error(fmt.Sprintf("[SMTP] Error accepting connection: %s", err.Error()))
			continue
		}

		logManager().Warning("TODO: SMTP add Monkey")

		go smtpSession(connection, context)
	}
}

func smtpSession(connection net.Conn, context *Context) {
	defer connection.Close()
	logManager().Debug("Start smtp session")

	logManager().Debug("End smtp session")
}
