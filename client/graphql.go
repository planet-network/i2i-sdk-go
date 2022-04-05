package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/machinebox/graphql"
)

const (
	i2iTokenHeader = "i2iToken"
)

type query struct {
	filePath  string
	query     string
	variables map[string]interface{}
	timeout   time.Duration
	response  interface{}
}

type GraphqlResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []struct {
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors"`
}

func (c *Client) query(query *query) (*GraphqlResponse, error) {
	var (
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		address       = c.nodeGraphqlAddress()
		graphQlClient = graphql.NewClient(address, graphql.WithHTTPClient(client))
	)

	req := graphql.NewRequest(query.query)
	req.Header.Add(i2iTokenHeader, c.acl)

	for k, v := range query.variables {
		req.Var(k, v)
	}

	if c.debug {
		graphQlClient.Log = func(s string) { fmt.Println(s) }
	}

	ctx, cancel := context.WithTimeout(context.Background(), query.timeout)
	defer cancel()

	if err := graphQlClient.Run(ctx, req, query.response); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) queryUpload(query *query) (*GraphqlResponse, error) {
	var (
		address       = c.nodeGraphqlAddress()
		graphQlClient = graphql.NewClient(address, graphql.UseMultipartForm())
	)

	req := graphql.NewRequest(query.query)
	req.Header.Add(i2iTokenHeader, c.acl)

	for k, v := range query.variables {
		req.Var(k, v)
	}

	if c.debug {
		graphQlClient.Log = func(s string) { fmt.Println(s) }
	}

	ctx, cancel := context.WithTimeout(context.Background(), query.timeout)
	defer cancel()

	if err := graphQlClient.Run(ctx, req, query.response); err != nil {
		return nil, err
	}

	return nil, nil
}
