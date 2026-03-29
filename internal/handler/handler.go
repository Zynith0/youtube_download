package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"text/template"
)

var folderName string

type Video struct {
	Url    string `json:"url"`
	Format string `json:"format"`
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/index.html")
}

func Download(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	format := r.FormValue("format")

	num1 := rand.Intn(9)
	num2 := rand.Intn(9)
	num3 := rand.Intn(9)
	num4 := rand.Intn(9)

	folderName = strconv.Itoa(num1) + strconv.Itoa(num2) + strconv.Itoa(num3) + strconv.Itoa(num4)

	os.Mkdir("./videos/"+folderName, 0755)

	cmd := exec.Command("yt-dlp", "-t", format, url)
	if format == "wav" {
		cmd = exec.Command("yt-dlp", "-x", "--audio-format", "wav", url)
	}
	cmd.Dir = "./videos/" + folderName
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	cmd2 := exec.Command("zip", "-r", folderName, folderName)
	cmd2.Dir = "./videos"
	output2, err := cmd2.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output2))
}

func HandleServeFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=\"Videos.zip\"")
	http.ServeFile(w, r, "./videos/"+folderName+".zip")
}

func renderTemplate(w http.ResponseWriter, file string) {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println("Failed to parse html template", err)
	}
	tmpl.Execute(w, nil)
}
