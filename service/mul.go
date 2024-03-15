package service

import "restful/domain/service_model"

func MulService(req service_model.MulSvcReq) (service_model.MulSvcRes, error) {
	return service_model.MulSvcRes{Result: req.Num1 * req.Num2}, nil
}
