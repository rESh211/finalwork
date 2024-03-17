package email //Этап 5. Сбор данных о системе Email

import (
	"bufio"
	"fmt"
	"os"
	"skillbox/go/basic/diplom2/pkg/service/bv"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func EmailReport(uri string, code string, min bool) []EmailData {

	result := make([]EmailData, 0)
	file, err := os.Open(uri)

	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return []EmailData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, ";")
		if len(data) != 3 {
			continue
		}

		Country := data[0]
		if !bv.IsCountry(Country) {
			continue
		}
		Provider := data[1]
		if !bv.IsProviderEmail(Provider) {
			continue
		}
		DeliveryTime, err := strconv.Atoi(data[2])
		if err != nil {
			continue
		}
		elem := EmailData{
			Country:      Country,
			Provider:     Provider,
			DeliveryTime: DeliveryTime,
		}
		if elem.Country == code {
			result = append(result, elem)
		}
	}

	if min {
		sort.SliceStable(result, func(i, j int) bool {
			return result[i].DeliveryTime < result[j].DeliveryTime
		})
	} else {
		sort.SliceStable(result, func(i, j int) bool {
			return result[i].DeliveryTime > result[j].DeliveryTime
		})
	}

	if len(result) < 3 {
		return result
	}

	return result[:3]
}

/*func EmailReport(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("../simulator/email.data")
	if err != nil {
		log.Println("Ошибка чтения файла:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	lines := strings.Split(string(data), "\n")
	var result []EmailData

	for _, line := range lines {
		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			continue
		}

		country := fields[0]
		provider := fields[1]
		deliveryTime, err := strconv.Atoi(fields[2])
		if err != nil || !isValidCountry(country) || !isValidProvider(provider) {
			continue
		}

		emailData := EmailData{
			Country:      country,
			Provider:     provider,
			DeliveryTime: deliveryTime,
		}

		result = append(result, emailData)
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
