package service

import (
	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/1612180/chat_stranger/internal/repository"
)

type FavoriteService struct {
	favRepo repository.FavoriteRepo
}

func NewFavoriteService(favRepo repository.FavoriteRepo) *FavoriteService {
	return &FavoriteService{
		favRepo: favRepo,
	}
}

func (s *FavoriteService) FetchAll() ([]*dtos.FavoriteResponse, []error) {
	favs, errs := s.favRepo.FetchAll()
	if len(errs) != 0 {
		return nil, errs
	}

	var favRess []*dtos.FavoriteResponse
	for _, fav := range favs {
		favRess = append(favRess, fav.ToResponse())
	}

	return favRess, nil
}

func (s *FavoriteService) Find(name string) (*dtos.FavoriteResponse, []error) {
	fav, errs := s.favRepo.Find(name)
	if len(errs) != 0 {
		return nil, errs
	}

	return fav.ToResponse(), nil
}

func (s *FavoriteService) Create(req dtos.FavoriteRequest) (int, []error) {
	fav := (&models.Favorite{}).FromRequest(req)

	return s.favRepo.Create(fav)
}

func (s *FavoriteService) Delete(id int) []error {
	return s.favRepo.Delete(id)
}
