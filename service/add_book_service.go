package service

import (
	"context"
	"restful/domain/service_model"
	"restful/entity"
	"restful/repository"
)

func AddBookService(ctx context.Context, req service_model.AddBookSvcReq) (service_model.AddBookSvcRes, error) {
	err := repository.AddBook(ctx, entity.BookEntity{
		Title:     req.Title,
		Count:     req.Count,
		UnitPrice: req.UnitPrice,
	})
	if err != nil {
		return service_model.AddBookSvcRes{}, err
	}
	return service_model.AddBookSvcRes{}, nil
}
