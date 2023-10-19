package delivery

import (
	"fmt"
	"io"
	"net/http"

	"file-storage/internal/entities"
	"file-storage/internal/repositories"
	"file-storage/internal/usecases"
	"file-storage/internal/utils"
)

type FileHandler struct {
	repo *repositories.ServerRepository
}

func NewFileHandler(repo *repositories.ServerRepository) *FileHandler {
	return &FileHandler{repo: repo}
}

func (h *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Прочитать файл из запроса
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.ErrorLogger.Println("Unable to read file:", err)
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		utils.ErrorLogger.Println("Unable to reading file:", err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Здесь можно добавить генерацию ID для файла, например, с помощью UUID
	fileID := "some_generated_id"
	fileEntity := entities.File{
		ID:   fileID,
		Data: data,
	}

	// Разделить файл на части
	parts, err := usecases.SplitFileIntoParts(fileEntity, 6)
	if err != nil {
		utils.ErrorLogger.Println("Unable to split file:", err)
		http.Error(w, "Error splitting file", http.StatusInternalServerError)
		return
	}

	// Сохранить части на серверах хранения
	err = h.repo.StoreFileParts(parts)
	if err != nil {
		utils.ErrorLogger.Println("Unable to save file:", err)
		http.Error(w, "Error storing file parts", http.StatusInternalServerError)
		return
	}

	// Вернуть ответ
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded with ID: %s", fileID)
}

func (h *FileHandler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	// Получить ID файла из параметров запроса
	fileID := r.URL.Query().Get("id")
	if fileID == "" {
		utils.ErrorLogger.Println("Unable to download file:", fileID)
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	// Получить части файла с серверов хранения
	parts, err := h.repo.RetrieveFileParts(fileID, 6)
	if err != nil {
		utils.ErrorLogger.Println("Unable to retrieve file:", err)
		http.Error(w, "Error retrieving file parts", http.StatusInternalServerError)
		return
	}

	// Собрать файл из частей
	combinedFile, err := usecases.CombineFileFromParts(parts)
	if err != nil {
		utils.ErrorLogger.Println("Unable to combine file:", err)
		http.Error(w, "Error combining file parts", http.StatusInternalServerError)
		return
	}

	// Отправить файл обратно пользователю
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileID))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(combinedFile.Data)
	if err != nil {
		utils.ErrorLogger.Println("Unable to send file:", err)
		return
	}
}
