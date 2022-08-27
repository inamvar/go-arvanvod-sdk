package arvanvod

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) NewFileUpload(ctx context.Context, channel string, length int64, meta map[string]string) (string, error) {

	requestURL := fmt.Sprintf("%s/channels/%s/files", c.options.BaseUrl, channel)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return "", err
	}
	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	req.Header.Add("tus-resumable", "1.0.0")
	req.Header.Add("upload_length", fmt.Sprintf("%d", length))

	if meta != nil {
		metaData := make([]string, len(meta))
		for k, v := range meta {
			encodedValue := base64.StdEncoding.EncodeToString([]byte(v))
			metaData = append(metaData, fmt.Sprintf("%s %s", k, encodedValue))
		}
		req.Header.Add("upload-metadata", strings.Join(metaData, ","))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return "", err
	}

	location := res.Header.Get("Location")
	return location, nil
}

func (c *Client) GetUploadOffset(ctx context.Context, channel, file string) (offset int, length int, err error) {

	requestURL := fmt.Sprintf("%s/channels/%s/files/%s", c.options.BaseUrl, channel, file)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, requestURL, nil)
	if err != nil {
		return -1, -1, err
	}
	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	req.Header.Add("tus-resumable", "1.0.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, -1, err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return -1, -1, err
	}

	uploadOffset := res.Header.Get("Upload-Offset")

	offset, err = strconv.Atoi(uploadOffset)
	if err != nil {
		return -1, -1, err
	}
	uploadLength := res.Header.Get("Upload-Length")
	length, err = strconv.Atoi(uploadLength)
	if err != nil {
		return -1, -1, err
	}
	return offset, length, nil
}

func (c *Client) UlpoadFileBytes(ctx context.Context, channel, file string, data []byte) (int, error) {
	requestURL := fmt.Sprintf("%s/channels/%s/files/%s", c.options.BaseUrl, channel, file)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, requestURL, bytes.NewReader(data))
	if err != nil {
		return -1, err
	}
	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	req.Header.Add("tus-resumable", "1.0.0")
	req.Header.Add("Content-Type", "application/offset+octet-stream")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return -1, err
	}

	uploadOffset := res.Header.Get("Upload-Offset")

	offset, err := strconv.Atoi(uploadOffset)
	if err != nil {
		return -1, err
	}
	return offset, nil
}
