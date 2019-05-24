package myml

import (
    "github.com/mercadolibre/ejercicioCircuitBreaker/src/api/domain/myml"
    "github.com/mercadolibre/ejercicioCircuitBreaker/src/api/utils/apierrors"
)

func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {
    user := &myml.User{
        ID: userID,
    }
    if apiErr := user.Get(); apiErr != nil {
        return nil, apiErr
    }
    return user, nil
}
