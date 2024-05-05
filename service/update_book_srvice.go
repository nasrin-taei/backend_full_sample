package service

import (
	"context"
	"restful/domain/service_model"
	"restful/entity"
	"restful/repository"
)

func UpdateBookService(ctx context.Context, req service_model.UpdateBookSvcReq) (service_model.UpdateBookSvcRes, error) {

	err := repository.UpdateBook(ctx, entity.BookEntity{
		Title:     req.Title,
		Count:     req.Count,
		UnitPrice: req.UnitPrice,
		Id:        req.Id,
	})
	if err != nil {
		return service_model.UpdateBookSvcRes{}, err
	}

	return service_model.UpdateBookSvcRes{}, nil
}
