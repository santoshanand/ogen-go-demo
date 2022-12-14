package services

import (
	"context"
	"sync"

	petstore "github.com/santoshanand/ogen-go-demo/petstore"
)

// PetsService -
type PetsService struct {
	Pets map[int64]petstore.Pet
	id   int64
	mux  sync.Mutex
}

// AddPet -
func (p *PetsService) AddPet(ctx context.Context, req petstore.Pet) (petstore.Pet, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.Pets[p.id] = req
	p.id++
	return req, nil
}

// DeletePet -
func (p *PetsService) DeletePet(ctx context.Context, params petstore.DeletePetParams) (petstore.DeletePetOK, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	delete(p.Pets, params.PetId)
	return petstore.DeletePetOK{}, nil
}

// GetPetById -
func (p *PetsService) GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet, ok := p.Pets[params.PetId]
	if !ok {
		// Return Not Found.
		return &petstore.GetPetByIdNotFound{}, nil
	}
	return &pet, nil
}

// UpdatePet -
func (p *PetsService) UpdatePet(ctx context.Context, params petstore.UpdatePetParams) (petstore.UpdatePetOK, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet := p.Pets[params.PetId]
	pet.Status = params.Status
	if val, ok := params.Name.Get(); ok {
		pet.Name = val
	}
	p.Pets[params.PetId] = pet

	return petstore.UpdatePetOK{}, nil
}
