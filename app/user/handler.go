package user

import (
	"fmt"
	"net/http"

	"github.com/yesleymiranda/go-toolkit/web/decoder"

	"github.com/yesleymiranda/go-toolkit/logger"
	"github.com/yesleymiranda/go-toolkit/web"
)

func MakeGetByIDHandler(svc *Service) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := decoder.IDInt64(r)
		if err != nil {
			logger.Info(err.Error())
			_, _ = fmt.Fprintf(w, "error:"+err.Error())
			return
		}

		res, err := svc.GetByID(r.Context(), id)
		if err != nil {
			logger.Info(err.Error())
			_, _ = fmt.Fprintf(w, "error:"+err.Error())
			return
		}

		err = web.EncodeJSON(w, res, http.StatusOK)
		if err != nil {
			logger.Info(err.Error())
			_, _ = fmt.Fprintf(w, "error:"+err.Error())
			return
		}
	}
}
