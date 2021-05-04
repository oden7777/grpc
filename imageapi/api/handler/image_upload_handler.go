package handler

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/oden7777/grpc/imageapi/api/gen/pb"
)

type ImageUploadHandler struct {
	sync.Mutex
	files map[string][]byte
	pb.UnsafeImageUploadServiceServer
}

func NewImageUploadHandler() *ImageUploadHandler {
	return &ImageUploadHandler{
		files: make(map[string][]byte),
	}
}

func (h *ImageUploadHandler) Upload(stream pb.ImageUploadService_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	meta := req.GetFileMeta()
	filename := meta.Filename

	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := u.String()
	buf := &bytes.Buffer{}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		chunk := r.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}
	}

	data := buf.Bytes()
	mimeType := http.DetectContentType(data)
	h.files[filename] = data

	return stream.SendAndClose(&pb.ImageUploadResponse{
		Uuid:        uuid,
		Size:        int32(len(data)),
		Filename:    filename,
		ContentType: mimeType,
	})
}
