package worker

import (
	"context"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/theoden9014/exercitu/utils"
)

const (
	DefaultConnections = 10000
	DefaultKeepAlive   = 30 * time.Second
	DefaultTimeout     = 30 * time.Second
)

var (
	DefaultLocalAddr = net.IPAddr{IP: net.IPv4zero}
)

type HTTP struct {
	// *Common

	url    *url.URL
	client *http.Client
	dialer *net.Dialer
	log    *log.Logger
}

func NewHTTP() Worker {
	h := &HTTP{}

	h.dialer = &net.Dialer{
		LocalAddr: &net.TCPAddr{IP: DefaultLocalAddr.IP, Zone: DefaultLocalAddr.Zone},
		KeepAlive: DefaultKeepAlive,
		Timeout:   DefaultTimeout,
	}

	h.client = &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           h.dialer.DialContext,
			ResponseHeaderTimeout: DefaultTimeout,
			MaxIdleConnsPerHost:   DefaultConnections,
		},
	}

	return h
}

func (h *HTTP) Configure(f *flag.FlagSet) {
}

func (h HTTP) Try(ctx context.Context) error {
	// h.log.Debugf("Get %s", h.Target)

	resp, err := h.client.Get(h.url.String())
	if resp != nil {
		defer resp.Body.Close()
		// defer h.log.Debugf("Close %s", h.Target)
	}
	if err != nil {
		return err
	}

	return nil
}

func (h *HTTP) Run(ctx context.Context) (*Result, error) {
	str := make([]string, 10)
	utils.GenerateRandString(str)
	urlPath := strings.Join(str, "")
	getUrl := joinUrl(h.url.String(), urlPath)

	// h.log.Debugf("Get %s", getUrl)
	req, err := h.newRequest(ctx, "GET", getUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
		// defer h.log.Debugf("Close %s", getUrl)
	}
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *HTTP) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *h.url
	u.Path = path.Join(h.url.Path, spath)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

func joinUrl(rawurl string, filepath string) string {
	u, err := url.Parse(rawurl)
	if err != nil {
		return rawurl
	}
	u.Path = path.Join(u.Path, filepath)

	return u.String()
}
