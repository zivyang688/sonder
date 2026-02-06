package app

import (
	"fmt"
	"log"
	"sync"
)

var (
	once     sync.Once
	instance *Apps
)

type BaseApp struct {
	Id      string
	Name    string
	Version string
}

type Apps struct {
	apps []IApp
	mu   sync.RWMutex
}

type IApp interface {
	GetInfo() string
	GetName() string
	GetId() string
	GetVersion() string
	// Init() error
	Start() error
	Stop() error
}

func (b *BaseApp) GetId() string {
	return b.Id
}
func (b *BaseApp) GetName() string {
	return b.Name
}

func (b *BaseApp) GetVersion() string {
	return b.Version
}

func (b *BaseApp) GetInfo() string {
	return fmt.Sprintf("app_id: %s  app_name: %s, app_version: %s", b.Id, b.Name, b.Version)
}

func NewApps() *Apps {
	once.Do(func() {
		instance = &Apps{
			apps: make([]IApp, 64),
		}
	})
	return instance
}

func (a *Apps) RegisterApp(app IApp) {
	if app == nil {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.apps = append(a.apps, app)
}

// func (a *Apps) Inits() {
// 	a.mu.RLock()
// 	defer a.mu.RUnlock()
// 	for _, app := range a.apps {
// 		err := app.Init()
// 		if err != nil {
// 			log.Fatalf("[APP INIT ERROR] app_id: %s  app_name: %s, app_version: %s, err: %v",
// 				app.GetId(), app.GetName(), app.GetVersion(), err)
// 			continue
// 		}
// 	}
// }

func (a *Apps) Starts() {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, app := range a.apps {
		err := app.Start()
		if err != nil {
			log.Fatalf("[APP START ERROR] app_id: %s  app_name: %s, app_version: %s, err: %v",
				app.GetId(), app.GetName(), app.GetVersion(), err)
			continue
		}
	}
}

func (a *Apps) Statics() {

}

func (a *Apps) Stops() {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, app := range a.apps {
		err := app.Stop()
		if err != nil {
			log.Fatalf("[APP STOP ERROR] app_id: %s  app_name: %s, app_version: %s, err: %v",
				app.GetId(), app.GetName(), app.GetVersion(), err)
			continue
		}
	}
}
