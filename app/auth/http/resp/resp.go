package resp

import "github.com/gin-gonic/gin"
import "fmt"

const (
	AcceptJson string = "application/json"
	AcceptXML = "application/xml"
)

func Resp (c *gin.Context) {
	accept := c.Request.Header.Get("Accept")

	if len(accept) == 0 {
		accept = AcceptJson	
	}

	switch accept {
	case AcceptJson:
		payload :=  gin.H{
			"status": "Ok",	
			"Content-Type": "application/json",
		}

		fmt.Println(c.Query("json_expanded"))
		if c.Query("json_expanded") == "true" {
			c.IndentedJSON(200, payload)
		} else {
			c.JSON(200, payload)
		} 
	case AcceptXML:
		c.XML(200, gin.H{
			"status": "Ok",
			"Content-Type": "application/xml",
		})
	}
}
