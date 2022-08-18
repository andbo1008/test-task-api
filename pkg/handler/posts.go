package handler

import (
	"boosters/api/post"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Hello(c echo.Context) error {
	return c.JSON(200, "Hello")
}

// @Summary Create Post
// @Tags post
// @Description Create our post
// @ID create-post
// @Accept  json
// @Produce  json
// @Param input body post.Post true "post info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /posts [post]
func (h *Handler) CreatePost(c echo.Context) error {
	var post post.Post
	if err := c.Bind(&post); err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info("Create post: ", post)
	id, err := h.service.CreatePost(post)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []post.Post `json:"data"`
}

// @Summary Get All Post
// @Tags post
// @Description Get all our posts
// @ID getall-post
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /posts [get]
func (h *Handler) GetAllPosts(c echo.Context) error {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, getAllListsResponse{
		Data: posts,
	})
}

// @Summary Get Post by id
// @Tags post
// @Description Get post by id
// @ID getbyid-post
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /posts/{id} [get]
func (h *Handler) GetById(c echo.Context) error {
	id := c.Param("id")
	post, err := h.service.GetById(id)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, post)
}

// @Summary Update Post
// @Tags post
// @Description Update Post by id
// @ID update-post
// @Accept  json
// @Produce  json
// @Param input body post.Post true "post info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /posts/{id} [put]
func (h *Handler) UpdatePost(c echo.Context) error {
	var post post.Post
	id := c.Param("id")
	if err := c.Bind(&post); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := h.service.UpdatePost(post, id); err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, statusResponse{"updated"})
}

// @Summary Delete Post
// @Tags post
// @Description Delete our post
// @ID create-post
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /posts/{id} [delete]
func (h *Handler) DeletePost(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeletePost(id); err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, statusResponse{"deleted"})
}
