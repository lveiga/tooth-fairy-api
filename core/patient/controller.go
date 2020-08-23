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

// ListPatients godoc
// @Summary List patients
// @Description get patients
// @Accept  json
// @Produce  json
// @Success 200 {array} patient.Patient
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} http.HTTPError
// @Failure 404 {object} http.HTTPError
// @Failure 500 {object} http.HTTPError
// @Router /patients [get]
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

// CreatePatient godoc
// @Summary Create a new patient
// @Description create patient with body patient json
// @Accept  json
// @Produce  json
// @Param patients body patient.Patient true "Patient"
// @Success 200 {object} patient.Patient
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} http.HTTPError
// @Failure 404 {object} http.HTTPError
// @Failure 500 {object} http.HTTPError
// @Router /patients [post]
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

// ShowPatient godoc
// @Summary Show a patient
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} patient.Patient
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} http.HTTPError
// @Failure 404 {object} http.HTTPError
// @Failure 500 {object} http.HTTPError
// @Router /patients/{id} [get]
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

// UpdatePatient godoc
// @Summary Update a patient
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Param patients body patient.Patient true "Patient"
// @Success 200 {object} patient.Patient
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} http.HTTPError
// @Failure 404 {object} http.HTTPError
// @Failure 500 {object} http.HTTPError
// @Router /patients/{id} [put]
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

// DeletePatient godoc
// @Summary Delete a patient
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} patient.Patient
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} http.HTTPError
// @Failure 404 {object} http.HTTPError
// @Failure 500 {object} http.HTTPError
// @Router /patients/{id} [delete]
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
