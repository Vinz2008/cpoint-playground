package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func status(c *gin.Context) {
	c.Status(http.StatusOK)
}

type code_json struct {
	Code                        string `json:"code"`
	Should_return_ir            bool   `json:"should_return_ir"`
	Should_return_after_imports bool   `json:should_return_after_imports`
}

func run_code(c *gin.Context) {
	var code_to_run code_json
	if err := c.BindJSON(&code_to_run); err != nil {
		fmt.Println("error BindJSON")
		return
	}
	_, err := os.Stat("temp/")
	if !os.IsNotExist(err) {
		os.RemoveAll("temp/")
	}
	os.Mkdir("temp", os.ModePerm)
	f, err := os.Create("temp.cpoint")
	if err != nil {
		fmt.Println("error os.Create(\"temp.cpoint\")")
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(code_to_run.Code)
	out_stdout_compiler, err := exec.Command("cpoint", "temp.cpoint", "-no-delete-import-file").Output()
	if err != nil {
		fmt.Println("error exec.Command(\"cpoint\", \"temp.cpoint\").Run()")
		log.Fatal(err)
	}

	var ir string = ""

	if code_to_run.Should_return_ir {
		out_ir, err := os.ReadFile("out.ll")
		if err != nil {
			fmt.Println("error exec.ReadFile(\"out.ll\")")
			log.Fatal(err)
		}
		ir = string(out_ir)
		//fmt.Println(ir)
	}

	var after_imports string = ""

	if code_to_run.Should_return_after_imports {
		out_after_imports, err := os.ReadFile("temp.cpoint.temp")
		if err != nil {
			fmt.Println("error exec.ReadFile(\"temp.cpoint.temp\")")
			log.Fatal(err)
		}
		after_imports = string(out_after_imports)
	}
	//a_out_abs_path, err := filepath.Abs("./a.out")
	//cmd := exec.Command(a_out_abs_path)
	cmd := exec.Command("../a.out")
	cmd.Dir = "./temp/"
	fmt.Println(cmd.Dir)
	out_ran, err := cmd.Output()
	if err != nil {
		fmt.Println("error exec.Command(\"./a.out\").Output()")
		log.Fatal(err)
	}
	fmt.Println("out ran : ", string(out_ran))
	//c.IndentedJSON(http.StatusOK, code_to_run)
	//c.IndentedJSON(http.StatusOK, string(out))
	c.JSON(http.StatusOK, gin.H{"stdout_compiler": string(out_stdout_compiler), "output": string(out_ran), "output_ir": ir, "output_after_imports": after_imports})
	//c.JSON(http.StatusOK, code_to_run)
}

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	})

	//router.GET("/run-code", run_code)
	router.POST("/run-code", run_code)

	router.GET("/status", status)

	router.Run("localhost:8080")
}
