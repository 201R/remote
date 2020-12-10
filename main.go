package main

import (
	"encoding/json"

	"github.com/201RichK/remote/remote"
	"github.com/sirupsen/logrus"
)

func main() {
	client := remote.NewRemote(remote.Config{})

	res, err := client.POST(remote.Options{
		URL: "https://api.bizao.com/mobilemoney/v1",
		Header: map[string]string{
			"Authorization": "Bearer 6487f028-0335-3839-915e-8dc08392a59f",
			"country-code":  "ci",
			"channel":       "web",
			"lang":          "fr",
			"content-type":  "application/json",
			"mno-name":      "orange",
		},
		Body: map[string]interface{}{
			"currency":    "XOF",
			"order_id":    "AmosPascal_54",
			"amount":      5000,
			"reference":   "KouaKou Amos Pascal",
			"user_msisdn": "22578674311",
			"state":       "param1%3Dvalue1%26param2%3Dvalue2",
			"notif_url":   "https://residences-h2o.com/api/v1/reservation/checkout",
			"return_url":  "https://residences-h2o.com/api/v1/reservation/checkout",
			"cancel_url":  "https://residences-h2o.com/api/v1/reservation/checkout",
		},
	})

	if err != nil {
		logrus.Error("err : ", err)
	}
	var data interface{}
	json.NewDecoder(res.Body).Decode(&data)
	logrus.Info(data)
}
