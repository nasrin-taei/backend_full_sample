package service

import (
	"restful/domain/service_model"
)

func DivService(req service_model.DivSvcReq) (service_model.DivSvcRes, error) {
	return service_model.DivSvcRes{Result: req.Num1 / req.Num2}, nil
}
