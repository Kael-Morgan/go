package handlers

import (
	"encoding/json"
	"fmt"
	"go-beyond/services"
	"net/http"
)

func HandleUpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartName := r.PathValue("name")
	itemId := r.URL.Query().Get("id")
	if itemId == "" {
		fmt.Println("Empty param", cartName)
	}
	var item CartItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range item {
		redisHSETValue, err := json.Marshal(value)
		if err != nil {
			fmt.Println(err)
		}
		services.GetRedisClient().HSet(r.Context(), cartName, key, redisHSETValue)
	}

	fmt.Println(cartName, item)
	w.Write([]byte("received"))

}

func HandleDeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartName := r.URL.Query().Get("name")
	itemId := r.URL.Query().Get("id")
	if itemId == "" {
		w.Write([]byte("Error: no id in Parameters"))
	}

	services.GetRedisClient().HDel(r.Context(), cartName, itemId)

	w.Write([]byte("Deleted " + itemId))

}

func HandleGetCartItems(w http.ResponseWriter, r *http.Request) {
	cartName := r.PathValue("name")
	data := services.GetRedisClient().HGetAll(r.Context(), cartName)

	res := CartItem{}

	for key, item := range data.Val() {
		unjsoned := ItemInfo{}
		json.Unmarshal([]byte(item), &unjsoned)
		res[key] = unjsoned
	}

	items, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte("read error" + err.Error()))
	}

	w.Write(items)
}

type CartItem map[string]ItemInfo

type ItemInfo struct {
	Name      string `json:"name"`
	IsChecked bool   `json:"isChecked"`
}

// redisClient := services.GetRedisClient()

// redisClient.Del(ctx, "initCart")

// text := redisClient.HGetAll(ctx, "initCart")

// testData := CartItem{
// 	"1":        {Name: "halberk", IsChecked: false},
// 	"2":        {Name: "M249", IsChecked: true},
// 	"lz9vz8ia": {Name: "M4A1-S", IsChecked: false},
// }

// for key, item := range testData {
// jsoned, _ := json.Marshal(item)
// var unjsoned_data *ItemInfo
// json.Unmarshal(jsoned, &unjsoned_data)
// redisClient.HSet(ctx, "initCart", key, jsoned)
// fmt.Println(key, jsoned, unjsoned_data)
// }

// data := redisClient.HGetAll(ctx, "initCart")

// for _, item := range data.Val() {
// 	unjsoned := &ItemInfo{}
// 	json.Unmarshal([]byte(item), unjsoned)
// 	fmt.Println(item, unjsoned.Name, unjsoned.IsChecked)
// }
// fmt.Println(data.Val())

// jsoned, _ := json.Marshal(data)
// fmt.Println(jsoned)
