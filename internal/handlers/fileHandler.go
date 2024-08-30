package handlers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/markbates/goth/gothic"
	"github.com/mwaurathealex/mbumwa3d/internal/stl"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
)

type FileHandler struct {
	SessionName string
}

type FileHandlerParams struct {
	SessionName string
}

func NewFileHandler(params FileHandlerParams) *FileHandler {
	if params.SessionName == "" {
		panic("NewFileHandler is missing SessionName")
	}
	return &FileHandler{
		SessionName: params.SessionName,
	}
}

func (h *FileHandler) Post(w http.ResponseWriter, r *http.Request) {
	errorEventPayload := &GetToastPayloadParams{
		EventName: "FileConfigUploadEvent",
		Message:   "error",
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		errorEventPayload.Description = "File size should be less than 10MB"
		w.Header().Add("HX-Trigger", GetToastPayload(errorEventPayload))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Parsing STL File Data: ", err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		errorEventPayload.Description = "An STL file is required."
		w.Header().Add("HX-Trigger", GetToastPayload(errorEventPayload))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Parsing STL File Data: ", err)
		return
	}

	filestore := dbstore.NewFileStore()
	fileNameInDisk, err := filestore.SaveToDisk(file, handler.Filename)

	filePath := filepath.Join(filestore.FileDir, fileNameInDisk)
	stlCalc, err := stl.NewSTLCalc(filePath)
	defer stlCalc.Close()
	fileVolume, err := stlCalc.GetVolume("cm")

	if err != nil {
		errorEventPayload.Description = "Something went wrong. Please try again"
		w.Header().Add("HX-Trigger", GetToastPayload(errorEventPayload))
		fmt.Println("ERROR: ", err)
		return
	}

	printConfig := store.PrintConfig{
		FileVolume: fileVolume,
		FileID:     fileNameInDisk,
	}

	session, _ := gothic.Store.Get(r, h.SessionName)
	session.Values["print_config"] = printConfig
	err = session.Save(r, w)
	if err != nil {
		w.Header().Add("HX-Trigger", GetToastPayload(&GetToastPayloadParams{
			EventName:   "FileConfigUploadEvent",
			Message:     "error",
			Description: "Something went wrong. Please try again.",
		}))
		log.Fatalf("Error adding PrintConfig to session: ", err)
		return
	}

	w.Header().Add("HX-Trigger", GetToastPayload(&GetToastPayloadParams{
		EventName:   "FileConfigUploadEvent",
		Message:     "success",
		Description: "File uploaded successfully",
	}))

	w.WriteHeader(http.StatusOK)
}
