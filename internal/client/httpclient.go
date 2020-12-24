package client

import (
	"context"
	"devstats/internal/utils"
	"golang.org/x/oauth2"
	"net/http"
)

type HttpClientWrapper struct {
	*http.Client
}

func NewClient() HttpClientWrapper {
	return HttpClientWrapper{
		&http.Client{},
	}
}

func NewClientWithOauth(key string, ctx context.Context) HttpClientWrapper {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: key},
	)
	return HttpClientWrapper{
		oauth2.NewClient(ctx, tokenSource),
	}
}

func (client *HttpClientWrapper) MakeRequest(req *HttpRequestWrapper, postProcessors ...ResponsePostProcessor) error {
	resp, err := client.Do(req.Request)
	if err != nil {
		return err
	}

	for _, postprocessor := range postProcessors {
		err := postprocessor(resp)
		utils.CheckPrintError(err)
	}

	defer resp.Body.Close()
	return nil
}
