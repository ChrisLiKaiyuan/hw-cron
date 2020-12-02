package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/imroc/req"
	"github.com/wuhan005/hw-cron/cas"
	"github.com/wuhan005/hw-cron/fanya"
	log "unknwon.dev/clog/v2"
)

func init() {
	_ = log.NewConsole()
}

func main() {
	casauth, err := cas.NewSession(os.Getenv("HDU_NO"), os.Getenv("HDU_PASSWORD"))
	if err != nil {
		log.Fatal("Failed to login: %v", err)
	}

	fy := fanya.New()
	err = casauth.ServiceLogin(fy)
	if err != nil {
		log.Fatal("Failed to login to fanya: %v", err)
	}

	terms, err := fy.GetAllTerm()
	if err != nil {
		log.Fatal("Failed to get all term: %v", err)
	}

	// HACK: get the first and the third term.
	courses, err := fy.GetCourseList(terms[14])
	if err != nil {
		log.Fatal("Failed to get courses list: %v", err)
	}

	Acontent := ""

	for _, course := range courses {
		homeworks, err := fy.GetHomeworks(course)

		//log.Warn("%s", course.Name)

		if err != nil {
			log.Warn("Failed to get homework of %s: %v", course.Name, err)
			continue
		}
		for _, hw := range homeworks {
			if hw.Status == fanya.EXPIRED || hw.Status == fanya.FINISHED {
				continue
			}
			timeToEnd := time.Duration(hw.End.UnixNano() - time.Now().UnixNano())

			//title := fmt.Sprintf("%s 作业即将截止", course.Name)
			if timeToEnd < 0 {
				content := fmt.Sprintf("%s - %s 需要提交，无截止日期。", course.Name, hw.Title)
				//sendAlert(title, content)
				//Acontent = Acontent + fmt.Sprintf("%s\r\n\r\n", content)
				log.Trace(content)
			} else if timeToEnd < 1*time.Hour { // < 1
				content := fmt.Sprintf("%s - %s 还有不到 1 小时截止提交。【 %s 】", course.Name, hw.Title, hw.End.Format("2006-01-02 15:04:05"))
				//sendAlert(title, content)
				Acontent = Acontent + fmt.Sprintf("%s\r\n\r\n", content)
				log.Trace(content)
			} else if timeToEnd < 3*time.Hour { // < 3
				content := fmt.Sprintf("%s - %s 还有不到 3 小时截止提交。【 %s 】", course.Name, hw.Title, hw.End.Format("2006-01-02 15:04:05"))
				//sendAlert(title, content)
				Acontent = Acontent + fmt.Sprintf("%s\r\n\r\n", content)
				log.Trace(content)
			} else if timeToEnd < 12*time.Hour { // < 12
				content := fmt.Sprintf("%s - %s 还有不到 12 小时截止提交。【 %s 】", course.Name, hw.Title, hw.End.Format("2006-01-02 15:04:05"))
				//sendAlert(title, content)
				Acontent = Acontent + fmt.Sprintf("%s\r\n\r\n", content)
				log.Trace(content)
			} else if timeToEnd < 24*time.Hour { // < 24
				content := fmt.Sprintf("%s - %s 还有不到 24 小时截止提交。【 %s 】", course.Name, hw.Title, hw.End.Format("2006-01-02 15:04:05"))
				//sendAlert(title, content)
				Acontent = Acontent + fmt.Sprintf("%s\r\n\r\n", content)
				log.Trace(content)
			} else if timeToEnd < 48*time.Hour { // < 48
				content := fmt.Sprintf("%s - %s 还有不到 48 小时截止提交。【 %s 】", course.Name, hw.Title, hw.End.Format("2006-01-02 15:04:05"))
				//sendAlert(title, content)
				Acontent = Acontent + fmt.Sprintf("%s\r\n\r\n", content)
				log.Trace(content)
			} else if timeToEnd < 72*time.Hour { // < 72
				content := fmt.Sprintf("%s - %s 还有不到 72 小时截止提交。【 %s 】", course.Name, hw.Title, hw.End.Format("2006-01-02 15:04:05"))
				//sendAlert(title, content)
				Acontent = Acontent + fmt.Sprintf("%s\r\n\r\n", content)
				log.Trace(content)
			}
		}
	}

	sendAlert("作业通知", Acontent)

	log.Stop()
}

// sendAlert 发送推送，可以接入 Bark 等服务
func sendAlert(title, content string) {
	alertURL := os.Getenv("ALERT_URL")
	alertURL = strings.ReplaceAll(alertURL, "{{title}}", url.QueryEscape(title))
	alertURL = strings.ReplaceAll(alertURL, "{{content}}", url.QueryEscape(content))
	_, _ = req.Get(alertURL)
	log.Trace("ALERT_MESSAGE sent")
}
