package jwtx

import "time"

type Manager struct {
	mapping map[string]*JWT   // app -> jwt config
}

func (mgr *Manager) WithApp(app string) *JWT {
	return mgr.mapping[app]
}

func NewManager(configList []*Config) *Manager {
	mgr := &Manager{mapping: make(map[string]*JWT, len(configList))}
	for _, c := range configList {
		mgr.mapping[c.Name] = &JWT{
			AccessSecret: c.AccessSecret,
			AccessExpire: time.Duration(c.AccessExpire) * time.Second,
		}
	}
	return mgr
}
