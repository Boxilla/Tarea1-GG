package grb

type Ciudad struct{
	Name string `json:"name"`
}

type Conexion struct{
	CityDesde string   `json:"from"`
	CityHasta string   `json:"to"`
	Costo     int      `json:"cost"`
}


