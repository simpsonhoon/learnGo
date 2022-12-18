package controller

// /controller.go : 실제 비지니스 로직 및 프로세스가 처리후 결과 전송
import (
	"lecture/go-mvc/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

func (p *Controller) RespOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

// router 에서 수행될 함수
func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
	return
}

func (p *Controller) GetData(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok, I will get data"})
	p.md.GetPerson() // 목록 가져오자..
	return
}
