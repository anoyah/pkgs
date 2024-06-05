package main

import (
	"context"
	"encoding/json"
	"log"
	"rstool/rstool"
	"time"
)

type Data struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func main() {
	rs, err := rstool.NewRstool(&rstool.Cfg{
		Network:    "tcp",
		Addr:       "192.168.233.202:6379",
		ClientName: "yother-redis",
		DB:         1,
		PrefixKey:  "yother",
	})
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	var d Data
	rs.Set(context.Background(), "hot-data", "a101000", 0)
	if err := rs.Cache(context.Background(), "data:one", func() (any, time.Duration, error) {
		b, err := json.Marshal(&Data{
			Name:  "Yother",
			Age:   26,
			Hobby: []string{"Code", "Movie", "Sport", "Music"},
		})
		if err != nil {
			return nil, 0, err
		}
		return b, 10 * time.Minute, nil
	}, &d); err != nil {
		panic(err)
	}
	// encoding.BinaryMarshaler
	log.Printf("%+v", d)
}
