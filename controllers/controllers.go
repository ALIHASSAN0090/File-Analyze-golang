package controllers

import (
	"io"
	"net/http"
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
// @Param file formData file true "Text file to analyze"
// @Success 200 {object} map[string]interface{} "Analysis results"
// @Failure 400 {object} map[string]string "Invalid input or number of routines out of range"
// @Failure 500 {object} map[string]string "Error opening file or inserting analysis results"
// @Router /stats [post]
func stats(c *gin.Context) {
	routinesStr := c.PostForm("routines")
	routines, err := strconv.Atoi(routinesStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Please enter a number between 1 and 4"})
		return
	}

	if routines < 1 || routines > 4 {
		c.JSON(400, gin.H{"error": "Number of routines must be between 1 and 4"})
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error opening file"})
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error reading file"})
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

	err = db.CreateUser(totalCounts["vowels"], totalCounts["capital"], totalCounts["small"], totalCounts["spaces"])
	if err != nil {
		c.JSON(500, gin.H{"error": "Error inserting analysis results"})
		return
	}

	c.JSON(200, gin.H{
		"total_vowels":   totalCounts["vowels"],
		"total_capitals": totalCounts["capital"],
		"total_small":    totalCounts["small"],
		"total_spaces":   totalCounts["spaces"],
		"process_time":   milliSec,
	})
}

// DisplayAll godoc
// @Summary Display all file analysis statistics
// @Description Retrieve all records of file analysis statistics from the database.
// @Produce  json
// @Success 200 {array} db.FileStats "List of file statistics"
// @Failure 500 {object} map[string]string "Error fetching or processing data"
// @Router / [get]
func DisplayAll(c *gin.Context) {
	rows, err := db.DbConn.Query("SELECT * FROM file_stats")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	defer rows.Close()

	var results []db.FileStats
	for rows.Next() {
		var stat db.FileStats
		if err := rows.Scan(&stat.ID, &stat.Vowels, &stat.Capital, &stat.Small, &stat.Spaces); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
			return
		}
		results = append(results, stat)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error completing query"})
		return
	}

	c.JSON(http.StatusOK, results)
}
