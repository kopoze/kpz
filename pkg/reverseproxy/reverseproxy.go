package reverseproxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kopoze/kpz/pkg/app"
	"github.com/kopoze/kpz/pkg/config"
	"golang.org/x/net/http2"
)

func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.LoadConfig()
		appDomain := fmt.Sprintf(".%s", conf.Kopoze.Domain)
		sub := strings.Replace(c.Request.Host, appDomain, "", 1)

		var currApp app.App

		if err := app.DB.Where("subdomain = ?", sub).First(&currApp).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "App not found!"})
			return
		}

		currUrl := fmt.Sprintf("http://localhost:%s", currApp.Port)

		appUrl, err := url.Parse(currUrl)
		if err != nil {
			log.Fatal(err)
		}
		// proxy := httputil.NewSingleHostReverseProxy(demoUrl)

		c.Request.Host = appUrl.Host
		c.Request.URL.Host = appUrl.Host
		c.Request.URL.Scheme = appUrl.Scheme
		c.Request.RequestURI = ""
		// c.Request.URL.Path = c.Param("proxyPath")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		host, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		c.Request.Header.Set("X-Forwarded-For", host)

		http2.ConfigureTransport(http.DefaultTransport.(*http.Transport))

		resp, err := http.DefaultClient.Do(c.Request)

		if err != nil {
			// TODO: Render a good 404 page
			c.Writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(c.Writer, "No App with the current subdomain found!")
			return
		}

		for key, values := range resp.Header {
			for _, value := range values {
				c.Writer.Header().Set(key, value)
			}
		}

		// Stream
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-time.Tick(10 * time.Millisecond):
					c.Writer.(http.Flusher).Flush()
				case <-done:
					return
				}
			}
		}()

		// Trailer
		trailerKeys := []string{}
		for key := range resp.Trailer {
			trailerKeys = append(trailerKeys, key)
		}

		c.Writer.Header().Set("Trailer", strings.Join(trailerKeys, ","))

		c.Writer.WriteHeader(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)

		for key, values := range resp.Trailer {
			for _, value := range values {
				c.Writer.Header().Set(key, value)
			}
		}

		close(done)
		// c.Abort()
	}
}
