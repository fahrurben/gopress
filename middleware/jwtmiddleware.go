package middleware

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type UserClaims struct {
	Email string `json:"auth_email"`
	Name  string `json:"auth_name"`
	Type  string `json:"auth_type"`
}

func (u UserClaims) Validate(ctx context.Context) error {
	return nil
}

func JwtAuthMiddleware(next http.Handler) http.Handler {
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

		usertype, _ := strconv.Atoi(userClaims.Type)
		ctx := context.WithValue(r.Context(), "auth_email", userClaims.Email)
		ctx = context.WithValue(ctx, "auth_name", userClaims.Name)
		ctx = context.WithValue(ctx, "auth_type", usertype)

		next.ServeHTTP(w, r.WithContext(ctx))
	}))
}
