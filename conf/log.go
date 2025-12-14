package conf

type LogConfig struct {
	Level       string `json:"Level"`
	Output      string `json:"Output"`
	DnsLog      bool   `json:"DnsLog"`      // 是否启用 DNS 查询日志
	MaskAddress string `json:"MaskAddress"` // IP 地址遮罩: quarter, half, full
}
