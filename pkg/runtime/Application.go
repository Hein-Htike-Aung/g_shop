package runtime

import (
	"github.com/casbin/casbin/v2"
	"github.com/robfig/cron/v3"
	"sync"
)

type Application struct {
	casbins map[string]*casbin.SyncedEnforcer
	crontab map[string]*cron.Cron
	mux     sync.RWMutex
}

func (e *Application) SetCasbin(key string, enforcer *casbin.SyncedEnforcer) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.casbins[key] = enforcer
}

func (e *Application) GetCasbin() map[string]*casbin.SyncedEnforcer {
	return e.casbins
}

// GetCasbinKey get casbin by key
func (e *Application) GetCasbinKey(key string) *casbin.SyncedEnforcer {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e, ok := e.casbins["*"]; ok {
		return e
	}
	return e.casbins[key]
}

// SetCrontab set crontab for key
func (e *Application) SetCrontab(key string, crontab *cron.Cron) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.crontab[key] = crontab
}

// GetCrontab get all crontab entries from map
func (e *Application) GetCrontab() map[string]*cron.Cron {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.crontab
}

// GetCrontabKey get crontab by key
func (e *Application) GetCrontabKey(key string) *cron.Cron {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e, ok := e.crontab["*"]; ok {
		return e
	}
	return e.crontab[key]
}

// get default instance
func GetInstance() *Application {
	return &Application{
		casbins: make(map[string]*casbin.SyncedEnforcer),
		crontab: make(map[string]*cron.Cron),
	}
}
