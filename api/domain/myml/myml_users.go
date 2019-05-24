package myml

import (
    "encoding/json"
    "fmt"
    "github.com/mercadolibre/ejercicioCircuitBreaker/src/api/utils/apierrors"
    "github.com/sony/gobreaker"
    "io/ioutil"
    "net/http"
)

type User struct {
    ID               int64  `json:"id"`
    Nickname         string `json:"nickname"`
    RegistrationDate string `json:"registration_date"`
    CountryID        string `json:"country_id"`
    Email            string `json:"email"`
}


const urlUsers = "http://localhost:8081/.users/"

var cb *gobreaker.CircuitBreaker

func init() {
    var st gobreaker.Settings
    st.Name = "HTTP GET"
    st.ReadyToTrip = func(counts gobreaker.Counts) bool {
        failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
        return counts.Requests >= 3 && failureRatio >= 0.6
    }

    cb = gobreaker.NewCircuitBreaker(st)
}

func (user *User) Get() *apierrors.ApiError {

    if user.ID == 0 {
        return &apierrors.ApiError{
            Message: "userID is empty",
            Status:  http.StatusBadRequest,
        }
    }
  final := fmt.Sprintf("%s%d", urlUsers, user.ID)
    // Se implementa patron
    _, err := cb.Execute(func() (interface{}, error) {
        resp, err := http.Get(final)
        if err != nil {
            return nil, err
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }
        if err := json.Unmarshal([]byte(body), &user); err != nil {
            return nil,err
        }
        return body, nil
    })
    if err != nil {
        return &apierrors.ApiError{
            Message: err.Error(),
            Status:  http.StatusInternalServerError,
        }
    }
    return nil
}

