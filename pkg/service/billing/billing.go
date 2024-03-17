package billing //Этап 6. Сбор данных о системе Billing*

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func billingData(uri string) []BillingData {
	file, err := os.Open(uri)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return []BillingData{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()

		if len(line) != 6 {
			return []BillingData{}
		}
		var bytes []byte
		runes := []rune(line)
		for _, symbol := range runes {
			switch string(symbol) {
			case "0":
				bytes = append(bytes, 0)
			case "1":
				bytes = append(bytes, 1)
			default:
				return []BillingData{}
			}
		}
		bytes[0], bytes[5] = bytes[5], bytes[0]
		bytes[1], bytes[4] = bytes[4], bytes[1]
		bytes[2], bytes[3] = bytes[3], bytes[2]

		var digit uint8
		for i := 0; i < 6; i++ {
			digit = digit + bytes[i]*uint8(math.Pow(2, float64(i)))
		}

		CreateCustomer := digit&1 == 1
		Purchase := digit>>1&1 == 1
		Payout := digit>>2&1 == 1
		Recurring := digit>>3&1 == 1
		FraudControl := digit>>4&1 == 1
		CheckoutPage := digit>>5&1 == 1

		elems := BillingData{
			CreateCustomer: CreateCustomer,
			Purchase:       Purchase,
			Payout:         Payout,
			Recurring:      Recurring,
			FraudControl:   FraudControl,
			CheckoutPage:   CheckoutPage,
		}

		return []BillingData{elems}
	}

	return []BillingData{}
}

/*for _, line := range lines {
		bits := make([]int, len(line))
		for i, char := range line {
			if char == '1' {
				bits[i] = 1
			}
		}

		sum := countNumber(bits)

		billingData := BillingData{
			CreateCustomer: (sum & (1 << 0)) != 0,
			Purchase:       (sum & (1 << 1)) != 0,
			Payout:         (sum & (1 << 2)) != 0,
			Recurring:      (sum & (1 << 3)) != 0,
			FraudControl:   (sum & (1 << 4)) != 0,
			CheckoutPage:   (sum & (1 << 5)) != 0,
		}

		results = append(results, billingData)
	}

	jsonResponse, err := json.Marshal(results)
	if err != nil {
		log.Println("Ошибка при преобразовании данных в JSON:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}*/
