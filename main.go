package main

//source: https://github.com/rumyantseva/advent-2017/blob/master/main.go
// https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/
//testing: https://splice.com/blog/lesser-known-features-go-test/?utm_source=google&utm_medium=cpc&utm_campaign=Google_Search_Acquisition_Sounds_Nonbrand_DSA_ROW&utm_content=sounds&utm_term=&campaignid=13577111017&adgroupid=123041963239&adid=528665014304&gclid=CjwKCAjw9e6SBhB2EiwA5myr9hdj9Q9_BQ7aGeLsT2Jo-vAhU_jW0MydZubt6Yl9ASyZ1i9Lh9zmjRoC1YoQAvD_BwE

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os" // to access env vars
	"os/signal"
	"syscall"

	"log" //"github.com/golang/glog"

	"microservice2/handlers"
	"microservice2/vn"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

// How to try it: PORT=80 go run main.go
func main() {
	/*
		glog.V(2).Info("main called")
		glog.Error("test Err Msg")
		glog.Infof(
			"Starting the service...\n\tcommit: %s, build time: %s, release: %s",
			version.Commit, version.BuildTime, version.Release,
		)
	*/
	log.Print("main called")
	log.Printf(
		"Starting the service...\n\tcommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		//glog.Fatal("env var PORT is not set.")
		log.Print("env var PORT is not set, using 80")
		port = "80"
	}

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// this channel is for graceful shutdown:
	// if we receive an error, we can send it here to notify the server to be stopped
	shutdown := make(chan struct{}, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			//glog.Errorf("%v", err)
			log.Fatalf("%v", err)
		}
	}()
	//glog.Info("The service is ready to listen and serve.")
	log.Printf("The service is ready to listen and serve on port %s.", port)

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	case <-shutdown:
		log.Fatal("Got an error...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
