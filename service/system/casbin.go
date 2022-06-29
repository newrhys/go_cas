package system

import (
	"errors"
	"strconv"
	"wave-admin/global"
	"wave-admin/model/system/request"
)

type CasbinService struct{}

func (casbinService *CasbinService) UpdateCasbin(roleId uint64, casbinInfos []request.CasbinInfo) error {
	id := strconv.Itoa(int(roleId))
	casbinService.ClearCasbin(0, id)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{id, v.Path, v.Method})
	}
	e := global.GnCasbin
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := global.GnCasbin
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
