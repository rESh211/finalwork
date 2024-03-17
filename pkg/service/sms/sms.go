package sms //Этап 2. Сбор данных о системе SMS

import (
	"bufio"
	"fmt"
	"os"
	"skillbox/go/basic/diplom2/pkg/service/bv"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func smsData(uri string) []SMSData {
	result := make([]SMSData, 0)

	file, err := os.Open(uri)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return []SMSData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ";")
		if len(data) != 4 {
			continue
		}
		if !bv.IsCountry(data[0]) || !bv.IsBandwidth(data[1]) || !bv.IsResponseTime(data[2]) || !bv.IsProviderSMSandMMS(data[3]) {
			continue
		}
		sms := SMSData{
			Country:      data[0],
			Bandwidth:    data[1],
			ResponseTime: data[2],
			Provider:     data[3],
		}

		result = append(result, sms)
	}
	return []SMSData{}
}
