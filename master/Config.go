package master

import (
	"io/ioutil"
	"encoding/json"
	"sync"
)
var lock sync.Mutex
// 程序配置
type Config struct {
	ApiPort int	`json:"apiPort"`
	ApiReadTimeout int	`json:"apiReadTimeout"`
	ApiWriteTimeout int	`json:"apiWriteTimeout"`
	EtcdEndpoints []string `json:"etcdEndpoints"`
	EtcdDialTimeout int `json:"etcdDialTimeout"`
}

var (
	// 单例
	G_config *Config
)

// 加载配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf Config
	)

	// 1,读配置
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	//单例
	if G_config == nil {
		lock.Lock()
		if(G_config == nil){
			G_config = &conf
		}
		lock.Unlock()
	}

	return
}