package logstreamer

import (
	"fmt"
	"github.com/buildkite/agent/buildkite/http"
	"github.com/buildkite/agent/buildkite/logger"
	"time"
)

type Chunk struct {
	// The ID of the chunk as assigned by Buildkite
	ID string `json:"id,omitempty`

	// The sequence number of this chunk
	Order int `json:"order"`

	// The contents of the chunk
	Blob string `json:"blob"`

	// The size of the chunk
	Bytes int `json:"bytes"`

	// The HTTP request we'll send logs to
	Request http.Request

	// If this chunk has been uploaded
	Uploaded bool
}

func (chunk *Chunk) Upload() {
	logger.Debug("Uploading %d bytes of content at order %d", chunk.Bytes, chunk.Order)
	logger.Debug("%s", chunk.Request.String())
	chunk.Request.Method = fmt.Sprintf("foo %d", chunk.Order)
	time.Sleep(time.Second * 5)
	logger.Debug("Finished %d", chunk.Order)
}
