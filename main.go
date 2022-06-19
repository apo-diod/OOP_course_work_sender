package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"

	"io/ioutil"
	"net/http"
)

const PORT = "6122"

func main() {
	r := gin.Default()
	r.POST("/use_module", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		id := req["id"].(string)
		log.Println(id)
		data := req["data"]
		toUsebts, _ := json.Marshal(data)
		toUse := string(toUsebts)
		log.Println(id, toUse)
		go useModule(id, toUse)
		ctx.JSON(http.StatusOK, nil)
	})

	r.POST("/add_module", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		module := req["module"].(string)
		settingsmap := req["settings"]
		settingsbts, _ := json.Marshal(settingsmap)
		settings := string(settingsbts)
		id := newModule(module, settings)
		if id == "0" {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.Run(":" + PORT)
}
