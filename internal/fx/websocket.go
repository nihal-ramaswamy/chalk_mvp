package fx_utils

import (
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
)

var WebsocketModule = fx.Module(
	"WebsocketModule",
	fx.Provide(NewWebsocketUpgrader),
)

func NewWebsocketUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{}
}
