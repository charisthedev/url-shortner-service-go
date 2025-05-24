package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	dbconfig "url-shortner/db/dbConfig"
	"url-shortner/internal/database"
	"url-shortner/internal/utils"
)

type CreatePayload struct {
	URL string `json:"url"`
}

func CreateShortenedUrl (w http.ResponseWriter, r *http.Request){
	var input CreatePayload
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w,http.StatusBadRequest,"Invalid request body")
		return
	}
	db := dbconfig.ConnectDb();
	defer db.Close()
	ctx := context.Background();
	tx, err := db.BeginTx(ctx, nil);
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback();
	qtx := database.New(tx)
	createdUrl, err := qtx.CreateURL(ctx,database.CreateURLParams{OriginalUrl: input.URL})
	if err != nil{
		utils.RespondWithError(w,http.StatusInternalServerError,"error creating short url");
		return
	}
	shortCode := utils.HashUrl(createdUrl.ID);
	err = qtx.UpdateShortCode(ctx,database.UpdateShortCodeParams{ShortCode:sql.NullString{String: shortCode, Valid: true},ID: createdUrl.ID});
	if err != nil {
		utils.RespondWithError(w,http.StatusInternalServerError,"Failed to create url")
		return
	}
	if err := tx.Commit(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError,"Failed to commit transaction")
		return
	}
	utils.RespondWithSuccess(w,200,"url created",map[string]string{"url": "http://localhost:5000/"+shortCode})
}