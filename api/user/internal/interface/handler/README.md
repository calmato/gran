# Interface層

HTTPリクエストを受け取り、Application層を使って処理を行い、結果をクライアントに返したり、サーバのログに出力する。  
外部データとの差異を吸収してApplication層に渡し、結果を返却する役割を担うのがInterface層の役割。

## テンプレート

```go
package v1

import (
	"net/http"

	"github.com/calmato/gran/api/sample/internal/application"
	"github.com/calmato/gran/api/sample/internal/application/request"
	"github.com/calmato/gran/api/sample/internal/interface/handler"
	"github.com/gin-gonic/gin"
)

// APIV1SampleHandler - SampleHandlerインターフェース
type APIV1SampleHandler interface {
	Create(ctx *gin.Context)
}

type apiV1SampleHandler struct {
	sampleApplication application.SampleApplication
}

// NewAPIV1SampleHandler - APIV1SampleHandlerの生成
func NewAPIV1SampleHandler(ua application.SampleApplication) APIV1SampleHandler {
	return &apiV1SampleHandler{
		sampleApplication: sa,
	}
}

func (sh *apiV1SampleHandler) Create(ctx *gin.Context) {
	req := request.CreateSample{}
	if err := ctx.BindJSON(&req); err != nil {
		hanlder.ErrorHandling(ctx, err)
		return
	}

	if err := uh.sampleApplication.Create(ctx, &req); err != nil {
		hanlder.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
```
