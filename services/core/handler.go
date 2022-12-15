package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
// @Param		 prefix query string false "prefix of words"
// @Param        keyword query string false "keyword to search"
// @Param        language query string false "search word in language"
// @Param		 sort query string false "sort by alphabet asc/ desc, default asc"
// @Success      200 {object} SearchRes
// @Router       /api/v1/words/search [get]
func (h *Handler) Search(c *gin.Context) {
	keyword, _ := c.GetQuery("keyword")
	language, _ := c.GetQuery("language")
	sort, _ := c.GetQuery("sort")
	prefix, _ := c.GetQuery("prefix")

	fmt.Println("Searching word: ", keyword, " ", prefix, " ", language, " ", sort)
	res, err := h.service.Search(keyword, prefix, language, sort)
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

// Upsert godoc
// @Summary      Add or update word
// @Tags         Core
// @Accept       json
// @Produce      json
// @Param        upsertWordReq body UpsertWord true "if id = 0 means create new word, else update word"
// @Success      200 {object} BaseRes
// @Router       /api/v1/words/upsert [post]
func (h *Handler) Upsert(c *gin.Context) {
	var word UpsertWord

	err := c.Bind(&word)
	if err != nil {
		c.JSON(http.StatusBadRequest, BaseRes{
			Message: err.Error(),
		})
		return
	}

	id, err := h.service.Upsert(word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseRes{
		Message: SuccessMessage,
		Data: UpsertWordRes{
			ID: id,
		},
	})
}

// AddPronounce godoc
// @Summary      Add sound of pronunciation
// @Tags         Core
// @Accept			multipart/form-data
// @Produce		json
// @Param			file	formData	file			true	"this is a test file"
// @Param 		id path int true "id of word to upload pronounce"
// @Success      200 {object} BaseRes
// @Router       /api/v1/words/pronounce/:id [post]
func (h *Handler) AddPronounce(c *gin.Context) {
	wordID, _ := strconv.Atoi(c.Param("id"))
	file, _ := c.FormFile("file")

	err := h.service.AddPronounce(wordID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, BaseRes{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseRes{
		Message: SuccessMessage,
	})
	return
}

// Translate godoc
// @Summary      Translate from Kinh to another language
// @Tags         Core
// @Accept	    json
// @Produce		json
// @Param	 	keyword query string true "keyword in Kinh"
// @Param 		language query string true "language want to translate into"
// @Success      200 {object} BaseRes
// @Router       /api/v1/words/translate [get]
func (h *Handler) Translate(c *gin.Context) {
	resp, _ := h.service.Search("", "", "", "")
	c.JSON(http.StatusOK, BaseRes{
		Message: SuccessMessage,
		Data:    resp,
	})
	return
}

// ListWords godoc
// @Summary      List words by id
// @Tags         Core
// @Accept       json
// @Produce      json
// @Param		 ids query string false "ids seperated  by commas, eg 1,2,3,4"
// @Success      200 {object} SearchRes
// @Router       /api/v1/words/list [get]
func (h *Handler) ListWords(c *gin.Context) {
	ids, _ := c.GetQuery("ids")

	ls := strings.Split(ids, ",")
	var listIDs []int
	for _, x := range ls {
		y, _ := strconv.Atoi(x)
		listIDs = append(listIDs, y)
	}

	res, err := h.service.List(listIDs)
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
