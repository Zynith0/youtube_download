package ytdlp

import (
	"log"
	"os/exec"
)

func DownloadVideo(url string, format string) {
	if format == "wav" {
		cmd := exec.Command("yt-dlp", "-x", "--audio-format", "wav", url)
		cmd.Dir = "./videos"
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		log.Println(string(output))
	}
	cmd := exec.Command("yt-dlp", "-t", format, url)
	cmd.Dir = "./videos"
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(output))
}
