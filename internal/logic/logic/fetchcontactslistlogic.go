package logic

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bot/internal/svc"
	"bot/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchContactsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchContactsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchContactsListLogic {
	return &FetchContactsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchContactsListLogic) FetchContactsList(req *types.FetchContactsListReq) error {
	req2, err := http.NewRequest("POST",
		"http://127.0.0.1:2531/v2/api/contacts/fetchContactsList",
		strings.NewReader(fmt.Sprintf(`{"appId": "%s"}`, l.svcCtx.Config.Name)))
	if err != nil {
		l.Error(err)
		return err
	}
	req2.Header.Add("X-GEWE-TOKEN", "1a856e3f100e48bbaf8ff3cb68bfe3d8")
	req2.Header.Add("Content-Type", "application/json")

	res, err := l.svcCtx.HttpCli.Do(req2)
	if err != nil {
		l.Error(err)
		return err
	}
	defer res.Body.Close()
	//////////////////////////////////////////
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l.Error(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
