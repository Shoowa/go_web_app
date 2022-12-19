package webserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type BulkDetails struct {
	Thing   string `csv:"thing" json:"thing"`
	Details string `csv:"details" json:"details"`
}

func split_or_not(item series.Element) series.Element {
	list := strings.Fields(item.String())
	if len(list) > 1 {
		descriptors := strings.Join(list[:], ",")
		series.Element.Set(item, descriptors)
		return item
	}
	return item
}

func (e *Env) BulkCsvPOST(c *gin.Context) {

	// Identify file in web request.
	file, uploadErr := c.FormFile("file")
	if uploadErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": uploadErr.Error()})
	}

	// Open file.
	content, openErr := file.Open()
	if openErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": openErr.Error()})
	}

	// Read CSV content into a Gota Dataframe, then close file.
	df := dataframe.ReadCSV(content)
	content.Close()

	// Amend a Gota column.
	details := df.Col("details")
	details.Map(split_or_not)

	// Replace column in Gota Dataframe.
	df2 := df.Mutate(details)

	// Convert Gota Dataframe to JSON, and write it into a buffer.
	buf := new(bytes.Buffer)
	if bufErr := df2.WriteJSON(buf); bufErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": bufErr.Error()})
	}

	// Read from buffer.
	var deets []*BulkDetails
	if err := json.Unmarshal(buf.Bytes(), &deets); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Deliver JSON to client.
	c.IndentedJSON(http.StatusOK, deets)
}
