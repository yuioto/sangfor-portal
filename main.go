package main

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	configFileName = "sangfor-portal"
)

func main() {
	// 设置zerolog日志格式
	// set log style, look like: 2006-01-02T15:04:05Z07:00 INFO Msg Str
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// 读取配置文件
	const configFileName = "sangfor-portal"

	err := readConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}

	// check_url := "https://ping.archlinux.org/nm-check.txt"
	url := "http://captive.apple.com/hotspot-detect.html"
	//url := viper.GetString("url")
	//account := viper.GetString("account")
	//password := viper.GetString("password")
	checkCaptivePortalAsync(url, login)

	log.Info().Msg("アトリは、高性能ですから!")

	// 在这里可以继续执行其他任务，不会被阻塞
	// 这里我简单地让程序等待一段时间来模拟其他任务
	for {
		time.Sleep(1000 * time.Second)
	}
}

func checkCaptivePortalAsync(url string, loginFunc func()) {
	go func() {
		for {
			if detectCaptivePortal(url) {
				log.Info().Msg("Detected Captive Portal. Initiating login process")
				loginFunc()
				break
			}
			time.Sleep(5 * time.Second) // 每隔5秒检测一次
		}
	}()
}

func detectCaptivePortal(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		log.Warn().Err(err).Msg("Error can't connect to network")
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true
	} else {
		log.Warn().Err(err).Msg("Error while detecting Captive Portal")
		return false
	}

}

func login() {
	log.Info().Msg("Login function called")
	// 在这里执行登录操作
	return
}

func readConfig() error {
	// 设置配置文件搜索路径
	viper.AddConfigPath("$HOME/.config") // 设置.config目录
	viper.AddConfigPath("/etc/")         // 设置etc目录
	viper.AddConfigPath(".")             // 设置当前目录
	viper.SetConfigName(configFileName)  // 设置配置文件名称
	viper.SetConfigType("toml")          // 设置配置文件类型为TOML

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Err(err).Msg("Config file not found")
			return err
		}
		log.Warn().Err(err).Msg("Failed to read config file")
		return err
	}

	log.Info().Msg("Config file loaded successfully")
	return nil
}
