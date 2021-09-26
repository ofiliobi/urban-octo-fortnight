package http

import (
	"context"
	"encoding/json"
	"os"

	"github.com/ofiliobi/urban-octo-fortnight/adapter/logger"
	"github.com/ofiliobi/urban-octo-fortnight/domain/entity"
	"github.com/ofiliobi/urban-octo-fortnight/usecase"
	"github.com/pkg/errors"
)

const (
	autorizado = "Autorizado"
)

var (
	errAuthorizationDenied = errors.New("authorization denied")
)

type (
	authorizer struct {
		client HTTPGetter
		log    logger.Logger
		logKey string
	}

	authorizerResponse struct {
		Message string
	}
)

// NewAuthorizer creates new authorizer with its dependencies
func NewAuthorizer(client HTTPGetter, l logger.Logger) usecase.Authorizer {
	return authorizer{
		client: client,
		log:    l,
		logKey: "send_authorized",
	}
}

// Authorized authorizes a transfer
func (a authorizer) Authorized(_ context.Context, _ entity.Transfer) (bool, error) {
	res, err := a.client.Get(os.Getenv("AUTHORIZER_URI"))
	if err != nil {
		a.log.WithFields(logger.Fields{
			"key":   a.logKey,
			"error": err.Error(),
		}).Errorf("failed to client")

		return false, errAuthorizationDenied
	}

	b := &authorizerResponse{}
	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		a.log.WithFields(logger.Fields{
			"key":   a.logKey,
			"error": err.Error(),
		}).Errorf("failed to marshal message")

		return false, errAuthorizationDenied
	}

	if b.Message != autorizado {
		return false, errAuthorizationDenied
	}

	a.log.WithFields(logger.Fields{
		"key":         a.logKey,
		"http_status": res.StatusCode,
	}).Infof("success to authorized")

	return true, nil
}