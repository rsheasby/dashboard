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
	"github.com/rs/cors" // Add this import
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
	mux.HandleFunc("/next-meeting", handleNextMeeting)

	handler := cors.Default().Handler(mux) // Add CORS handler

	fmt.Println("Started running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
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

	http.Redirect(w, r, "/next-meeting", http.StatusTemporaryRedirect)
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

type MeetingResponse struct {
	CurrentMeeting *Meeting `json:"currentMeeting,omitempty"`
	NextMeeting    *Meeting `json:"nextMeeting,omitempty"`
}

func handleNextMeeting(w http.ResponseWriter, r *http.Request) {
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

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).EventTypes("default").
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		response := "Unable to retrieve events: " + err.Error() + "\n<a href=\"/login\">Login</a>"
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, response)
		return
	}

	now := time.Now()
	var currentMeeting, nextMeeting *calendar.Event
	for _, event := range events.Items {
		endTime, err := time.Parse(time.RFC3339, event.End.DateTime)
		if err != nil {
			endTime, _ = time.Parse("2006-01-02", event.End.Date)
		}
		if now.Before(endTime) {
			if currentMeeting == nil {
				startTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
				if err != nil {
					startTime, _ = time.Parse("2006-01-02", event.Start.Date)
				}
				if now.After(startTime) {
					currentMeeting = event
				} else {
					nextMeeting = event
					break
				}
			} else {
				nextMeeting = event
				break
			}
		}
	}

	var response MeetingResponse
	if currentMeeting != nil {
		response.CurrentMeeting = &Meeting{
			Name:      currentMeeting.Summary,
			StartTime: currentMeeting.Start.DateTime,
			EndTime:   currentMeeting.End.DateTime,
		}
	}

	if nextMeeting != nil {
		response.NextMeeting = &Meeting{
			Name:      nextMeeting.Summary,
			StartTime: nextMeeting.Start.DateTime,
			EndTime:   nextMeeting.End.DateTime,
		}
	}

	if currentMeeting == nil && nextMeeting == nil {
		response = MeetingResponse{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
