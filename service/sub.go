package service

import "restful/domain/service_model"

func SubService(req service_model.SubSvcReq) (service_model.SubSvcRes, error) {
	return service_model.SubSvcRes{Result: req.Num1 - req.Num2}, nil
}
