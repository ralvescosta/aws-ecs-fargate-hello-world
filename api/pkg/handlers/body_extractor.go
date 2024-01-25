package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/pkg/views"
)

func ExtractBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	var body T
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(views.UnformattedBody().ToBuffer())
		return nil, err
	}

	err, message := validate(w, body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(views.BadRequest(message).ToBuffer())

		return nil, err
	}

	return &body, nil
}

func validate(w http.ResponseWriter, body any) (error, *map[string]string) {
	val := validator.New()
	err := val.Struct(body)

	if err == nil {
		return nil, nil
	}

	validationErrors := err.(validator.ValidationErrors)
	messages := make(map[string]string, len(validationErrors))

	for _, validationErr := range validationErrors {
		message := ""
		if validationErr.Tag() == "required" {
			message = fmt.Sprintf("%s is required", validationErr.Field())
		} else {
			message = fmt.Sprintf("%s invalid %s", validationErr.Field(), validationErr.Tag())
		}
		messages[validationErr.Field()] = message
	}

	return err, &messages
}
