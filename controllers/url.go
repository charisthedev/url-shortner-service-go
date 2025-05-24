package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	dbconfig "url-shortner/db/dbConfig"
	"url-shortner/internal/database"
	"url-shortner/internal/utils"

	"github.com/go-chi/chi/v5"
)

type CreatePayload struct {
	URL string `json:"url"`
}

func CreateShortenedUrl (w http.ResponseWriter, r *http.Request){
	var input CreatePayload
	hostUrl := os.Getenv("HOST_URL")
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
	log.Println(shortCode,createdUrl.ID)
	err = qtx.UpdateShortCode(ctx,database.UpdateShortCodeParams{ShortCode:sql.NullString{String: shortCode, Valid: true},ID: createdUrl.ID});
	if err != nil {
		utils.RespondWithError(w,http.StatusInternalServerError,"Failed to create url")
		return
	}
	if err := tx.Commit(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError,"Failed to commit transaction")
		return
	}
	utils.RespondWithSuccess(w,200,"url created",map[string]string{"url": hostUrl+shortCode})
}

func VisitUrl (w http.ResponseWriter, r *http.Request) {
	input := chi.URLParam(r,"short_code");
	db := dbconfig.ConnectDb();
	defer db.Close()
	ctx := context.Background();
	queries := database.New(db);
	data, err := queries.GetURLByShortCode(ctx,sql.NullString{String: input, Valid: true});
	if err != nil{
		utils.RespondWithError(w,http.StatusNotFound,"invalid url")
		return
	}
	utils.RespondWithRedirect(w,r,data.OriginalUrl,http.StatusFound)
}