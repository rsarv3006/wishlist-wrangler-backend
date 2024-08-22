package auth

import (
	"testing"
	"time"
	"wishlist-wrangler-api/ent"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestGenerateJWT(t *testing.T) {
	user := &ent.User{
		ID:          uuid.New(),
		DisplayName: "Test User",
		Email:       "test@test.com",
	}

	app := fiber.New()
	fastReq := fasthttp.RequestCtx{}
	ctx := app.AcquireCtx(&fastReq)
	ctx.Locals("JwtSecret", "testSecret")

	token, err := GenerateJWT(user, ctx)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotEmpty(t, *token)
}

func TestValidateToken(t *testing.T) {
	user := &ent.User{
		ID:          uuid.New(),
		DisplayName: "Test User",
		Email:       "test@test.com",
	}

	app := fiber.New()
	fastReq := fasthttp.RequestCtx{}
	ctx := app.AcquireCtx(&fastReq)
	ctx.Locals("JwtSecret", "testSecret")

	token, _ := GenerateJWT(user, ctx)

	t.Run("Valid token", func(t *testing.T) {
		validatedUser, err := ValidateToken(*token, ctx)

		assert.NoError(t, err)
		assert.NotNil(t, validatedUser)
		assert.Equal(t, user.ID, validatedUser.ID)
		assert.Equal(t, user.DisplayName, validatedUser.DisplayName)
	})

	t.Run("Expired token", func(t *testing.T) {
		claims := &JWTClaims{
			User: *user,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(-time.Hour).Unix(),
			},
		}
		expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		expiredTokenString, _ := expiredToken.SignedString([]byte("testSecret"))

		_, err := ValidateToken(expiredTokenString, ctx)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token is expired")
	})

	t.Run("Invalid token", func(t *testing.T) {
		_, err := ValidateToken("invalidToken", ctx)

		assert.Error(t, err)
		assert.NotEqual(t, ErrExpired, err)
		assert.NotEqual(t, ErrInvalid, err)
	})
}

func TestJWTExpiration(t *testing.T) {
	user := &ent.User{
		ID:          uuid.New(),
		DisplayName: "Test User",
		Email:       "test@test.com",
	}

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Locals("JwtSecret", "testSecret")

	token, err := GenerateJWT(user, ctx)
	assert.NoError(t, err)

	claims := &JWTClaims{}
	_, err = jwt.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("testSecret"), nil
	})
	assert.NoError(t, err)

	expectedExpiration := time.Now().Add(SevenDays)
	actualExpiration := time.Unix(claims.ExpiresAt, 0)

	assert.WithinDuration(t, expectedExpiration, actualExpiration, 5*time.Second)
}
