package grb

type Ciudad struct{
	Name string `json:"name"`
}

type Coneccion struct{
	CityDesde string   `json:"from"`
	CityHasta string   `json:"to"`
	Costo     int      `json:"cost"`
}


