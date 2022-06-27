package excel

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func PreprocessImport(c *gin.Context) (*excelize.File, error) {
	currentTime := time.Now().Unix()

	r := c.Request
	defer r.Body.Close()

	file, fileHeader, err := r.FormFile("excel-file")
	if err != nil {
		return nil, errors.New("Invalid uploaded file")
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Invalid uploaded file")
	}

	fileNameWithExt := fileHeader.Filename
	fileExt := filepath.Ext(fileNameWithExt)
	fileNameOnly := strings.TrimSuffix(fileNameWithExt, fileExt)
	fileNameSaved := fmt.Sprintf("%s_%d%s", fileNameOnly, currentTime, fileExt)
	filePath := "public/import/" + fileNameSaved

	err = ioutil.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		return nil, errors.New("Internal server error")
	}

	excel, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, errors.New("Internal server error")
	}

	return excel, nil
}
