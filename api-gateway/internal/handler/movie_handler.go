package handler

import (
	"net/http"
	"strconv"

	"pb"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	client pb.MovieServiceClient
}

func NewMovieHandler(client pb.MovieServiceClient) *MovieHandler {
	return &MovieHandler{client: client}
}

// GetMovie godoc
// @Summary Busca um filme pelo ID
// @Description Retorna os detalhes de um filme específico usando seu ID
// @Tags movies
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} pb.Movie
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	req := &pb.GetMovieRequest{Id: id}
	res, err := h.client.GetMovie(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
		return
	}

	c.JSON(http.StatusOK, res.Movie)
}

// ListMovies godoc
// @Summary Lista todos os filmes
// @Description Retorna uma lista de todos os filmes disponíveis no banco de dados
// @Tags movies
// @Produce json
// @Success 200 {array} pb.Movie
// @Router /movies/ [get]
func (h *MovieHandler) ListMovies(c *gin.Context) {
	req := &pb.ListMoviesRequest{}
	res, err := h.client.ListMovies(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res.Movies == nil {
		c.JSON(http.StatusOK, []pb.Movie{})
		return
	}

	c.JSON(http.StatusOK, res.Movies)
}

// CreateMovie godoc
// @Summary Cria um novo filme
// @Description Cria um novo filme no banco de dados
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body pb.CreateMovieRequest true "Dados do Filme"
// @Success 201 {object} pb.Movie
// @Router /movies/ [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var req pb.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido"})
		return
	}

	res, err := h.client.CreateMovie(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res.Movie)
}

// DeleteMovie godoc
// @Summary Remove um filme
// @Description Remove um filme do banco de dados usando seu ID
// @Tags movies
// @Produce json
// @Param id path int true "Movie ID"
// @Success 204
// @Router /movies/{id} [delete]
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	req := &pb.DeleteMovieRequest{Id: id}
	_, err = h.client.DeleteMovie(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar filme"})
		return
	}

	c.Status(http.StatusNoContent)
}
