package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LoggingMiddleware struct {
	logger *logrus.Logger
}

func NewLoggingMiddleware() *LoggingMiddleware {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	return &LoggingMiddleware{
		logger: logger,
	}
}

func (m *LoggingMiddleware) LogRequestResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Simpan waktu mulai
		start := time.Now()

		// Baca request body
		var reqBody []byte
		if c.Request().Body != nil {
			reqBody, _ = io.ReadAll(c.Request().Body)
			// Kembalikan body ke request agar bisa dibaca lagi
			c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		// Buat response recorder
		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)

		// Simpan writer asli dan ganti dengan recorder
		writer := c.Response().Writer
		c.Response().Writer = &responseWriter{Writer: mw, ResponseWriter: writer.(http.ResponseWriter)}

		// Proses request
		err := next(c)

		// Hitung durasi
		duration := time.Since(start)

		// Log request dan response
		logData := map[string]interface{}{
			"timestamp":     start.Format(time.RFC3339),
			"method":        c.Request().Method,
			"url":           c.Request().URL.String(),
			"status":        c.Response().Status,
			"duration_ms":   duration.Milliseconds(),
			"request_body":  string(reqBody),
			"response_body": resBody.String(),
			"user_agent":    c.Request().UserAgent(),
			"client_ip":     c.RealIP(),
		}

		// Jika ada error
		if err != nil {
			logData["error"] = err.Error()
		}

		// Log berdasarkan status code
		if c.Response().Status >= 500 {
			m.logger.WithFields(logData).Error("Server Error")
		} else if c.Response().Status >= 400 {
			m.logger.WithFields(logData).Warn("Client Error")
		} else {
			m.logger.WithFields(logData).Info("Request Processed")
		}

		return err
	}
}

// responseWriter adalah wrapper untuk menangkap response body
type responseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}