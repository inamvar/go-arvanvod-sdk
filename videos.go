package arvanvod

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VideoConvertMode string

var (
	VideoConvertModeAuto    VideoConvertMode = "auto"
	VideoConvertModeManual  VideoConvertMode = "manual"
	VideoConvertModeProfile VideoConvertMode = "profile"
)

type WatermarkArea string

var (
	WatermarkCenter             WatermarkArea = "center"
	WatermarkFixTopLeft         WatermarkArea = "fix_top_left"
	WatermarkFixTopRight        WatermarkArea = "fix_top_right"
	WatermarkFixTopCenter       WatermarkArea = "fix_top_center"
	WatermarkFixBottomLeft      WatermarkArea = "fix_bottom_left"
	WatermarkFixBottomRight     WatermarkArea = "fix_bottom_right"
	WatermarkFixBottomCenter    WatermarkArea = "fix_bottom_center"
	WatermarkAnimateLeftToRight WatermarkArea = "animate_left_to_right"
	WatermarkAnimateTopToBottom WatermarkArea = "animate_top_to_bottom"
)

type SaveVideoReq struct {
	Title           string             `json:"title"`
	ConvertMode     VideoConvertMode   `json:"convert_mode"`
	Description     string             `json:"description,omitempty"`
	VideoUrl        string             `json:"video_url,omitempty"`
	FileId          string             `json:"file_id,omitempty"`
	ParallelConvert bool               `json:"parallel_convert"`
	ThumbnailTime   int                `json:"thumbnail_time,omitempty"`
	WatermarkId     string             `json:"watermark_id,omitempty"`
	WatermarkArea   WatermarkArea      `json:"watermark_area,omitempty"`
	ConvertInfo     []VideoConvertInfo `json:"convert_info,omitempty"`
	Options         []VideoOption      `json:"options,omitempty"`
}

//TODO: define SaveVideoResp properties
type SaveVideoResp struct {
}

type VideoOption struct {
	Bframe              int    `json:"bframe"`
	Level               string `json:"level"`
	Cabac               bool   `json:"cabac"`
	Crf                 int    `json:"crf"`
	MinGop              int    `json:"minGop"`
	MinKeyframeInterval int    `json:"minKeyframeInterval"`
	BitrateTolerance    string `json:"bitrate_tolerance"`
	Fps                 int    `json:"fps"`
	Profile             string `json:"profile"`
}

type VideoConvertInfo struct {
	AudioBitrate int    `json:"audio_bitrate"`
	VideoBitrate int    `json:"video_bitrate"`
	Resolution   string `json:"resolution"`
}

//TODO: define GetAllVideosResp properties
type GetAllVideosResp struct {
}

// store newly created video
// FileId should be string and it will be required whenever video_url is not available) required
// Convert_mode (could be auto or manual or profile)
// Profile_id (required if convert mode has set to profile)
// Parallel_convert (boolean) required
// Thumbnail_time (numeric) required
// Convert_info (must be an array and this will be required if convert mode has set to manual)
// Watermark_area (should be one of: center, fix_top_left, fix_top_right, fix_top_center, fix_bottom_left, fix_bottom_right, fix_bottom_center, animate_left_to_right, animate_top_to_bottom)
func (c *Client) SaveVideo(ctx context.Context, channel string, model *SaveVideoReq) (*SaveVideoResp, error) {
	jsonBody, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("%s/channels/%s/videos", c.options.BaseUrl, channel)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(SaveVideoResp)
	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Get all channel's videos
func (c *Client) GetAllChannelVideos(ctx context.Context, channel, filter string, page, perPage int) (*GetAllVideosResp, error) {
	requestURL := fmt.Sprintf("%s/channels/%s/videos?filter=%s&page=%d&per_page=%d", c.options.BaseUrl, channel, filter, page, perPage)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s", c.options.ApiKey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = getErrorByStatus(res.StatusCode)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(GetAllVideosResp)
	err = json.Unmarshal(resBody, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
