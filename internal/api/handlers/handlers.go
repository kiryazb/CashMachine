package handlers

import (
	"CashMachine/internal/api/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Handlers struct {
	DB     *gorm.DB
	Logger *slog.Logger
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}	


func (h *Handlers) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item

	err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	if err := h.DB.Create(&item).Error; err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	h.Logger.Info("Add Item into db", "module", "handlers", "id", item.Id)
}


func (h *Handlers) ReadItem(w http.ResponseWriter, r *http.Request) {
	
	var item model.Item

	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))

	if err != nil {
		http.Error(w, "Incorrect user id", http.StatusBadRequest)
		h.Logger.Error("Bad Request with user id", "module", "handlers", "id", userID)
	}

	if err := h.DB.First(&item, userID).Error; err != nil {
		http.Error(w, "Incorrect item id", http.StatusBadRequest)
		h.Logger.Error("Bad Request with item id", "module", "handlers", "id", userID)
	}

	fmt.Fprintf(w, "Item model with ID %d\nTitle: %s\nPrice: %d", item.Id, item.Title, item.Price)
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) UpdateItem(w http.ResponseWriter, r *http.Request) {

	var UpdateItem model.Item
	var item model.Item

	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))

	if err != nil {
		http.Error(w, "Incorrect user id", http.StatusBadRequest)
		h.Logger.Error("Bad Request with user id", "module", "handlers", "id", userID)
	}

	if err := h.DB.First(&item, userID).Error; err != nil {
		http.Error(w, "Incorrect item id", http.StatusBadRequest)
		h.Logger.Error("Bad Request with item id", "module", "handlers", "id", userID)
	}

	if err := json.NewDecoder(r.Body).Decode(&UpdateItem); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
	}

	if err := h.DB.First(&item, userID).Error; err != nil {
		http.Error(w, "Incorrect item id", http.StatusBadRequest)
		h.Logger.Error("Bad Request with item id", "module", "handlers", "id", userID)
	}

	item.Title = UpdateItem.Title
	item.Price = UpdateItem.Price

	if err := h.DB.Save(&item).Error; err != nil {
		http.Error(w, "Error during save Update Item", http.StatusInternalServerError)
		h.Logger.Error("Error during save Update Item", "module", "handlers")
	}

	fmt.Fprintf(w, "Update Item with ID %d\nNew Title: %s\nNew Price: %d", userID, UpdateItem.Title, UpdateItem.Price)
	w.WriteHeader(http.StatusOK)

}


func (h *Handlers) DeleteItem(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))

	if err != nil {
		http.Error(w, "Incorrect user id", http.StatusBadRequest)
		h.Logger.Error("Bad Request with user id", "module", "handlers", "id", userID)
	}

	if err := h.DB.Where("id = ?", userID).Delete(&model.Item{}).Error; err != nil {
		http.Error(w, "Error during delete Item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}