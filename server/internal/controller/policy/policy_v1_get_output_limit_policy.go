package policy

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/policy/v1"
)

func (c *ControllerV1) GetOutputLimitPolicy(ctx context.Context, req *v1.GetOutputLimitPolicyReq) (res *v1.GetOutputLimitPolicyRes, err error) {
	var list []entity.OutputLimitRules
	err = dao.OutputLimitRules.Ctx(ctx).OrderAsc(dao.OutputLimitRules.Columns().Position).Scan(&list)
	if err != nil {
		return nil, err
	}

	network, err := service.Network().GetNetwork()
	if err != nil {
		return nil, err
	}

	return &v1.GetOutputLimitPolicyRes{
		Data:      list,
		Network:   *network,
		Total:     len(list),
		Timestamp: time.Now().Unix(),
	}, nil
}
