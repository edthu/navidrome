package repositories

import (
	"github.com/deluan/gosonic/models"
	"errors"
)

type ArtistIndex interface {
	Put(m *models.ArtistIndex) error
	Get(id string) (*models.ArtistIndex, error)
	GetAll() ([]models.ArtistIndex, error)
}

type ArtistIndexImpl struct {
	BaseRepository
}

func NewArtistIndexRepository() ArtistIndex {
	r := &ArtistIndexImpl{}
	r.init("index", &models.ArtistIndex{})
	return r
}

func (r *ArtistIndexImpl) Put(m *models.ArtistIndex) error {
	if m.Id == "" {
		return errors.New("Id is not set")
	}
	return r.saveOrUpdate(m.Id, m)
}

func (r *ArtistIndexImpl) Get(id string) (*models.ArtistIndex, error) {
	var rec interface{}
	rec, err := r.readEntity(id)
	return rec.(*models.ArtistIndex), err
}

func (r *ArtistIndexImpl) GetAll() ([]models.ArtistIndex, error) {
	var indices = make([]models.ArtistIndex, 0)
	err := r.loadAll(&indices)
	return indices, err
}


