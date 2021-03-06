// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	elastic "github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/turbit-nsi/internal/config"
)

type connector struct {
	Addresses []string
	Username  string
	Password  string
	Client    *elastic.Client
}

func NewConnectorImpl(opts ClientOptions) Connector {
	c := connector{
		Addresses: opts.Addresses,
		Username:  opts.Username,
		Password:  opts.Password,
	}
	if err := c.initClient(context.Background()); err != nil {
		log.Error(context.Background(), fmt.Sprintf("Unable to connect, Error %v", err.Error()))
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
	s.Client, err = elastic.NewClient(
		elastic.SetURL(config.ElasticsearchConf.Addresses...),
		elastic.SetErrorLog(NewErrorLogger()),
		elastic.SetInfoLog(NewInfoLogger()),
		elastic.SetTraceLog(NewTraceLogger()),
	)
	return
}

func (s *connector) AddDocument(ctx context.Context, index string, doc Document) (err error) {
	var indexResult *elastic.IndexResponse
	if s.Client == nil {
		if err = s.initClient(ctx); err != nil {
			log.Error(ctx, fmt.Sprintf("Unable to reconnect, Error : %v", err.Error()))
			return err
		}
	}

	indexResult, err = s.Client.Index().
		Index(index).
		BodyJson(&doc).
		Id(fmt.Sprintf("%v", doc.Id)).
		Do(ctx)

	if err != nil {
		log.Error(ctx, fmt.Sprintf("Unable to index the document Doc : %v  Error : %v", doc, err.Error()))
		return
	}
	if indexResult == nil {
		log.Error(ctx, fmt.Sprintf("Expected result to be != nil; got: %v", indexResult))
	}
	return
}

func (s *connector) GetDocuments(ctx context.Context, criteria Criteria) (docs []Document, err error) {
	docs = make([]Document, 0)
	var searchResult *elastic.SearchResult
	if s.Client == nil {
		if err = s.initClient(ctx); err != nil {
			log.Error(ctx, fmt.Sprintf("Unable to reconnect, Error : %v", err.Error()))
			return
		}
	}

	var categoriesStr []interface{}
	for _, v := range criteria.Categories {
		categoriesStr = append(categoriesStr, v)
	}
	var query *elastic.BoolQuery

	if criteria.TextSearchEnable {
		//query = elastic.NewBoolQuery().Must(
		//	elastic.NewMatchQuery("company_name", criteria.SearchTerm),
		//)
		query = elastic.NewBoolQuery().Must(
			elastic.NewMultiMatchQuery(criteria.SearchTerm, "company_name^3", "categories^2", "content"),
		)
	} else {
		query = elastic.NewBoolQuery().Must(
			elastic.NewTermsQuery("geo_hex_ids", strings.Join(criteria.GeoHexId, ",")),
			elastic.NewTermsQuery("categories", categoriesStr...),
			elastic.NewRangeQuery("id").Gt(criteria.LastConsumedId).IncludeLower(true).IncludeUpper(false),
		)
	}

	src, _ := query.Source()
	data, _ := json.MarshalIndent(src, "", " ")
	fmt.Printf("%s", data)
	searchResult, err = s.Client.Search().
		Index(criteria.Index).
		Query(query).
		From(criteria.PageIndex).
		Size(criteria.PageSize).Do(ctx)
	if searchResult != nil {
		if len(searchResult.Hits.Hits) > 0 {
			for _, hit := range searchResult.Hits.Hits {
				doc := Document{}
				sourceStr, _ := hit.Source.MarshalJSON()
				err := json.Unmarshal(sourceStr, &doc)
				if err != nil {
					continue
				}
				docs = append(docs, doc)
			}
		}
	}
	return
}
