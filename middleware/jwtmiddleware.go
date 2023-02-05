package middleware

import (
	"context"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/fahrurben/gopress/internal/user"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type UserClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

func (u UserClaims) Validate(ctx context.Context) error {
	return nil
}

func AdminOnly(next http.Handler) http.Handler {
	keyFunc := func(ctx context.Context) (interface{}, error) {
		// Our token must be signed using this data.
		return []byte(os.Getenv("JWT_AUTH_SECRET")), nil
	}

	userClaims := func() validator.CustomClaims {
		return &UserClaims{}
	}

	// Set up the validator.
	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
		os.Getenv("JWT_AUTH_ISSUER"),
		[]string{os.Getenv("JWT_AUTH_AUDIENCE")},
		validator.WithCustomClaims(userClaims),
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	// Set up the middleware.
	jwtmiddlewareObj := jwtmiddleware.New(jwtValidator.ValidateToken)

	return jwtmiddlewareObj.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if !ok {
			http.Error(w, "failed to get validated claims", http.StatusInternalServerError)
			return
		}

		userClaims := claims.CustomClaims.(*UserClaims)

		if role, _ := strconv.Atoi(userClaims.Type); role != user.TYPE_ADMIN {
			fmt.Printf("%+v\n", claims.CustomClaims)
			http.Error(w, http.StatusText(403), 403)
			return
		}

		next.ServeHTTP(w, r)
	}))
}
