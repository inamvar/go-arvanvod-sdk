package arvanvod_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	client := getClient()
	filename := "sample.mp4"
	path := filepath.Join("samples", filename)
	file, err := os.Open(path)

	if assert.NoError(t, err) {
		data, err := ioutil.ReadAll(file)
		assert.Equal(t, nil, err)
		length := len(data)
		channel := "19fee3b0-a850-4fa6-bfb0-9563a19c811f"
		meta := map[string]string{
			"filename": filename,
			"filetype": "video/mp4",
		}
		location, err := client.NewFileUpload(context.Background(), channel, int64(length), meta)
		if assert.NoError(t, err) {
			assert.NotEmpty(t, location)
			ss := strings.Split(location, "/")
			fileId := ss[len(ss)-1]

			step := int64(1000000)
			for {
				offset, len, err := client.GetUploadOffset(context.Background(), channel, fileId)
				assert.Equal(t, nil, err)
				assert.Equal(t, length, len)
				if offset == len || err != nil {
					break
				}
				end := len
				if offset+step < len {
					end = offset + step
				}
				bytes := data[offset:end]
				offset, err = client.UlpoadFileBytes(context.Background(), channel, fileId, bytes)
				assert.Equal(t, nil, err)
				fmt.Printf("\n len: %d, offset: %d", len, offset)
				if offset == len || err != nil {
					break
				}

			}
			fmt.Printf("\n finished!")
		}
	}
}
