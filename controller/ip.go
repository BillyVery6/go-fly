package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
)

func PostIpblack(c *gin.Context) {
	ip := c.PostForm("ip")
	if ip==""{
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "请输入IP!",
		})
		return
	}
	kefuId, _ := c.Get("kefu_name")
	models.CreateIpblack(ip,kefuId.(string))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加黑名单成功!",
	})
}
