package requestlog

import (
	"cloud-native/util/logger"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Handler struct {
	handler http.Handler
	logger  *logger.Logger
}

func NewHandler(h http.HandlerFunc, l *logger.Logger) *Handler {
	return &Handler{handler: h, logger: l}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	logEntry := &logEntry{
		ReceivedTime:       start,
		RequestMethod:      r.Method,
		RequestURL:         r.URL.String(),
		RequestHeaderSize:  headerSize(r.Header),
		RequestBodySize:    0,
		UserAgent:          r.UserAgent(),
		Referer:            r.Referer(),
		Proto:              r.Proto,
		RemoteIP:           ipFromHostPort(r.RemoteAddr),
		ServerIP:           "",
		Status:             0,
		ResponseHeaderSize: 0,
		ResponseBodySize:   0,
		Latency:            0,
	}

	if addr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
		logEntry.ServerIP = ipFromHostPort(addr.String())
	}
	newRequest := new(http.Request)
	*newRequest = *r
	rcc := &readCounterCloser{r: r.Body}
	newRequest.Body = rcc
	newWriter := &responseStats{w: w}

	h.handler.ServeHTTP(newWriter, newRequest)

	logEntry.Latency = time.Since(start)
	if rcc.err == nil && rcc.r != nil {
		io.Copy(ioutil.Discard, rcc)
	}

	logEntry.RequestBodySize = rcc.n
	logEntry.Status = newWriter.code
	if logEntry.Status == 0 {
		logEntry.Status = http.StatusOK
	}
	logEntry.ResponseHeaderSize, logEntry.ResponseBodySize = newWriter.size()

	h.logger.Info().
		Time("received_time", logEntry.ReceivedTime).
		Str("method", logEntry.RequestMethod).
		Str("url", logEntry.RequestURL).
		Int64("header_size", logEntry.RequestHeaderSize).
		Int64("body_size", logEntry.RequestBodySize).
		Str("agent", logEntry.UserAgent).
		Str("referer", logEntry.Referer).
		Str("proto", logEntry.Proto).
		Str("remote_ip", logEntry.RemoteIP).
		Str("server_ip", logEntry.ServerIP).
		Int("status", logEntry.Status).
		Int64("resp_header_size", logEntry.ResponseHeaderSize).
		Int64("resp_body_size", logEntry.ResponseBodySize).
		Dur("latency", logEntry.Latency).
		Msg("")
}
