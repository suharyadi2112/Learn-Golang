package imagetobase64

import (
  "encoding/base64"
  "io/ioutil"
  "os"
  "net/http"
)

// GetFileContentType ...
func GetFileContentType(out *os.File) (typefileWithbase, contentType string, err error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = out.Read(buffer)
	if err != nil {
		return "","", err
	}
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType = http.DetectContentType(buffer)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch contentType {
	case "image/jpeg":
		typefileWithbase += "data:image/jpeg;base64,"
	case "image/png":
		typefileWithbase += "data:image/png;base64,"
	}
	return typefileWithbase, contentType, nil
}
// PathToBase64 ...
func PathToBase64() (res, contentType, typefileWithbase string) {
	// Open file from path
  cekSumber := "static_file/sumber_image/download.jpg"
	file, err := os.Open(cekSumber)
	if err != nil {
		return res, contentType, typefileWithbase
	}
	defer file.Close()
	// Get the content
	typefileWithbase, contentType, err = GetFileContentType(file)
	if err != nil {
		return res, contentType, typefileWithbase
	}
	// Read file as byte[]
	data, err := ioutil.ReadFile(cekSumber)
	if err != nil {
		return res, contentType, typefileWithbase
	}
	// Encode as base64.
	res = base64.StdEncoding.EncodeToString(data)
	return res, contentType, typefileWithbase
}
