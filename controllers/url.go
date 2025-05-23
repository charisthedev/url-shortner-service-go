package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	dbconfig "url-shortner/db/dbConfig"
	"url-shortner/internal/database"
	"url-shortner/internal/utils"

	"github.com/mattheath/base62"
)

type CreatePayload struct {
	url string
}

func createShortenedUrl (w http.ResponseWriter, r *http.Request){
	var input CreatePayload
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	db := dbconfig.ConnectDb();
	ctx := context.Background();
	tx, err := db.BeginTx(ctx, nil);
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	qtx := database.New(tx)
	createdUrl, err := qtx.CreateURL(ctx,database.CreateURLParams{OriginalUrl: input.url})
	if err != nil{
		utils.RespondWithError(w,http.StatusInternalServerError,"error creating short url");
		return
	}
	shortCode := base62.EncodeInt64(int64(createdUrl.ID));
	err = qtx.UpdateShortCode(ctx,database.UpdateShortCodeParams{ShortCode:sql.NullString{String: shortCode, Valid: true},ID: createdUrl.ID});
	if err != nil {
		tx.Rollback()
		utils.RespondWithError(w,http.StatusInternalServerError,"Failed to create url")
		return
	}
	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}
	utils.RespondWithSuccess(w,200,"url created",map[string]string{"url": "http://localhost:5000/"+shortCode})
}