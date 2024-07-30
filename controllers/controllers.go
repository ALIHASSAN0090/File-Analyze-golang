package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/utils"
)

// Stats godoc
// @Summary Analyze text file content
// @Description Analyze the text file to count vowels, capital letters, small letters, and spaces.
// @Accept  multipart/form-data
// @Produce json
// @Param routines formData int true "Number of routines (1 to 4)"
// @Success 200 {object} map[string]interface{} "Analysis results"
// @Failure 400 {object} map[string]string "Invalid input or number of routines out of range"
// @Failure 500 {object} map[string]string "Error opening file or inserting analysis results"
// @Router /stats [post]
func stats(g *gin.Context) {
	routinesStr := g.PostForm("routines")

	routines, err := strconv.Atoi(routinesStr)
	fmt.Println(routines)
	if err != nil {
		g.JSON(400, gin.H{"error": "You entered Aphabet. please enter a number between 1 and 4 "})
		return
	}

	if routines > 4 || routines < 1 {
		g.JSON(400, gin.H{"error": "Number of routines must be between 1 and 4"})
		return
	}

	filePath := "src/test.txt"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		g.JSON(500, gin.H{"error": "Error opening file"})
		return
	}

	text := string(fileContent)

	chunk := len(text) / routines
	var wg sync.WaitGroup
	results := make(chan map[string]int, routines)
	totalCounts := map[string]int{
		"vowels":  0,
		"capital": 0,
		"small":   0,
		"spaces":  0,
	}

	startTime := time.Now()

	for i := 0; i < routines; i++ {
		startId := i * chunk
		endId := startId + chunk
		if i == routines-1 {
			endId = len(text)
		}
		wg.Add(1)
		go utils.GetData(text[startId:endId], &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for count := range results {
		for key, value := range count {
			totalCounts[key] += value
		}
	}
	endTime := time.Now()
	processTime := endTime.Sub(startTime)
	milliSec := processTime.Milliseconds()
	//

	fmt.Println(totalCounts["vowels"], totalCounts["capital"], totalCounts["small"], totalCounts["spaces"])

	err = db.CreateUser(totalCounts["vowels"], totalCounts["capital"], totalCounts["small"], totalCounts["spaces"])
	if err != nil {
		fmt.Printf("Error inserting analysis results: %v\n", err)
		g.JSON(500, gin.H{"error": "Error inserting analysis results"})
		return
	}

	g.JSON(200, gin.H{

		"total_vowels":   totalCounts["vowels"],
		"total_capitals": totalCounts["capital"],
		"total_small":    totalCounts["small"],
		"total_spaces":   totalCounts["spaces"],
		"process_time":   milliSec,
	})
	fmt.Println("Analysis results inserted successfully")
}

// DisplayAll godoc
// @Summary Display all file analysis statistics
// @Description Retrieve all records of file analysis statistics from the database.
// @Produce  json
// @Success 200 {array} db.FileStats "List of file statistics"
// @Failure 500 {object} map[string]string "Error fetching or processing data"
// @Router / [get]
func DisplayAll(g *gin.Context) {

	rows, err := db.DbConn.Query("SELECT * FROM file_stats")
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	defer rows.Close()

	var results []db.FileStats

	for rows.Next() {
		var stat db.FileStats
		if err := rows.Scan(&stat.ID, &stat.Vowels, &stat.Capital, &stat.Small, &stat.Spaces); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}
		results = append(results, stat)
	}

	if err := rows.Err(); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error completing query"})
		return
	}

	g.JSON(http.StatusOK, results)
}
