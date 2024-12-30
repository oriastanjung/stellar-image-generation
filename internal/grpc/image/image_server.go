package image_server

import (
	"context"
	"fmt"

	services "github.com/oriastanjung/stellar/internal/services/image" // your existing usecase package
	pb "github.com/oriastanjung/stellar/proto/image"                   // generated from image_service.proto
)

// imageServer implements pb.ImageServiceServer
type imageServer struct {
	pb.ImageServiceServer
	imageService services.ImageService
}

// NewImageServer constructs our server with the needed service(s).
func NewImageServer(imageService services.ImageService) pb.ImageServiceServer {
	return &imageServer{
		imageService: imageService,
	}
}

// GenerateImage is our gRPC method that constructs the prompt string,
// calls the usecase, and returns the result.
func (s *imageServer) GenerateImage(ctx context.Context, req *pb.ImageRequest) (*pb.ImageResponse, error) {
	// Build the final prompt string (mirroring your original approach)
	prompt := fmt.Sprintf(`Core Subject: %s
			Key Descriptors: %s
			Environment: %s
			Style: %s
			Mood/Tone: %s
			Composition: %s
			Additional Instructions: %s`,
		req.CoreSubject,
		req.KeyDescriptors,
		req.Environment,
		req.Style,
		req.MoodTone,
		req.Composition,
		req.AdditionalInstructions,
	)

	imageURL, filename, err := s.imageService.GenerateImage(ctx, prompt)
	if err != nil {
		return &pb.ImageResponse{
			ImageUrl: "",
			Filename: "",
			Error:    err.Error(),
		}, nil
	}

	return &pb.ImageResponse{
		ImageUrl: imageURL,
		Filename: filename,
		Error:    "",
	}, nil
}

// DownloadAndSaveImage uses the image URL and filename to download locally.
func (s *imageServer) DownloadAndSaveImage(ctx context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	err := s.imageService.DownloadAndSaveImages(ctx, req.GetImageUrl(), req.GetFilename())
	if err != nil {
		return &pb.DownloadResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	return &pb.DownloadResponse{
		Success: true,
		Error:   "",
	}, nil
}
