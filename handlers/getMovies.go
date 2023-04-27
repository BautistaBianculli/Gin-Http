package handlers

import (
	"GORUTINE/api"
	"GORUTINE/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func GetMoovies(c *gin.Context) {
	wg := sync.WaitGroup{}
	moovie := make(chan []models.Moovies, 10)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go api.GetMoovies(moovie, i, &wg)
	}

	var allMovies []models.Moovies

	go func() {
		wg.Wait()
		close(moovie)
	}()

	for movies := range moovie {
		allMovies = append(allMovies, movies...)
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "ok",
		"value"	:	allMovies,
	})
}
