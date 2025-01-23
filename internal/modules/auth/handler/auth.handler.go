package handler

import (
	"encoding/json"
	"net/http"
	"vibe-user/internal/config"
	"vibe-user/internal/modules/user/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type AuthHandler interface {
	GoogleLogin(c echo.Context) error
	GoogleCallback(c echo.Context) error
}

type authHandler struct {
	userService service.UserService
	oauthConfig *oauth2.Config
	userinfoURL string
}

func NewAuthHandler(userService service.UserService, oauthConfig *config.Oauth) AuthHandler {
	return &authHandler{userService, &oauth2.Config{
		ClientID:     oauthConfig.Google.ClientID,
		ClientSecret: oauthConfig.Google.ClientSecret,
		RedirectURL:  oauthConfig.Google.RedirectURL,
		Scopes:       oauthConfig.Google.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  oauthConfig.Google.Endpoint.AuthURL,
			TokenURL: oauthConfig.Google.Endpoint.TokenURL,
		},
	}, oauthConfig.Google.Endpoint.UserinfoURL}
}

func (h *authHandler) GoogleLogin(c echo.Context) error {
	url := h.oauthConfig.AuthCodeURL(h.randomState(), oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *authHandler) GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	reqCtx := c.Request().Context()
	token, err := h.oauthConfig.Exchange(reqCtx, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	client := h.oauthConfig.Client(reqCtx, token)
	resp, err := client.Get(h.userinfoURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userInfo)
}

func (h *authHandler) randomState() string {
	return uuid.New().String()
}
