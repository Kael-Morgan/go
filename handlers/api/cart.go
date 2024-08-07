package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-beyond/server"
	"go-beyond/services"
	"net/http"
	"strconv"
	"time"

	"nhooyr.io/websocket"
)

type CartItem map[string]ItemInfo

type ItemInfo struct {
	Name      string `json:"name"`
	IsChecked bool   `json:"isChecked"`
}

type Command map[string]CartItem

// func HandleUpdateCartItem(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithCancel(r.Context())
// 	defer cancel()

// 	cartName := r.PathValue("name")

// 	if cartName == "" {
// 		return
// 	}
// 	itemId := r.URL.Query().Get("id")

// 	if itemId == "" {
// 		return
// 	}

// }

func HandleUpdateCartItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	cartName := r.PathValue("name")

	var item ItemInfo
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		fmt.Println(err)
		return
	}

	operation := "modify"

	itemId := r.URL.Query().Get("id")
	if itemId == "" {
		operation = "add"
		itemId = getItemId()
	}

	redisHSETValue, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	// for services.GetRedisClient().HGet(ctx, cartName, itemId) != nil {
	// 	itemId = getItemId()
	// }

	services.GetRedisClient().HSet(r.Context(), cartName, itemId, redisHSETValue)

	jsonItem, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	command := Command{operation: {itemId: item}}

	jsonRes, err := json.Marshal(command)

	if err != nil {
		fmt.Println(err)
		return
	}

	sendCommand(ctx, cartName, jsonRes)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonItem)

}

func HandleDeleteCartItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	cartName := r.PathValue("name")
	if cartName == "" {
		return
	}
	itemId := r.URL.Query().Get("id")

	if itemId == "" {
		w.Write([]byte("Error: no id in Parameters"))
	}

	services.GetRedisClient().HDel(ctx, cartName, itemId)

	res, err := json.Marshal("Deleted" + itemId)
	if err != nil {
		fmt.Println(err)
		return
	}

	command := Command{"remove": {itemId: {}}}

	jsonCommand, err := json.Marshal(command)

	if err != nil {
		fmt.Println(err)
		return
	}

	sendCommand(ctx, cartName, jsonCommand)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func HandleGetCartItems(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	cartName := r.PathValue("name")
	data := services.GetRedisClient().HGetAll(ctx, cartName)

	res := CartItem{}

	for key, item := range data.Val() {
		unjsoned := ItemInfo{}
		json.Unmarshal([]byte(item), &unjsoned)
		res[key] = unjsoned
	}

	jsoned, err := json.Marshal(res)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(jsoned)
}

func sendCommand(ctx context.Context, cartName string, content []byte) {
	for ws, client := range server.GetClients() {
		if client.CartName == cartName {
			if err := ws.Write(ctx, websocket.MessageText, content); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getItemId() string {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	return strconv.FormatInt(now, 36)
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
