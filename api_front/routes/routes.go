package routes

import (
	"encoding/base64"
	"fmt"
	"github.com/Eitol/document_color_meter/api_front/config"
	"github.com/Eitol/document_color_meter/api_front/core"
	"github.com/Eitol/document_color_meter/api_front/core/binarizer"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Request struct {
	Binarize       *bool                 `form:"binarize" json:"binarize,omitempty"`
	WhiteThreshold uint32                `form:"whiteThreshold" json:"whiteThreshold"`
	BlackThreshold uint32                `form:"blackThreshold" json:"blackThreshold"`
	GrayThreshold  uint32                `form:"grayThreshold"  json:"grayThreshold"`
	File           *multipart.FileHeader `form:"file" binding:"required"`
}

type ResponseStatus string

const (
	BadRequest      = ResponseStatus("bad_request")
	BadFile         = ResponseStatus("bad_file")
	EmptyFile       = ResponseStatus("empty_file_file")
	ServerError     = ResponseStatus("server_error")
	ProcessingError = ResponseStatus("processing_error")
	BinarizingError = ResponseStatus("binarizing_error")
	Success         = ResponseStatus("success")
)

type Document struct {
	Name  string `json:"name" bson:"name"`
	Pages int    `json:"pages" bson:"pages"`
}

type Response struct {
	DocumentName   string                      `json:"documentName" bson:"documentName"`
	ResponseStatus ResponseStatus              `json:"responseStatus" bson:"responseStatus"`
	Result         *core.DocColorMeasureResult `json:"result" bson:"result"`
	Details        string                      `json:"details" bson:"details"`
	BinarizedPages []string                    `json:"binarizedPages" bson:"binarizedPages"`
	OutPath        string                      `json:"outPath" bson:"outPath"`
}

func multipartFileToBytes(mpFile *multipart.FileHeader) ([]byte, error) {
	if mpFile == nil {
		return nil, fmt.Errorf("nil file")
	}
	f, err := mpFile.Open()
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil || f == nil {
		return nil, err
	}
	return fileBytes, nil
}

func binarizeImages(images [][]byte) ([]string, string, error) {
	b := binarizer.NewBinarizerInstance(
		config.GetConfig().BinarizerHost,
		config.GetConfig().ExternalServicesTimeout,
	)
	binarizedImages, path, err := b.Binarize(images, binarizer.SAUVOLA)
	if err != nil {
		return nil, "", err
	}
	out := make([]string, 0, len(images))
	for _, image := range binarizedImages {
		strImg := base64.StdEncoding.EncodeToString(image)
		out = append(out, strImg)
	}
	return out, path, nil
}

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/measure_color", func(c *gin.Context) {
		var request Request

		// Bind multipartFile
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				DocumentName:   "",
				ResponseStatus: BadRequest,
				Details:        err.Error(),
			})
			return
		}
		bytes, err := multipartFileToBytes(request.File)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				DocumentName:   request.File.Filename,
				ResponseStatus: BadFile,
				Details:        err.Error(),
			})
			return
		}
		if len(bytes) == 0 {
			c.JSON(http.StatusBadRequest, Response{
				DocumentName:   request.File.Filename,
				ResponseStatus: EmptyFile,
			})
			return
		}
		meter := core.NewColorMeter()
		result, imgsBytes, err := meter.MeasureColor(bytes, core.Options{
			WhiteThreshold: request.WhiteThreshold,
			BlackThreshold: request.BlackThreshold,
			GrayThreshold:  request.GrayThreshold,
		})
		if err != nil {
			c.JSON(http.StatusOK, Response{
				DocumentName:   request.File.Filename,
				ResponseStatus: ProcessingError,
				Result:         nil,
				Details:        err.Error(),
			})
			return
		}
		var binarizedImagesB64 []string
		var path string
		if request.Binarize != nil && *request.Binarize {
			binarizedImagesB64, path, err = binarizeImages(imgsBytes)
		}
		if err != nil {
			c.JSON(http.StatusOK, Response{
				DocumentName:   request.File.Filename,
				ResponseStatus: BinarizingError,
				Result:         nil,
				Details:        err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			DocumentName:   request.File.Filename,
			ResponseStatus: Success,
			Result:         result,
			BinarizedPages: binarizedImagesB64,
			OutPath:        path,
		})
	})
	return router
}
