package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/play", func(ctx *gin.Context) {
		file, err := os.Open("../files/tvari.mp3")
		if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
            return
        }
		defer file.Close()

		ctx.Header("Content-Description", "File Transfer")
        ctx.Header("Content-Type", "audio/mpeg")
		ctx.Header("Accept-Ranges", "bytes")
        ctx.Header("Content-Transfer-Encoding", "binary")
        ctx.Header("Expires", "0")
        ctx.Header("Cache-Control", "must-revalidate")
        ctx.Header("Pragma", "public")
        ctx.Header("Content-Length", "application/octet-stream")

		_, err = io.Copy(ctx.Writer, file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to stream file"})
            return
		}

	})

	server.Run()

}