package models

import (
	"errors"
	"strings"
)

type Book struct {
	ID     uint64 `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Gender string `json:"gender,omitempty"`
	UserID uint64 `json:"usuario_id,omitempty"`
}

func (b *Book) Prepare() error {
	if err := b.validate(); err != nil {
		return err
	}
	b.format()
	return nil
}

func (b *Book) validate() error {
	if b.Title == "" {
		return errors.New("O título é obrigatório")
	}
	if b.Author == "" {
		return errors.New("O autor é obrigatório")
	}
	if b.Gender == "" {
		return errors.New("O gênero é obrigatório")
	}
	return nil
}

func (b *Book) format() {
	b.Title = strings.TrimSpace(b.Title)
	b.Author = strings.TrimSpace(b.Author)
	b.Gender = strings.TrimSpace(b.Gender)
}
