package models

/*RespuestaConsultaRelacion tiene el true o false que se obtiene de consultar la realcion entre dos usuarios*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
