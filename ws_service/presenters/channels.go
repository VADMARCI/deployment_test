package presenters

import (
	"sync"
)

type ConnectionStore struct {
	connections sync.Map
}

type ChannelStore struct {
	channels    sync.Map
	MainFactory *MainFactory
}
