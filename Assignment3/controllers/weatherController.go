package controllers

import (
	"assignment-3/database"
	"assignment-3/helpers"
	"assignment-3/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) {
	var result models.WeatherData
	var err error
	var water, wind int
	rows, err := database.DB.Query("SELECT water, wind FROM weather")

	for rows.Next() {
		err = rows.Scan(&water, &wind)
		result.Water = water
		result.Wind = wind
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err.Error())
		return
	}

	// getData, err := json.MarshalIndent(result, "", " ")

	// if err != nil {
	// 	log.Println("Error:", err)
	// }

	// waterStatus := helpers.GetWaterStatus(water)
	// windStatus := helpers.GetWindStatus(wind)

	// fmt.Printf("%s\nstatus water : %s\nstatus wind : %s\n", string(getData), waterStatus, windStatus)

	ctx.JSON(http.StatusOK, result)
}

func UpdateData(ctx *gin.Context) {
	var reqdata models.WeatherData

	if err := ctx.ShouldBindJSON(&reqdata); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	sqlStatement := `
	UPDATE weather SET water = $2, wind = $3, updated_at = $4 
	WHERE id = $1
	`
	_, err := database.DB.Exec(sqlStatement, 1, reqdata.Water, reqdata.Wind, time.Now())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Success update data"})
}

func RequestUpdateData() {
	client := &http.Client{}

	requestBody := models.WeatherData{
		Water: helpers.RandomNumber(1, 100),
		Wind:  helpers.RandomNumber(1, 100),
	}

	body, err := json.Marshal(requestBody)

	if err != nil {
		log.Println("Error:", err)
	}

	request, err := http.NewRequest("PUT", "http://localhost:8080/weather", bytes.NewBuffer(body))

	if err != nil {
		log.Println("Error:", err)
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		log.Println("Error:", err)
	}

	defer response.Body.Close()

	getData, err := json.MarshalIndent(requestBody, "", " ")

	if err != nil {
		log.Println("Error:", err)
	}

	waterStatus := helpers.GetWaterStatus(requestBody.Water)
	windStatus := helpers.GetWindStatus(requestBody.Wind)

	fmt.Printf("%s\nstatus water : %s\nstatus wind : %s\n", string(getData), waterStatus, windStatus)

}
