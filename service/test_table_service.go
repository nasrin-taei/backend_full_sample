package service

import (
	"context"
	"restful/domain/service_model"
	"restful/repository"
)

func FetchTestTableRecords(ctx context.Context, req service_model.FetchTableSvcReq) (service_model.FetchTableSvcRes, error) {
	recs, err := repository.FetchAllTestTableRecs(ctx)
	if err != nil {
		return service_model.FetchTableSvcRes{}, err
	}

	resRecords := make([]service_model.TestTableSvcRecord, 0)
	t := service_model.TestTableSvcRecord{}
	for _, r := range recs {
		t = service_model.TestTableSvcRecord{
			Col1: r.Col1,
			Col2: r.Col2,
		}
		resRecords = append(resRecords, t)
	}

	return service_model.FetchTableSvcRes{
		Records: resRecords,
	}, nil
}
