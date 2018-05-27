package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/aikizoku/go-gae-template/src/model"

	"github.com/mjibson/goon"

	"github.com/aikizoku/go-gae-template/src/infrastructure"
	"google.golang.org/appengine/log"
)

type sample struct {
	http *infrastructure.HTTP
	csql *sql.DB
}

func (s *sample) Hoge(ctx context.Context) {
	log.Debugf(ctx, "call repository hoge")

}

func (s *sample) testDatastore(ctx context.Context) {
	now := time.Now().Unix()

	client := goon.FromContext(ctx)

	v := &model.Sample{
		ID:        123,
		Name:      "広瀬",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// key := client.Key(v)
	// log.Infof(ctx, "goon.Key => %v", key)

	key, err := client.Put(v)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return
	}
	log.Infof(ctx, "%v", key)

	err = client.Get(&model.Sample{ID: 123})
	if err != nil {
		log.Errorf(ctx, err.Error())
		return
	}
	log.Infof(ctx, "%v", key)
}

// NewSample ...
func NewSample(http *infrastructure.HTTP, csql *sql.DB) Sample {
	return &sample{
		http: http,
		csql: csql,
	}
}
