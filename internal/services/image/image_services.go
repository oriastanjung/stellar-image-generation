package services

import (
	"context"

	usecase "github.com/oriastanjung/stellar/internal/usecase/image"
)

// ImageService defines the contract for image-related operations.
type ImageService interface {
	GenerateImage(ctx context.Context, prompt string) (string, string, error)
	DownloadAndSaveImages(ctx context.Context, imageURL, filename string) error
}

// imageService is the concrete implementation of ImageService.
type imageService struct {
	imageUseCase usecase.ImageUseCase
}

// NewImageService returns an instance of ImageService.
func NewImageService(imageUseCase usecase.ImageUseCase) ImageService {
	return &imageService{
		imageUseCase: imageUseCase,
	}
}

// GenerateImage delegates the call to the usecase layer.
func (s *imageService) GenerateImage(ctx context.Context, prompt string) (string, string, error) {
	return s.imageUseCase.GenerateImage(prompt)
}

// DownloadAndSaveImages delegates the call to the usecase layer.
func (s *imageService) DownloadAndSaveImages(ctx context.Context, imageURL, filename string) error {
	return s.imageUseCase.DownloadAndSaveImages(imageURL, filename)
}
