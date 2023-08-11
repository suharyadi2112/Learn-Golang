package main

import (
  // "os"
  "github.com/sirupsen/logrus"
  "errors"
)

var log = logrus.New()

func init() {
  // Log as JSON instead of the default ASCII formatter.
  // log.SetFormatter(&logrus.JSONFormatter{})
  log.SetFormatter(&logrus.TextFormatter{
    DisableColors: false,
    FullTimestamp: true,  
    ForceColors:  true,
  })
  // Only log the debug severity or above.
  log.SetLevel(logrus.DebugLevel)
}

func main() {
  ctx := "LogrusToLogFile" //context

  // log.Out = os.Stdout
  // // You could set this to any `io.Writer` such as a file
  // file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  // if err == nil {
  //  log.Out = file
  // } else {
  //   log.Info("Failed to log to file, using default stderr")
  // }

  //make fake error
  err := errors.New("math: square root of negative number")
  if err != nil {
    log.WithFields(logrus.Fields{
      "ctx": ctx,
      "contoh": "tes",
    }).Error("Write to file")
  }

}

// <img width="100%" style="width:100%" src="https://media.giphy.com/media/25KEhzwCBBFPb79puo/giphy-downsized.gif">

 // log.SetFormatter(&log.JSONFormatter{

 //  })

 //  log.SetFormatter(&log.TextFormatter{
 //    DisableColors: true,
 //    FullTimestamp: true,  
 //  })

 //  // log.SetReportCaller(true)

   // contextLogger := log.WithFields(log.Fields{
  //   "common": "this is a common field",
  //   "other": "I also should be logged always",
  // })

  // contextLogger.Info("I'll be logged with common and other field")
  // log.Warn("Warn Me too")
  // log.Fatal("Fatal Me too")