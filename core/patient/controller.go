package patient

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	httptransf "github.com/tooth-fairy/infrastructure/http"
)

type controller struct {
	patientRepository Repository
}

func (c *controller) GetPacients(ctx *gin.Context) {
	var patient Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		httptransf.BadRequest(ctx, err)
		return
	}

	patients, err := c.patientRepository.FindAllPatients()
	if err != nil {
		httptransf.InternalServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, patients)
}

func (c *controller) NewPatient(ctx *gin.Context) {
	var patient Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		httptransf.BadRequest(ctx, err)
		return
	}

	result, err := c.patientRepository.CreatePatient(&patient)
	if err != nil {
		httptransf.InternalServerError(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *controller) GetPatient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		httptransf.BadRequest(ctx, errors.New("parametro id deve ser um número inteiro"))
		return
	}

	patients, err := c.patientRepository.GetPatient(uint32(id))
	if err != nil {
		httptransf.InternalServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, patients)
}

func (c *controller) UpdatePatient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		httptransf.BadRequest(ctx, errors.New("parametro id deve ser um número inteiro"))
		return
	}
	var patient Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		httptransf.BadRequest(ctx, err)
		return
	}

	result, err := c.patientRepository.UpdatePatient(&patient, uint32(id))
	if err != nil {
		httptransf.InternalServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *controller) DeletePatient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		httptransf.BadRequest(ctx, errors.New("parametro id deve ser um número inteiro"))
		return
	}

	result, err := c.patientRepository.DeletePatient(uint32(id))
	if err != nil {
		httptransf.InternalServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func newController(patientRepository Repository) *controller {
	return &controller{
		patientRepository: patientRepository,
	}
}
