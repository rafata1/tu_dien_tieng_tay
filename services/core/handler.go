package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

// Search godoc
// @Summary      Search word in language
// @Tags         Core
// @Accept       json
// @Produce      json
// @Param        keyword query string false "keyword to search"
// @Param        language query string false "search word in language"
// @Success      200 {object} SearchRes
// @Router       /api/v1/words/search [get]
func (h *Handler) Search(c *gin.Context) {
	keyword, _ := c.GetQuery("keyword")
	language, _ := c.GetQuery("language")

	res, err := h.service.Search(keyword, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseRes{
		Message: SuccessMessage,
		Data:    res,
	})
}
