package main

import(
  "fmt"
  "encoding/base64"
  "bytes"
  "image"
  "image/png"
  "image/jpeg"
  "os"
  itob "basetoimagegolang/imagetobase64"
  uuid "github.com/satori/go.uuid"
)
func main(){
  fileName := uuid.NewV4().String()
  //strbase64 string base64 |  content type hanya ext | typefileWithbase type file base64
  strbase64, contentType , typefileWithbase:=  itob.PathToBase64()

  fmt.Println(contentType, "cek type ---")
  fmt.Println(typefileWithbase, "cek type with base ---")

  dec, err := base64.StdEncoding.DecodeString(strbase64)
  if err != nil {
      fmt.Println(err.Error())
  }

  var varImageCek image.Image
  if contentType == "image/jpeg" {
    varImageCek, err = jpeg.Decode(bytes.NewReader(dec))
    if err != nil {
      fmt.Println(err.Error())
    }
  }else if contentType == "image/png"{
    varImageCek, err = png.Decode(bytes.NewReader(dec))
    if err != nil {
      fmt.Println(err.Error())
    }
  }
  f_two, _ := os.OpenFile("static_file/output_versi_ngarang/"+fileName+".jpg", os.O_WRONLY | os.O_CREATE, 0777)
  jpeg.Encode(f_two, varImageCek, &jpeg.Options{Quality: 1080})
  fmt.Println(fileName, "gambar telah dibuat")

}
