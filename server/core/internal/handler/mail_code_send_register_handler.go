package handler

import (
	"net/http"

	"gcloud/core/internal/logic"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MailCodeSendRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMailCodeSendRegisterLogic(r.Context(), svcCtx)
		resp, err := l.MailCodeSendRegister(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
