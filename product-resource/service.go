package main

import (
	"context"
	"net/http"
	"time"

	"github.com/rsfreitas/go-pocket-utils/errors"
	"github.com/rsfreitas/go-pocket-utils/logger"
	"github.com/rsfreitas/go-pocket-utils/response"

	product_resourcev1 "github.com/rsfreitas/product-api-example/protobuf/gen/go/services/product_resource/v1"

	"github.com/rsfreitas/product-api-example/product-resource/internal/database"
)

type Service struct {
	logger *logger.Logger
	errors *errors.Factory
	db     *database.Database
}

func NewService(serviceName string) *Service {
	log := logger.New(logger.Options{
		FixedAttributes: map[string]string{
			"service.name":    serviceName,
			"service.version": "v1.0.0",
		},
	})

	errorFactory := errors.NewFactory(errors.FactoryOptions{
		HideMessageDetails: false,
		ServiceName:        serviceName,
		Logger:             log,
	})

	db, err := database.New(database.NewOptions{
		Port:         3306,
		Username:     "user",
		Password:     "password",
		Hostname:     "0.0.0.0",
		DatabaseName: "example",
		Errors:       errorFactory,
	})
	if err != nil {
		log.Fatal(context.TODO(), "could not initialize database", logger.Error(err))
	}

	return &Service{
		logger: log,
		errors: errorFactory,
		db:     db,
	}
}

func (s *Service) CreateProduct(ctx context.Context, req *product_resourcev1.CreateProductRequest) (*product_resourcev1.CreateProductResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, s.errors.InvalidArgument(err).Submit(ctx)
	}

	now := time.Now()
	p := &product_resourcev1.ProductModel{
		Id:        database.NewID("prd"),
		Name:      req.GetName(),
		CreatedAt: &now,
	}

	if err := s.db.Insert(ctx, "product", p); err != nil {
		return nil, err
	}

	response.SetResponseCode(ctx, http.StatusCreated)
	return &product_resourcev1.CreateProductResponse{
		Product: p.ProtoResponse(),
	}, nil
}

func (s *Service) GetProductByID(ctx context.Context, req *product_resourcev1.GetProductByIDRequest) (*product_resourcev1.GetProductByIDResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, s.errors.InvalidArgument(err).Submit(ctx)
	}

	var p product_resourcev1.ProductModel
	if err := s.db.FindOneByID(ctx, "product", req.GetId(), &p); err != nil {
		return nil, err
	}

	if isProductDeleted(&p) && !req.GetReturnDeleted() {
		return nil, s.errors.NotFound().Submit(ctx)
	}

	return &product_resourcev1.GetProductByIDResponse{
		Product: p.ProtoResponse(),
	}, nil
}

func (s *Service) UpdateProduct(ctx context.Context, req *product_resourcev1.UpdateProductRequest) (*product_resourcev1.UpdateProductResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, s.errors.InvalidArgument(err).Submit(ctx)
	}

	var p product_resourcev1.ProductModel
	if err := s.db.FindOneByID(ctx, "product", req.GetId(), &p); err != nil {
		return nil, err
	}
	if isProductDeleted(&p) {
		return nil, s.errors.NotFound().Submit(ctx)
	}

	now := time.Now()
	p.UpdatedAt = &now
	p.Name = req.GetName()

	if err := s.db.Update(ctx, "product", &p, &p); err != nil {
		return nil, err
	}

	return &product_resourcev1.UpdateProductResponse{
		Product: p.ProtoResponse(),
	}, nil
}

func (s *Service) DeleteProduct(ctx context.Context, req *product_resourcev1.DeleteProductRequest) (*product_resourcev1.DeleteProductResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, s.errors.InvalidArgument(err).Submit(ctx)
	}

	var p product_resourcev1.ProductModel
	if err := s.db.FindOneByID(ctx, "product", req.GetId(), &p); err != nil {
		return nil, err
	}
	if isProductDeleted(&p) {
		return nil, s.errors.NotFound().Submit(ctx)
	}

	now := time.Now()
	p.DeletedAt = &now

	if err := s.db.Update(ctx, "product", &p, &p); err != nil {
		return nil, err
	}

	return &product_resourcev1.DeleteProductResponse{
		Product: p.ProtoResponse(),
	}, nil
}

func isProductDeleted(p *product_resourcev1.ProductModel) bool {
	return p.DeletedAt != nil && !p.DeletedAt.IsZero()
}
