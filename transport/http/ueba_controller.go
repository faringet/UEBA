package http

import (
	"UEBA/internal/csvreader"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UEBAController struct {
	CSVReader *csvreader.CSVReader
	logger    *zap.Logger
}

func NewConfigController(logger *zap.Logger, csvReader *csvreader.CSVReader) *UEBAController {
	return &UEBAController{logger: logger, CSVReader: csvReader}
}

func (uc *UEBAController) GetItemsByID(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No id provided"})
		uc.logger.Error("No id provided")
		return
	}

	record, err := uc.CSVReader.GetRecordByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling JSON data"})
		uc.logger.Error("Error marshalling JSON data")
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
