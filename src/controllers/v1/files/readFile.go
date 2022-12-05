package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"video-conversion-service/src/configs/funtions"
)

// ReadAFile @BasePath /api/v1
// @Tags Files
// readFile godoc
// @Summary read a file from specific location
// @Schemes
// @Description execution will read a file from specific location and make a stream
// @Param mountPathWithName query string true "File Name with the full mounted path"
// @Accept json
// @Produce json
// @Success 200
// @Failure 404 "page not found"
// @Router /files/read [get]
func ReadAFile(gtx *gin.Context) {
	mountPathWithName, ok := gtx.GetQuery("mountPathWithName")
	if ok != true {
		funtions.ErrorResponse(gtx, "File read payload mulfuntion! ", nil)
		return
	}
	fileName := strings.Split(mountPathWithName, "/")
	fileBytes, err2 := os.ReadFile(mountPathWithName)
	if err2 != nil {
		panic(err2)
	}
	gtx.Header("Content-Disposition", "attachment; filename="+fileName[len(fileName)-1])
	gtx.Data(http.StatusOK, "application/octet-stream", fileBytes)
}
