package voice //Этап 4. Сбор данных о системе Voice Call

import (
	"bufio"
	"fmt"
	"os"
	"skillbox/go/basic/diplom2/pkg/service/bv"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_call_time"`
}

func voiceCallData(uri string) []VoiceCallData {
	result := make([]VoiceCallData, 0)

	file, err := os.Open(uri)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return []VoiceCallData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, ";")
		if len(data) != 8 {
			continue
		}

		if !bv.IsCountry(data[0]) || !bv.IsBandwidth(data[1]) || !bv.IsResponseTime(data[2]) || !bv.IsProviderVoiceCall(data[3]) {
			continue
		}

		ConnectionStability, err := strconv.ParseFloat(data[4], 32)
		if err != nil {
			continue
		}

		TTFB, err := strconv.Atoi(data[5])
		if err != nil {
			continue
		}

		VoicePurity, err := strconv.Atoi(data[6])
		if err != nil {
			continue
		}

		MedianOfCallsTime, err := strconv.Atoi(data[7])
		if err != nil {
			continue
		}

		elem := VoiceCallData{
			Country:             data[0],
			Bandwidth:           data[1],
			ResponseTime:        data[2],
			Provider:            data[3],
			ConnectionStability: float32(ConnectionStability),
			TTFB:                TTFB,
			VoicePurity:         VoicePurity,
			MedianOfCallsTime:   MedianOfCallsTime,
		}

		result = append(result, elem)
	}

	return result
}

/*country := fields[0]
		bandwidth := fields[1]
		responseTime := fields[2]
		provider := fields[3]

		if !isValidProvider(provider) {
			continue
		}

		connectionStability, err := strconv.ParseFloat(fields[4], 32)
		if err != nil {
			continue
		}

		ttfb, err := strconv.Atoi(fields[5])
		if err != nil {
			continue
		}

		voicePurity, err := strconv.Atoi(fields[6])
		if err != nil {
			continue
		}

		medianOfCallsTime, err := strconv.Atoi(fields[7])
		if err != nil {
			continue
		}

		result = append(result, VoiceCallData{
			Country:             country,
			Bandwidth:           bandwidth,
			ResponseTime:        responseTime,
			Provider:            provider,
			ConnectionStability: float32(connectionStability),
			TTFB:                ttfb,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   medianOfCallsTime,
		})
	}

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		log.Println("Ошибка при преобразовании данных в JSON:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}*/
