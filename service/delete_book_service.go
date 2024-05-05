package service

import (
	"context"
	"restful/domain/service_model"
	"restful/entity"
	"restful/repository"
)

func DeleteBookService(ctx context.Context, req service_model.DeleteBookSvcReq) (service_model.DeleteBookSvcRes, error) {
	err := repository.DeleteBook(ctx, entity.BookEntity{
		Id: req.Id,
	})
	if err != nil {
		return service_model.DeleteBookSvcRes{}, err
	}

	return service_model.DeleteBookSvcRes{}, nil
}
