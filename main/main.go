package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/songtianyi/wechat-go/plugins/wxweb/gifer"
	"github.com/songtianyi/wechat-go/plugins/wxweb/replier"
	"github.com/songtianyi/wechat-go/plugins/wxweb/switcher"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	faceplusplus.Register(session)
	replier.Register(session)
	switcher.Register(session)
	gifer.Register(session)

	session.HandlerRegister.DisableByName("faceplusplus")

	if err := session.LoginAndServe(false); err != nil {
		logs.Error("session exit, %s", err)
	}
}
