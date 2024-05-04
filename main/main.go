package main

import (
	"fmt"
	"hapemu/model"
	"hapemu/service"
	"log"
	"net/http"
)

func main() {
	var dxomarkList []model.DxoMark = service.InsertDxoMarkToDatabase()
	for _, element := range dxomarkList {
		if element.Camera != 0 {
			fmt.Println(element.Name)
			fmt.Printf("Photo: %d\n", element.Mobile.Subscores.Photo)
			fmt.Printf("Bokeh: %d\n", element.Mobile.Subscores.Bokeh)
			fmt.Printf("Preview: %d\n", element.Mobile.Subscores.Preview)
			fmt.Printf("Zoom: %d\n", element.Mobile.Subscores.Zoom)
			fmt.Printf("Video: %d\n", element.Mobile.Subscores.Video)
			// fmt.Println(element.Battery)
			// fmt.Println(element.Camera)
			// fmt.Println(element.Display)
			// fmt.Println(element.Selfie)
			fmt.Println()
		}
	}

	// http.HandleFunc("/antutu", service.GetAntutuList)
	// http.HandleFunc("/dxomark", service.GetDxoMarkList)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
