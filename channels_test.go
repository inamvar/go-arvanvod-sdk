package arvanvod_test

import (
	"context"
	"testing"

	"github.com/inamvar/go-arvanvod-sdk"
	"github.com/stretchr/testify/assert"
)

func TestGetChannels(t *testing.T) {
	client := getClient()

	result, err := client.GetChannels(context.Background(), "", 1, 10)

	if assert.NoError(t, err) {
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.Meta.CurrentPage)
		assert.Equal(t, "10", result.Meta.PerPage)
	}

}

func TestCreateAndDeleteChannel(t *testing.T) {
	client := getClient()

	channel := arvanvod.CreateUpdateChannelModel{
		Title:             "test",
		Description:       "some description",
		SecureLinkEnabled: 0,
		SecureLinkKey:     "",
		SecureLinkWithIp:  false,
		AdsEnabled:        0,
		PresentType:       arvanvod.ChannelPresentNone,
		CampaignId:        "",
	}

	result, err := client.CreateChannel(context.Background(), channel)
	if assert.NoError(t, err) {
		assert.NotNil(t, result)
		assert.NotNil(t, result.Data)
		assert.NotEmpty(t, result.Data.Id)
		assert.NotEmpty(t, result.Data.Title)

		err = client.DeleteChannel(context.Background(), result.Data.Id)
		assert.Nil(t, err)
	}

}

func TestGetChannel(t *testing.T) {
	channelId := "19fee3b0-a850-4fa6-bfb0-9563a19c811f"
	client := getClient()

	result, err := client.GetChannel(context.Background(), channelId)
	if assert.NoError(t, err) {
		assert.NotNil(t, result)
		assert.NotNil(t, result.Data)
		assert.NotEmpty(t, result.Data.Id)
		assert.NotEmpty(t, result.Data.Title)
	}
}

func TestUpdateChannel(t *testing.T) {
	channelId := "19fee3b0-a850-4fa6-bfb0-9563a19c811f"
	client := getClient()

	result, err := client.GetChannel(context.Background(), channelId)
	if assert.NoError(t, err) {
		assert.NotNil(t, result)
		assert.NotNil(t, result.Data)
		assert.NotEmpty(t, result.Data.Id)
		assert.NotEmpty(t, result.Data.Title)
	}

	channelToUpdate := arvanvod.CreateUpdateChannelModel{
		Title:             "new title",
		Description:       result.Data.Description,
		SecureLinkEnabled: result.Data.SecureLinkEnabled,
		SecureLinkKey:     result.Data.SecureLinkKey,
		SecureLinkWithIp:  result.Data.SecureLinkWithIp,
		AdsEnabled:        result.Data.AdsEnabled,
		PresentType:       result.Data.PresentType,
		CampaignId:        result.Data.CampaignId,
	}

	err = client.UpdateChannel(context.Background(), result.Data.Id, channelToUpdate)
	assert.NoError(t, err)

}
