package ark

import (
	"context"
	"errors"
	"io"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

const modelName = "ark-visual-v1"

func (c *Client) ImageRecognition(ctx context.Context, prompt string, imageURL string) (string, error) {
	if c == nil || c.aiClient == nil {
		return "", errors.New("client is not initialized")
	}
	inputMessage := &responses.ItemInputMessage{
		Role: responses.MessageRole_user,
		Content: []*responses.ContentItem{
			{
				Union: &responses.ContentItem_Image{
					Image: &responses.ContentItemImage{
						Type:     responses.ContentItemType_input_image,
						ImageUrl: volcengine.String(imageURL),
					},
				},
			},
			{
				Union: &responses.ContentItem_Text{
					Text: &responses.ContentItemText{
						Type: responses.ContentItemType_input_text,
						Text: prompt,
					},
				},
			},
		},
	}
	aiResponse, err := c.aiClient.CreateResponses(ctx, &responses.ResponsesRequest{
		Model: modelName,
		Input: &responses.ResponsesInput{
			Union: &responses.ResponsesInput_ListValue{
				ListValue: &responses.InputItemList{
					ListValue: []*responses.InputItem{
						{
							Union: &responses.InputItem_InputMessage{
								InputMessage: inputMessage,
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

}
