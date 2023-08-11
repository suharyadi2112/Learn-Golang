package main

import(
  "fmt"
  "encoding/base64"
  "os"
  "errors"

  "net/http"
  "strings"
  "github.com/jasonlvhit/gocron"
  itob "basetoimagegolang/imagetobase64"
  uuid "github.com/satori/go.uuid"
)
var (
	imageContentTypeWhitelist = []string{"image/gif", "image/jpeg", "image/png"}
)
//cron job menjalankan print gambar dari base64 ke image
func main(){
  gocron.Every(60).Second().Do(GetData)
  <- gocron.Start()

}

func GetData(){
  fileName := uuid.NewV4().String()
  //strbase64 string base64 |  content type hanya ext | typefileWithbase type file base64
  strbase64, _ , _:=  itob.PathToBase64()

  dec, err := base64.StdEncoding.DecodeString(strbase64)
  if err != nil {
      fmt.Println(err.Error())
  }
  // Get content type and file extention
  _, extention, err := GetFileContentTypeAndExtention(dec)
  if err != nil {
      fmt.Println(err.Error())
  }

  extention = "." + extention

  localPath := "static_file/output_versi_mansek/" + fileName + extention
  f, err := os.Create(localPath)
  if err != nil {
       fmt.Println(err.Error())
  }
  defer f.Close()

  if _, err := f.Write(dec); err != nil {
     fmt.Println(err.Error())
  }
  if err := f.Sync(); err != nil {
     fmt.Println(err.Error())
  }

  fmt.Println(fileName, "gambar telah dibuat")

}

// GetFileContentTypeAndExtention ...
func GetFileContentTypeAndExtention(file []byte) (contentType string, extention string, err error) {
	contentType = http.DetectContentType(file)

	if !Contains(imageContentTypeWhitelist, contentType) {
		return contentType, contentType, errors.New("invalid_content_type")
	}

	contentTypeArr := strings.Split(contentType, "/")
	if len(contentTypeArr) != 2 {
		return contentType, contentType, errors.New("invalid_content_type_length")
	}

	extention = contentTypeArr[1]

	return contentType, extention, nil
}

// Contains ...
func Contains(slices []string, comparizon string) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}
	return false
}
