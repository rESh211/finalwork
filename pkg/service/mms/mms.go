package mms //Этап 3. Сбор данных о системе MMS

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"skillbox/go/basic/diplom2/pkg/service/bv"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func mmsData(uri string) []MMSData {
	result := make([]MMSData, 0)

	file, err := http.Get(uri)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return []MMSData{}
	}
	defer file.Body.Close()

	if file.StatusCode != http.StatusOK {
		return []MMSData{}
	}

	body, err := ioutil.ReadAll(file.Body)
	if err != nil {
		fmt.Println(err.Error())
		return []MMSData{}
	}

	var mms []MMSData

	if err := json.Unmarshal(body, &mms); err != nil {
		return []MMSData{}
	}
	for _, elem := range mms {
		if !bv.IsCountry(elem.Country) || !bv.IsBandwidth(elem.Bandwidth) || !bv.IsResponseTime(elem.ResponseTime) || !bv.IsProviderSMSandMMS(elem.Provider) {
			continue
		}

		country := bv.GetCountryForCode(elem.Country)
		elem.Country = country
		result = append(result, elem)
	}

	return result
}

/*result := make([]MMSData, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			continue
		}
		var mms MMSData
		mms.Country = fields[0]
		mms.Provider = fields[1]
		mms.Bandwidth = fields[2]
		mms.ResponseTime = fields[3]

		result = append(result, mms)
	}
	return result, nil
}
func mmsData(uri string) ([]MMSData, error) {
	file, err := os.Open(uri)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}
	defer file.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Ошибка: Код ответа не равен 200")
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении тела ответа:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	var mmsData []MMSData
	err = json.Unmarshal(body, &mmsData)
	if err != nil {
		log.Println("Ошибка при разборе JSON в структуру MMSData:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	filteredMMSData := make([]MMSData, 0)
	for _, data := range mmsData {
		if isValidCountry(data.Country) && isValidProvider(data.Provider) {
			filteredMMSData = append(filteredMMSData, data)
		}
	}

	jsonResponse, err := json.Marshal(filteredMMSData)
	if err != nil {
		log.Println("Ошибка при преобразовании данных в JSON:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}*/
