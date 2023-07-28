package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var (
	Cfg *Config
)

// Config 对应配置文件结构
type Config struct {
	//Listen         string                       `toml:"listen"`
	//GrpcListen     string                       `toml:"grpc_listen"`
	//RoomCardAPIURL string                       `toml:"roomcardapi_url"`
	DBServers map[string]DBServer `toml:"dbservers"`
	//RedisServers   map[string]RedisServer       `toml:"redisservers"`
	//Grpc           map[string]Connect           `toml:"grpc"`
	//MetricConfig   *metricutil.MetricUtilConfig `toml:"metric"`
	//TrieFilePath   string                       `toml:"trie_file_path"`
	//UserInfoURL    string                       `toml:"user_info_url"`
	//WorkerInfoURL  string                       `toml:"worker_info_url"`
	//AgentSrvURL    string                       `toml:"agent_srv_url"`
	//MailAPI        string                       `toml:"mail_api"`
}

// DBServer 表示DB服务器配置
type DBServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type Connect struct {
	Addrs  []string `toml:"addrs"`
	Weight []int    `toml:"weight"`
}

// ConnectString 表示连接数据库的字符串
func (m DBServer) ConnectString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.Host, m.Port, m.User, m.Password, m.DBName)
}

// DBServerConf 获取数据库配置
func (c Config) DBServerConf(key string) (DBServer, bool) {
	s, ok := c.DBServers[key]
	return s, ok
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	if _, err := toml.DecodeFile(tomlfile, &Cfg); err != nil {
		return Cfg, err
	}
	return Cfg, nil
}
