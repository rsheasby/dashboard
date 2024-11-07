package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var (
	oauthConfig *oauth2.Config
	client      *http.Client
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	oauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{calendar.CalendarReadonlyScope},
		Endpoint:     google.Endpoint,
	}

	token, err := tokenFromFile()
	if err != nil {
		fmt.Println("Token file not found or invalid, please login to generate a new token.")
		return
	}

	ts := oauthConfig.TokenSource(context.Background(), token)
	client = oauth2.NewClient(context.Background(), ts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleMain)
	mux.HandleFunc("/login", handleGoogleLogin)
	mux.HandleFunc("/callback", handleGoogleCallback)
	mux.HandleFunc("/meetings", handleMeetings)

	handler := cors.Default().Handler(mux)

	certFile := os.Getenv("TLS_CERT_FILE")
	keyFile := os.Getenv("TLS_KEY_FILE")

	fmt.Println("Started running on https://localhost:443")
	log.Fatal(http.ListenAndServeTLS(":443", certFile, keyFile, handler))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html><body><a href="/login">Google Log In</a></body></html>`
	fmt.Fprint(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	saveToken(token)

	ts := oauthConfig.TokenSource(context.Background(), token)
	client = oauth2.NewClient(context.Background(), ts)

	http.Redirect(w, r, "/todays-meetings", http.StatusTemporaryRedirect)
}

func saveToken(token *oauth2.Token) {
	file, err := os.Create("token.json")
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(token)
}

func tokenFromFile() (*oauth2.Token, error) {
	file, err := os.Open("token.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	return token, err
}

type Meeting struct {
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func handleMeetings(w http.ResponseWriter, r *http.Request) {
	if client == nil {
		http.Error(w, "Client not initialized. Please login first.", http.StatusUnauthorized)
		return
	}

	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		response := "Unable to retrieve events: " + err.Error() + "\n<a href=\"/login\">Login</a>"
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, response)
		return
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Format(time.RFC3339)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999, now.Location()).Format(time.RFC3339)

	events, err := srv.Events.List("primary").ShowDeleted(false).EventTypes("default").
		SingleEvents(true).TimeMin(startOfDay).TimeMax(endOfDay).OrderBy("startTime").Do()
	if err != nil {
		response := "Unable to retrieve events: " + err.Error() + "\n<a href=\"/login\">Login</a>"
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, response)
		return
	}

	var meetings []Meeting
	for _, event := range events.Items {
		meetings = append(meetings, Meeting{
			Name:      event.Summary,
			StartTime: event.Start.DateTime,
			EndTime:   event.End.DateTime,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meetings)
}
