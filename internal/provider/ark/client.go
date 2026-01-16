package ark

import (
	"os"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
)

type Client struct {
	aiClient *arkruntime.Client
}

func NewArkClient() *Client {
	client := arkruntime.NewClientWithApiKey(
		os.Getenv("ARK_API_KEY"),
		arkruntime.WithBaseUrl("https://ark.cn-beijing.volces.com/api/v3"),
	)
	return &Client{
		aiClient: client,
	}
}
