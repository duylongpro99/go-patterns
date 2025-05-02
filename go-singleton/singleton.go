package singleton

// # Concurrency-safe Version #

import "sync"

type ConfigManager struct {
	AppConfig map[string]string
}

var (
	once     sync.Once
	instance *ConfigManager
)

func GetInstance() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			AppConfig: map[string]string{
				"env":     "dev",
				"port":    "3000",
				"timeout": "60s",
			},
		}
	})
	return instance
}

func TestConcurrency() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			println("\tRouting:", i)
			defer wg.Done()
			config := GetInstance()
			println("\tEnvironment:", config.AppConfig["env"], "Port:", config.AppConfig["port"], "Timeout:", config.AppConfig["timeout"])
		}()
	}

	wg.Wait()
}
