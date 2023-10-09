package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type code_json struct {
	Code string `json:"code"`
}

func run_code(c *gin.Context) {
	var code_to_run code_json
	if err := c.BindJSON(&code_to_run); err != nil {
		fmt.Print("error")
		return
	}

	f, err := os.Create("temp.cpoint")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(code_to_run.Code)
	err = exec.Command("cpoint", "temp.cpoint").Run()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command("./a.out").Output()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Print("code : ", code_to_run.code, "\n")
	//c.IndentedJSON(http.StatusOK, code_to_run)
	c.IndentedJSON(http.StatusOK, string(out))
	//c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	//c.JSON(http.StatusOK, gin.H{"status": code_to_run.code})
	//c.JSON(http.StatusOK, code_to_run)
}

func main() {
	router := gin.Default()
	//router.GET("/run-code", run_code)
	router.POST("/run-code", run_code)

	router.Run("localhost:8080")
}
