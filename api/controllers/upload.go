package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(50 << 20); err != nil { // 10MB max size
		SendErrorResponse(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		SendErrorResponse(w, "Failed to get file from request", http.StatusBadRequest)
		return
	}

	defer file.Close()
	currentTime := time.Now()

	// Format the current time to include date and time
	// Format: YYYY-MM-DD-HH-MM-SS
	currentDateTime := currentTime.Format("2006-01-02-15-04-05")
	valid, err := isValidFileType(file)
	if err != nil || !valid {
		SendErrorResponse(w, "Invalid file type", http.StatusBadRequest)
		return
	}
	// Create file on server
	fileName := fmt.Sprintf(currentDateTime + header.Filename)
	filePath := filepath.Join("../../client/public/uploads", fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		SendErrorResponse(w, "Failed to save file on server", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fileName)
}

func isValidFileType(file multipart.File) (bool, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return false, err
	}
	fileType := http.DetectContentType(buffer)
	validTypes := []string{"image/png", "image/jpeg", "image/gif", "image/webp"}

	for _, validType := range validTypes {
		if fileType == validType {
			if _, err := file.Seek(0, 0); err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}
