package conf

import (
	"fmt"
	"io"
	"os"

	"github.com/GoAsyncFunc/NeXT-Server/common/json5"

	"encoding/json/v2"
)

type Conf struct {
	LogConfig  LogConfig    `json:"Log"`
	XrayConfig *XrayConfig  `json:"Xray"`
	NodeConfig []NodeConfig `json:"Nodes"`
}

func New() *Conf {
	return &Conf{
		LogConfig: LogConfig{
			Level:  "info",
			Output: "",
		},
		XrayConfig: NewXrayConfig(),
	}
}

func (p *Conf) LoadFromPath(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open config file error: %s", err)
	}
	defer f.Close()

	reader := json5.NewTrimNodeReader(f)
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("read config file error: %s", err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		return fmt.Errorf("unmarshal config error: %s", err)
	}

	return nil
}
