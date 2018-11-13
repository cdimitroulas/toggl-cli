package api

import (
	"fmt"
	_ "github.com/cdimitroulas/toggl-cli/src/utils"
	"io/ioutil"
	"log"
	"net/http"
)

type Workspace struct {
	id   int
	name string
	at   string
}

type AuthenticationRequest struct {
	user     string
	password string
}

type AuthenticationResponseData struct {
	Id                    int64       `json:"id"`
	ApiToken              string      `json:"api_token"`
	DefaultWid            int         `json:"default_wid"`
	Email                 string      `json:"email"`
	Fullname              string      `json:"fullname"`
	JqueryTimeofdayFormat string      `json:"jquery_timeofday_format"`
	JqueryDateFormat      string      `json:"jquery_date_format"`
	TimeofdayFormat       string      `json:"timeofday_format"`
	DateFormat            string      `json:"date_format"`
	StoreStartAndStopTime bool        `json:"store_start_and_stop_time"`
	BeginningOfWeek       int64       `json:"beginning_of_week"`
	Language              string      `json:"language"`
	ImageUrl              string      `json:"image_url"`
	NewBlogPost           struct{}    `json:"new_blog_post"`
	Projects              []string    `json:"projects"`
	Tags                  []string    `json:"tags"`
	Tasks                 []string    `json:"tasks"`
	Workspaces            []Workspace `json:"workspaces"`
	Clients               []string    `json:"clients"`
}

type AuthenticationResponse struct {
	Since int                        `json:"since"`
	Data  AuthenticationResponseData `json:"data"`
}

func handleResponse(response *http.Response) (*AuthenticationResponse, error) {
	responseData := new(AuthenticationResponse)

	var httpError error
	if response.StatusCode != http.StatusOK {
		body, ioReadErr := ioutil.ReadAll(response.Body)
		if ioReadErr != nil {
			log.Fatalln(ioReadErr)
			return nil, ioReadErr
		}

		httpError = newApiError(response.StatusCode, string(body))
		return nil, httpError
	}

	return responseData, httpError
}

func AuthenticateWithToken(token string) *AuthenticationResponse {
	httpClient := &http.Client{}
	request, err := http.NewRequest(
		"GET",
		"https://www.toggl.com/api/v8/me",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	request.SetBasicAuth(token, "api_token")

	fmt.Println("Authenticating...")
	response, requestError := httpClient.Do(request)
	if requestError != nil {
		log.Fatalln(requestError)
	}

	defer response.Body.Close()

	data, responseError := handleResponse(response)
	if responseError != nil {
		log.Fatalln(responseError)
	}

	return data
}
