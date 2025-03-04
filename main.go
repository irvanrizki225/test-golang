package main

import (
	// "fmt"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences []string `json:"influences"`
}

type Language struct {
	Language string `json:"language"`
	Appeared int `json:"appeared"`
	Created []string `json:"created"`
	Functional bool `json:"functional"`
	ObjectOriented bool `json:"object-oriented"`
	Relation Relation `json:"relation"`
}


var (
	language = []Language{
		{
			Language: "Java",
			Appeared: 1995,
			Created: []string{"James Gosling"},
			Functional: true,
			ObjectOriented: true,
			Relation: Relation{
				InfluencedBy: []string{"C++", "C#"},
				Influences: []string{"Objective-C", "Swift"},
			},
		},
		{
			Language: "Go",
			Appeared: 2009,
			Created: []string{
				"Robert Griesemer",
				"Rob Pike",
				"Ken Thompson",
			},
			Functional: true,
			ObjectOriented: false,
			Relation: Relation{
				InfluencedBy: []string{
					"C++",
					"Java",
				},
				Influences: []string{
					"JavaScript",
					"Python",
				},
			},
		},
	}
)

func main(){
	ExampleRestApi()
}

//nomor 1 func palindrom dan reverse string
func isPalindrome(s string) bool{
	cleaned := strings.Map(func(r rune) rune {
		if (unicode.IsLetter(r) || unicode.IsNumber(r)) {
			return -1
		} 

		return unicode.ToLower(r)
	}, s)

	return cleaned == reverseString(cleaned)
}

func reverseString(s string) string{
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ExampleRestApi() {
	router := gin.Default()

	//endpoint no 3
	router.GET("/language", func(c *gin.Context) {
		//binding data model ke json
		data := language
		c.JSON(http.StatusOK, data)
	})

	//endpoint no 5
	router.POST("/language", func(c *gin.Context) {
		//binding request body ke json
		var request Language
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//binding data model ke json
		data := request

		//menambah data model request ke language array
		language = append(language, data)

		c.JSON(http.StatusOK, data)
	})
	router.GET("/language/:id", func(c *gin.Context) {
		//index
		id := c.Param("id")
		var index int

		//binding data model ke json
		for i := 0; i == len(language); i++ {
			if language[i].Language == id {
				index = i
			}
		}

		//check language id ada atau tidak
		if _, err := fmt.Sscanf(id, "%d", &index); err != nil || index < 0 || index >= len(language) {
			c.JSON(http.StatusNotFound, gin.H{"error": "language not found"})
			return
		}		

		c.JSON(http.StatusOK, language[index])
	})

	router.PATCH("/language/:id", func(c *gin.Context) {
		id := c.Param("id")
		var index int

		//binding data model ke json
		for i := 0; i < len(language); i++ {
			if language[i].Language == id {
				index = i
			}
		}


		//binding request body ke json
		var request Language
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//mengubah data model ke json
		data := request
		language[index] = data
		c.JSON(http.StatusOK, language[index])
	})

	router.DELETE("/language/:id", func(c *gin.Context) {
		id := c.Param("id")
		var index int

		//binding data model ke json
		for i := 0; i < len(language); i++ {
			if language[i].Language == id {
				index = i
			}
		}

		//menghapus data model ke json
		language = append(language[:index], language[index+1:]...)
		c.JSON(http.StatusOK, gin.H{"message": "language deleted"})
	})


	//endpoint no 4
	router.POST("/palindrome", func(c *gin.Context) {
		//binding request body ke json
		var data map[string]string
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//cek palindrom
		if isPalindrome(data["text"]) {
			c.JSON(http.StatusOK, gin.H{"message": "palindrome"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "not palindrome"})
		}
	})

	//endpoint no 5

	router.Run(":8080")
}