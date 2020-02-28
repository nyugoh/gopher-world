package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	db *sql.DB
)

func init()  {

	godotenv.Load()
	db, err := DbConnect()
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/bets/{bet_id}", getBets)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to read config file:", err.Error())
	}

	port := os.Getenv("PORT")
	log.Println("Starting server on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

type Outcome struct {
	MatchResultId int64 `json:"match_result_id"`
	SubTypeId int64 `json:"sub_type_id"`
	ParentMatchId int64 `json:"parent_match_id"`
	SpecialBetValue string `json:"special_bet_value"`
	LiveBet sql.NullInt64 `json:"live_bet"`
	CreatedBy string `json:"created_by"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Status sql.NullInt64 `json:"status"`
	ApprovedBy sql.NullString `json:"approved_by"`
	ApprovedStatus sql.NullInt64 `json:"approved_status"`
	DateApproved sql.NullString `json:"date_approved"`
	WinningoutCome string `json:"winning_outcome"`
}

func getBets(w http.ResponseWriter, r *http.Request)  {
	betId, _ := strconv.Atoi(mux.Vars(r)["bet_id"])
	fmt.Println("Request: ", betId)

	betoutcome := Outcome{}
	query := "SELECT * FROM uof_outcome WHERE match_result_id=? LIMIT 1"

	if err :=db.QueryRow(query, betId).Scan(&betoutcome.MatchResultId, &betoutcome.SubTypeId, &betoutcome.ParentMatchId,
		&betoutcome.SpecialBetValue, &betoutcome.LiveBet, &betoutcome.CreatedBy, &betoutcome.Created, &betoutcome.Modified,
		&betoutcome.Status, &betoutcome.ApprovedBy, &betoutcome.ApprovedStatus, &betoutcome.DateApproved, &betoutcome.WinningoutCome,); err!= nil {
		log.Println("Error fetching outcome: ", err.Error())
	}

	response, _ := json.Marshal(betoutcome)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func redisConnect() (*redis.Client, error) {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_SERVER"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"),
		DB:       db,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to redis:%s", err.Error()))
	}

	fmt.Println("Redis reply: ", pong)
	return client, nil
}

func DbConnect() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_URI"))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable too connect to DB: %s", err.Error()))
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("unable to ping DB")
	}
	fmt.Println("Connected to DB...")
	db.SetMaxIdleConns(80)
	db.SetMaxOpenConns(100)
	return db, nil
}