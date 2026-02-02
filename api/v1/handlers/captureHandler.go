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

	if req.IfaceName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "interface name is required"})
		return
	}
	if req.SnapshotLen <= 0 || req.SnapshotLen > 65535 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "snapshot length must be between 1 and 65535"})
		return
	}
	if req.Timeout == 0 || req.Timeout < -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "timeout must be -1 (block forever) or greater than 0"})
		return
	}

	devices, err := utils.GetAllInterfaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get interfaces"})
		return
	}

	interfaceExists := false
	for _, dev := range devices {
		if dev.Name == req.IfaceName {
			interfaceExists = true
			break
		}
	}
	if !interfaceExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "interface not found"})
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
