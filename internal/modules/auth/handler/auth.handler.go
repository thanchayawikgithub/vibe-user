package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"vibe-user/internal/config"
	"vibe-user/internal/modules/auth/model"
	"vibe-user/internal/modules/auth/service"
	"vibe-user/internal/modules/user/entity"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type AuthHandler interface {
	GoogleLogin(c echo.Context) error
	GoogleCallback(c echo.Context) error
}

type authHandler struct {
	authService service.AuthService
	oauthConfig *oauth2.Config
	userinfoURL string
}

func NewAuthHandler(authService service.AuthService, oauthConfig *config.Oauth) AuthHandler {
	return &authHandler{authService, &oauth2.Config{
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
	userInfo, err := h.getGoogleUserInfo(client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(userInfo)
	user := &entity.User{
		ID:        userInfo.ID,
		Email:     userInfo.Email,
		FirstName: userInfo.GivenName,
		LastName:  userInfo.FamilyName,
	}

	loginUser, err := h.authService.Login(reqCtx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, loginUser)
}

func (h *authHandler) randomState() string {
	return uuid.New().String()
}

func (h *authHandler) getGoogleUserInfo(client *http.Client) (*model.GoogleUserInfo, error) {
	res, err := client.Get(h.userinfoURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	userInfoBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	userInfo := new(model.GoogleUserInfo)
	if err := json.Unmarshal(userInfoBytes, &userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
