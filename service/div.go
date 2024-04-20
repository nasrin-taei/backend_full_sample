package service

import (
	"errors"
	"restful/domain/service_model"
)

func DivService(req service_model.DivSvcReq) (service_model.DivSvcRes, error) {
	if req.Num2 == 0 {
		return service_model.DivSvcRes{}, errors.New("invalid request data")
	}
	return service_model.DivSvcRes{Result: req.Num1 / req.Num2}, nil
}
