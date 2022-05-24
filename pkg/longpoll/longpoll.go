package longpoll

import "github.com/jcuga/golongpoll"

var longpollManager *golongpoll.LongpollManager

func Init() (err error) {
	manager, err := golongpoll.StartLongpoll(golongpoll.Options{})
	if err != nil {
		return err
	}
	longpollManager = manager
	return
}

func Manger() *golongpoll.LongpollManager {
	return longpollManager
}
