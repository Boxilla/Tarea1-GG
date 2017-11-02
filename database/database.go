package database

import (
	"github.com/JanoBotso/Tarea1-GG"
	"errors"
)

type Database struct {
	ciudades		map[string]*grb.Ciudad
	conexiones     map[string]*grb.Conexion
}


func New() *Database {
	return &Database{
		ciudades:   make(map[string]*grb.Ciudad),
		conexiones: make(map[string]*grb.Conexion),
	}
}



func (db *Database) ListaCiudad() []*grb.Ciudad {
	// Rellenamos una lista con los nombres de todas lo contenido en cities
	ciudades := make([]*grb.Ciudad, len(db.ciudades))
	i := 0
	for _, ciudad := range db.ciudades {
		ciudades[i] = ciudad
		i += 1
	}
	return ciudades
}


func (db *Database) ListaConexiones() []*grb.Conexion  {
	// Rellenamos una lista con todas lo contenido en connections
	conexiones := make([]*grb.Conexion, len(db.conexiones))
	i := 0
	for _, conexion := range db.conexiones {
		conexiones[i] = conexion
		i += 1
	}
	return conexiones
}


func (db *Database) AgragarCiudad (ciudad *grb.Ciudad) {
	db.ciudades[ciudad.Name] = ciudad
}
/*


type Database struct {
	entries map[string]*grb.Entry
}

func New() *Database {
	return &Database{
		entries: make(map[string]*grb.Entry),
	}
}

func (db *Database) AddEntry(entry *grb.Entry) {
	db.entries[entry.Key] = entry
}

func (db *Database) EntryList() []*grb.Entry {
	entries := make([]*grb.Entry, len(db.entries))

	var i = 0
	for _, entry := range db.entries {
		entries[i] = entry
		i += 1
	}

	return entries
}

func (db *Database) EntryGet(key string) (*grb.Entry, error) {
	value, ok := db.entries[key]
	if !ok {
		return nil, errors.New("entry not found")
	}
	return value, nil
}

*/