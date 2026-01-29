package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heshanthenura/sentrigov2/internal/capture"
	"github.com/heshanthenura/sentrigov2/internal/job"
	"github.com/heshanthenura/sentrigov2/internal/types"
	"github.com/heshanthenura/sentrigov2/internal/utils"
)

func GetAllInterfaces(c *gin.Context) {
	devices, err := utils.GetAllInterfaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

func StartCapture(c *gin.Context) {
	req, err := utils.BindJSON[types.CaptureConfig](c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.IfaceName == "" || req.SnapshotLen == 0 || req.Timeout == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "all fields must be filled"})
		return
	}

	if job.GetJobManager().IsRunning("captureJob") {
		c.JSON(http.StatusConflict, gin.H{"error": "capture already running"})
		return
	}

	job.GetJobManager().Start(
		"captureJob",
		func(ctx context.Context) error {
			return capture.StartCapture(ctx, req)
		},
		func(err error) {
			log.Printf("Capture job failed: %v", err)
		},
	)

	c.JSON(http.StatusOK, gin.H{"request": req})
}

func StopCapture(c *gin.Context) {
	job.GetJobManager().Cancel("captureJob")
	c.JSON(http.StatusOK, gin.H{"request": "stop capture called"})

}
