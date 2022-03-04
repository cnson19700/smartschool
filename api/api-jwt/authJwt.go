package api_jwt

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type GinJWTMiddleware struct {
	// Public key file for asymmetric algorithms
	PublicKeyFile string

	// Public key
	publicKey *rsa.PublicKey

	// Private key file for asymmetric algorithms
	PrivateKeyFile string

	// Private key
	privateKey *rsa.PrivateKey
}

func GetDefaultGinJWTMiddleware() (*GinJWTMiddleware, error) {
	mw := &GinJWTMiddleware{
		PublicKeyFile:  "resources/public.pem",
		PrivateKeyFile: "resources/private.pem",
	}

	err := mw.Init()
	if err != nil {
		return nil, err
	}

	return mw, nil
}

func (mw *GinJWTMiddleware) Init() error {
	err := mw.readPublicKeyFile()
	if err != nil {
		return err
	}

	err = mw.readPrivateKeyFile()
	if err != nil {
		return err
	}

	return nil
}

func (mw *GinJWTMiddleware) readPublicKeyFile() error {
	keyData, err := ioutil.ReadFile(mw.PublicKeyFile)
	if err != nil {
		return errors.New("no public key")
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return errors.New("public key is invalid")
	}

	mw.publicKey = key
	return nil
}

func (mw *GinJWTMiddleware) readPrivateKeyFile() error {
	keyData, err := ioutil.ReadFile(mw.PrivateKeyFile)
	if err != nil {
		return errors.New("no private key")
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return errors.New("private key is invalid")
	}

	mw.privateKey = key
	return nil
}

func (mw *GinJWTMiddleware) GetAuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		mw.authImpl(c)
	}
}

func (mw *GinJWTMiddleware) authImpl(c *gin.Context) {
	token, err := mw.parseToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "authorization failed")
		return
	}

	if token != nil {
		claims := token.Claims.(jwt.MapClaims)

		id := claims["id"]
		c.Set("userId", id)
	}

	c.Next()
}

func (mw *GinJWTMiddleware) parseToken(c *gin.Context) (*jwt.Token, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("authorization header is invalid")
	}

	return jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
		return mw.publicKey, nil
	})
}

func (mw *GinJWTMiddleware) GetSignedString(token *jwt.Token) (string, error) {
	return token.SignedString(mw.privateKey)
}
