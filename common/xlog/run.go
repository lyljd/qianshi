package xlog

import (
	"encoding/json"
	"fmt"
	"os"
	"qianshi/common/email"
)

// StartCollection Run a goroutine to receive msg from channel
func StartCollection(serviceName string, config *Config) {
	sn = serviceName

	if config == nil {
		config = &Config{}
	}

	if config.Cache <= 0 {
		config.Cache = 10
	}

	if config.PrintLevel < 0 {
		config.PrintLevel = 0
	}

	if file, err := os.OpenFile("log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777); err != nil {
		panic(err)
	} else {
		f = file
	}

	ch = make(chan map[string]interface{}, config.Cache)

	if config.Report.Enabled {
		sc := config.Report.Send
		emailDialer = email.NewDialer(sc.Host, sc.Port, sc.Sender, sc.Username, sc.Password)
	}

	go func() {
		for {
			select {
			case m := <-ch:
				level := m["level"].(Level)
				if level < config.PrintLevel {
					break
				}
				m["level"] = levelConvert(level)
				write(m)
				if emailDialer != nil && level >= config.Report.ReportLevel {
					report(m, config.Report.Receive)
				}
			case <-done:
				goto Done
			}
		}
	Done:
	}()
}

func write(m map[string]any) {
	jm, _ := json.Marshal(m)
	jms := string(jm)
	if _, err := f.WriteString(jms + "\n"); err != nil {
		fmt.Printf("log \"%s\" write fail! err: %s\n", jms, err)
	}
}

func report(m map[string]any, rc ReceiveConfig) {
	jm, _ := json.MarshalIndent(m, "", "    ")
	value := fmt.Sprintf(rc.ContentTmpl, string(jm))
	var content *email.Content
	if rc.ContentType == "html" {
		content = email.Html(value)
	} else {
		content = email.Text(value)
	}
	if err := emailDialer.SendToMany(rc.To, rc.Subject, content); err != nil {
		fmt.Printf("log report fail! err: %s\n", err)
	}
}

func StopCollection() {
	done <- struct{}{}
	_ = f.Close()
	close(ch)
}
