package service

import "restful/domain/service_model"

func AddService(req service_model.AddSvcReq) (service_model.AddSvcRes, error) {
	return service_model.AddSvcRes{Result: req.Num1 + req.Num2}, nil
}
