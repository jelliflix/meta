package meta

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Cinemeta struct {
	opts Options
}

type Options struct {
	URL string

	Timeout time.Duration
}

func NewCinemeta(opts Options) *Cinemeta {
	return &Cinemeta{opts: opts}
}

var DefaultOptions = Options{
	Timeout: 10 * time.Second,
	URL:     "https://cinemeta-live.strem.io/meta/",
}

func (c *Cinemeta) request(endpoint string) (reader io.ReadCloser, err error) {
	URL, err := url.Parse(c.opts.URL)
	if err != nil {
		return
	}

	URL.Path += endpoint

	client := &http.Client{Timeout: c.opts.Timeout}
	resp, err := client.Get(URL.String())
	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusNotFound {
		return reader, fmt.Errorf("not found")
	} else if resp.StatusCode != http.StatusOK {
		return reader, fmt.Errorf("got http error %q", resp.Status)
	}

	return resp.Body, err
}

func (c *Cinemeta) requestMeta(endpoint, id string) (meta Meta, err error) {
	resp, err := c.request(fmt.Sprintf("%s/%s.json", endpoint, id))
	if err != nil {
		return
	}

	defer func() {
		_ = resp.Close()
	}()

	dec := json.NewDecoder(resp)
	err = dec.Decode(&meta)
	if err != nil {
		return
	}

	return
}

func (c *Cinemeta) GetMovie(id string) (Meta, error) {
	return c.requestMeta("movie", id)
}

func (c *Cinemeta) GetSeries(id string) (Meta, error) {
	return c.requestMeta("series", id)
}
