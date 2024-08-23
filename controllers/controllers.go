package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/middleware"
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
	// Get number of routines from form data
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

	// Retrieve and read the file from form data
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

	// Process the text in parallel
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

	// Aggregate the results
	for count := range results {
		for key, value := range count {
			totalCounts[key] += value
		}
	}

	endTime := time.Now()
	processTime := endTime.Sub(startTime)
	milliSec := processTime.Milliseconds()

	// Store results in the database
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
	fmt.Println("Total counts:", totalCounts)

}

// DisplayAll godoc
// @Summary Display all file analysis statistics
// @Description Retrieve all records of file analysis statistics from the database.
// @Produce  json
// @Success 200 {array} db.FileStats "List of file statistics"
// @Failure 500 {object} map[string]string "Error fetching or processing data"
// @Router /display [get]
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

func Signup(c *gin.Context) {
	name := c.PostForm("Name")
	password := c.PostForm("Password")

	if name == "" && password == "" {
		c.JSON(http.StatusNoContent, gin.H{"Error": "One or More Fields are Empty"})
	}

	err := db.CreateUserData(name, password)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"Success": "Added User in Database Successfully"})
	}

}
func Login(c *gin.Context) {
	name := c.PostForm("Name")
	password := c.PostForm("Password")

	if name == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "One or More Fields are Empty"})
		return
	}

	// Fetch user from the database
	user, err := db.GetUserByName(name)
	if err != nil {
		fmt.Println("Error fetching user:", err) // Debug statement
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "User not found"})
		return
	}

	// Check if the password matches
	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(name)
	if err != nil {
		fmt.Println("Error generating token:", err) // Debug statement
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not generate token"})
		return
	}

	fmt.Println("Generated Token for user:", name, "Token:", token) // Debug statement

	// Send token in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}
