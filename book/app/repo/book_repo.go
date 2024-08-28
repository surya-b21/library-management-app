package repo

import "context"

type BookRepository struct{}

func (b *BookRepository) Create(ctx context.Context) {}

func (b *BookRepository) Get(ctx context.Context) {}

func (b *BookRepository) GetAll(ctx context.Context) {}

func (b *BookRepository) Update(ctx context.Context) {}

func (b *BookRepository) Delete(ctx context.Context) {}
