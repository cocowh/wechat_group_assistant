package main

import (
	"encoding/json"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
)

func main() {
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	//faceplusplus.Register(session)
	//replier.Register(session)
	//switcher.Register(session)
	//gifer.Register(session)
	//session.HandlerRegister.DisableByName("faceplusplus")
	Register(session)

	if err := session.LoginAndServe(false); err != nil {
		logs.Error("session exit, %s", err)
	}
}

func Register(session *wxweb.Session)  {
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(demo), "text-replier")
	if err := session.HandlerRegister.Add(wxweb.MSG_IMG, wxweb.Handler(demo), "img-replier"); err != nil {
		logs.Error(err)
	}

	if err := session.HandlerRegister.EnableByName("text-replier"); err != nil {
		logs.Error(err)
	}

	if err := session.HandlerRegister.EnableByName("img-replier"); err != nil {
		logs.Error(err)
	}
}

func demo(session *wxweb.Session, msg *wxweb.ReceivedMessage)  {
	contact := session.Cm.GetContactByUserName(msg.FromUserName)
	b, err := json.Marshal(msg)
	if err != nil {
		session.SendText( "error", session.Bot.UserName, msg.FromUserName)
	}
	if msg.MsgType == wxweb.MSG_IMG {
		session.SendText( string(b), session.Bot.UserName, msg.FromUserName)
	}

	if !msg.IsGroup && contact.PYQuanPin != "tingjianliangshan" {
		session.SendText(string(b), session.Bot.UserName, msg.FromUserName)
	}

	if msg.IsGroup && msg.Who != session.Bot.UserName {
		session.SendText(msg.FromUserName, session.Bot.UserName, msg.FromUserName)
	}
}