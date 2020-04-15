// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

type connector struct {
	Client    *elasticsearch.Client
	Addresses []string
	Username  string
	Password  string
	Logger    Logger
}

func NewConnectorImpl(opts ClientOptions) Connector {
	if opts.Logger == nil {
		opts.Logger = NewDefaultLogger()
	}
	c := connector{Logger: opts.Logger}
	if err := c.initClient(context.Background()); err != nil {
		c.Logger.Error(context.Background(), fmt.Sprintf("Unable to connect, Error %v", err.Error()))
	}
	return &c
}

func getJsonString(doc Document) (jsonString string, err error) {
	var jsonStringByteArr []byte
	jsonStringByteArr, err = json.Marshal(&doc)
	jsonString = fmt.Sprintf("%s", jsonStringByteArr)
	return
}

func (s *connector) initClient(ctx context.Context) (err error) {
	conf := elasticsearch.Config{
		Addresses: s.Addresses,
		Username:  "",
		Password:  "",
	}
	s.Client, err = elasticsearch.NewClient(conf)
	if err != nil {
		s.Logger.Error(ctx, fmt.Sprintf("Unable to create elastic search client Error : %v", err.Error()))
		return
	}
	return
}

func (s *connector) AddDocument(ctx context.Context, index string, doc Document) (err error) {
	var (
		bodyString string
		res        *esapi.Response
	)
	if bodyString, err = getJsonString(doc); err != nil {
		s.Logger.Error(ctx, fmt.Sprintf("Converting doc o json Error :%v", err.Error()))
		return
	}

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: fmt.Sprintf("%v", doc.Id),
		Body:       strings.NewReader(bodyString),
		Refresh:    "true",
	}

	s.Logger.Trace(ctx, fmt.Sprintf("Idex : %v Index request  : %v", req.Index, req.Body))

	// If client is does not exist,try to connected
	if s.Client == nil {
		if err = s.initClient(ctx); err != nil {
			s.Logger.Error(ctx, fmt.Sprintf("Unable to reconnect, Error %v", err.Error()))
			return err
		}
	}

	res, err = req.Do(ctx, s.Client)

	if err != nil {
		s.Logger.Error(ctx, fmt.Sprintf("Elasticsearch request Error :%v", err.Error()))
		return
	}

	err = res.Body.Close()
	if err != nil {
		s.Logger.Error(ctx, fmt.Sprintf("Elasticsearch request response body closing Error :%v", err.Error()))
	}

	return
}
