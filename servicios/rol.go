package servicios


import(
    "database/sql"
    "fmt"
)

func InsertarRol(db *sql.DB, descripcion string, menuInicio bool, menuMiembro bool, menuReporte bool, menuConfiguracion bool, menuInformacion bool, menuUsuario bool, menuReset bool) error {
	_, err := db.Exec("CALL InsertarRol(?, ?, ?, ?, ?, ?, ?, ?)", descripcion, menuInicio, menuMiembro, menuReporte, menuConfiguracion, menuInformacion, menuUsuario, menuReset)
	if err != nil {
		return fmt.Errorf("error insertano rol: %v", err)
	}
	return nil
}

func ActualizarRol(db *sql.DB, idRol int, descripcion string, menuInicio bool, menuMiembro bool, menuReporte bool, menuConfiguracion bool, menuInformacion bool, menuUsuario bool, menuReset bool) error {
	_, err := db.Exec("CALL ActualizarRol(?, ?, ?, ?, ?, ?, ?, ?)", idRol, descripcion, menuInicio, menuMiembro, menuReporte, menuConfiguracion, menuInformacion, menuUsuario, menuReset)
	if err != nil {
		return fmt.Errorf("error actualizando rol: %v", err)
	}
	return nil
}

func EliminarRol(db *sql.DB, idRol int) error {
	_, err := db.Exec("CALL EliminarRol(?)", idRol)
	if err != nil {
		return fmt.Errorf("error eliminando rol: %v", err)
	}
	return nil
}


