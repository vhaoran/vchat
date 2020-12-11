package yconfig

// for elasticsearch
type ESConfig struct {
	Url []string `json:"url,omitempty"`
	Sniff bool `json:"sniff,omitempty"`
}

