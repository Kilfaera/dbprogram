package servicios

import (
	"database/sql"
	"fmt"
	"time"
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

func InsertarUsuario(db *sql.DB,
	documentoUsuario string,
	nombre string,
	salario float32,
	idRol int,
	correo string,
	contraseña string,
	telefono string,
	telefonoEmergencia string,
	tema string,
	colorenfacis string,
	estado bool) error {
	_, err := db.Exec("CALL InsertarUsuario(?, ?, ?, ?, ?, ?, ?, ?, ?, ?,)", documentoUsuario, nombre, salario, idRol, correo, contraseña, telefono, telefonoEmergencia, tema, colorenfacis, estado)
	if err != nil {
		return fmt.Errorf("error insertando usuario: %v", err)
	}
	return nil
}

func ActualizarUsuario(db *sql.DB, idUsuario int, documentoUsuario string, nombre string, salario float32, idRol int, correo string, contraseña string, telefono string, telefonoEmergencia string, tema string, colorenfacis string, estado bool) error {
	_, err := db.Exec("CALL ActualizarUsuario(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", idUsuario, documentoUsuario, nombre, salario, idRol, correo, contraseña, telefono, telefonoEmergencia, tema, colorenfacis, estado)
	if err != nil {
		return fmt.Errorf("error actualizando usuario: %v", err)
	}
	return nil
}

func EliminarUsuario(db *sql.DB, idUsuario int) error {
	_, err := db.Exec("CALL EliminarUsuario(?)", idUsuario)
	if err != nil {
		return fmt.Errorf("error eliminando usuario: %v", err)
	}
	return nil
}

type UsuarioConRol struct {
	IdUsuario          int
	DocumentoUsuario   string
	Nombre             string
	Salario            float32
	Rol                string
	Correo             string
	Contraseña         string
	Telefono           string
	TelefonoEmergencia string
	Tema               string
	Colorenfacis       string
	Estado             bool
	FechaRegistro      string
}

func ObtenerUsuariosConRoles(db *sql.DB) ([]UsuarioConRol, error) {
	rows, err := db.Query("CALL ObtenerUsuariosConRoles()")
	if err != nil {
		return nil, fmt.Errorf("error ejecutando ObtenerUsuariosConRoles: %v", err)
	}
	defer rows.Close()

	//Slice para almacenar los resultados
	var usuarios []UsuarioConRol

	for rows.Next() {
		var usuario UsuarioConRol
		var estadoByte []byte

		err := rows.Scan(
			&usuario.IdUsuario,
			&usuario.DocumentoUsuario,
			&usuario.Nombre,
			&usuario.Salario,
			&usuario.Rol,
			&usuario.Correo,
			&usuario.Contraseña,
			&usuario.Telefono,
			&usuario.TelefonoEmergencia,
			&usuario.Tema,
			&usuario.Colorenfacis,
			&estadoByte,
			&usuario.FechaRegistro,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}

		// Convertir el valor del BIT a bool
		usuario.Estado = estadoByte[0] == 1
		usuarios = append(usuarios, usuario)
	}

	// Verificar errores de iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
	}

	return usuarios, nil
}

func InsertarProveedor(db *sql.DB, documento string, razonSocial string, correo string, telefono string, estado bool) error {
	_, err := db.Exec("CALL InsertarProveedor()")
	if err != nil {
		return fmt.Errorf("error insertando proveedor: %v", err)
	}
	return nil
}

func ActualizarProveedor(db *sql.DB, idProveedor int, documento string, razonSocial string, correo string, telefono string, estado bool) error {
	_, err := db.Exec("CALL ActualizarProveedor(?, ?, ?, ?, ?, ?)", idProveedor, documento, razonSocial, correo, telefono, estado)
	if err != nil {
		return fmt.Errorf("error actualizando proveedor: %v", err)
	}
	return nil
}

func EliminarProveedor(db *sql.DB, idProveedor int) error {
	_, err := db.Exec("CALL EliminarProveedor(?)", idProveedor)
	if err != nil {
		return fmt.Errorf("error eliminando proveedor: %v", err)
	}
	return nil
}

func InsertarTipoServicio(db *sql.DB, descripcion string, estado bool) error {
	_, err := db.Exec("CALL InsertarTipoServicio(?, ?)", descripcion, estado)
	if err != nil {
		return fmt.Errorf("error insertando tipo de servicio: %v", err)
	}
	return nil
}

func ActualizarTipoServicio(db *sql.DB, idTipoServicio int, descripcion string, estado bool) error {
	_, err := db.Exec("CALL ActualizarTipoServicio(?, ?, ?)", idTipoServicio, descripcion, estado)
	if err != nil {
		return fmt.Errorf("error actualizando tipo de servicio: %v", err)
	}
	return nil
}

func EliminarTipoServicio(db *sql.DB, idTipoServicio int) error {
	_, err := db.Exec("CALL EliminarTipoServicio(?)", idTipoServicio)
	if err != nil {
		return fmt.Errorf("error eliminando tipo de servicio: %v", err)
	}
	return nil
}

func InsertarCategoria(db *sql.DB, descripcion string, estado bool) error {
	_, err := db.Exec("CALL InsertarCategoria(?, ?)", descripcion, estado)
	if err != nil {
		return fmt.Errorf("error insertando categoria: %v", err)
	}
	return nil
}

func ActualizarCategoria(db *sql.DB, idCategoria int, descripcion string, estado bool) error {
	_, err := db.Exec("CALL ActualizarCategoria(?, ?, ?)", idCategoria, descripcion, estado)
	if err != nil {
		return fmt.Errorf("error actualizando categoria: %v", err)
	}
	return nil
}

func EliminarCategoria(db *sql.DB, idCategoria int) error {
	_, err := db.Exec("CALL EliminarCategoria(?)", idCategoria)
	if err != nil {
		return fmt.Errorf("error eliminando categoria: %v", err)
	}
	return nil
}

func InsertarProducto(db *sql.DB, idCategoria int, nombreProducto string, descripcion string, stock int, stockMin int, valorProductoCompra float32, valorProductoVenta float32, estado bool) error {
	_, err := db.Exec("CALL InsertarProducto(?, ?, ?, ?, ?, ?, ?, ?, ?)", idCategoria, nombreProducto, descripcion, stock, stockMin, valorProductoCompra, valorProductoVenta, estado)
	if err != nil {
		return fmt.Errorf("error insertando producto: %v", err)
	}
	return nil
}

func ActualizarProducto(db *sql.DB, idProducto int, idCategoria int, nombreProducto string, descripcion string, stock int, stockMin int, valorProductoCompra float32, valorProductoVenta float32, estado bool) error {
	_, err := db.Exec("CALL ActualizarProducto(?, ?, ?, ?, ?, ?, ?, ?, ?)", idProducto, idCategoria, nombreProducto, descripcion, stock, stockMin, valorProductoCompra, valorProductoVenta, estado)
	if err != nil {
		return fmt.Errorf("error actualizando producto: %v", err)
	}
	return nil
}

func EliminarProducto(db *sql.DB, idProducto int) error {
	_, err := db.Exec("CALL EliminarProducto(?)", idProducto)
	if err != nil {
		return fmt.Errorf("error eliminando producto: %v", err)
	}
	return nil
}

type ProductoConCategoria struct {
	IdProducto          int
	Categoria           string
	NombreProducto      string
	Descripcion         string
	Stock               int
	StockMin            int
	ValorProductoCompra float32
	ValorProductoVenta  float32
	Estado              bool
}

func ObtenerProductosConCategorias(db *sql.DB) ([]ProductoConCategoria, error) {
	rows, err := db.Query("CALL ObtenerProductosConCategorias()")
	if err != nil {
		return nil, fmt.Errorf("error ejecutando ObtenerProductosConCategorias: %v", err)
	}
	defer rows.Close()

	//Slice para almacenar los resultados
	var productos []ProductoConCategoria

	for rows.Next() {
		var producto ProductoConCategoria
		var estadoByte []byte

		err := rows.Scan(
			&producto.IdProducto,
			&producto.Categoria,
			&producto.NombreProducto,
			&producto.Descripcion,
			&producto.Stock,
			&producto.StockMin,
			&producto.ValorProductoCompra,
			&producto.ValorProductoVenta,
			&estadoByte,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}

		// Convertir el valor del BIT a bool
		producto.Estado = estadoByte[0] == 1
		productos = append(productos, producto)
	}

	// Verificar errores de iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
	}

	return productos, nil
}

// InsertarMiembro calls the InsertarMiembro stored procedure to insert a new member into the database.
func InsertarMiembro(db *sql.DB, documentoMiembro, nombreMiembro string, numeroAsistencias int, correo, telefono, telefonoEmergencia string, idServicio int, fechaCaducidad, ultimaRenovacion time.Time, estado bool, imagen string) error {
	// Prepare the SQL statement to call the stored procedure
	// Execute the SQL statement
	_, err := db.Exec(`CALL InsertarMiembro(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, documentoMiembro, nombreMiembro, numeroAsistencias, correo, telefono, telefonoEmergencia, idServicio, fechaCaducidad, ultimaRenovacion, estado, imagen)
	if err != nil {
		return fmt.Errorf("error insertando miembro: %v", err)
	}

	return nil
}

// ActualizarMiembro calls the ActualizarMiembro stored procedure to update an existing member in the database.
func ActualizarMiembro(db *sql.DB, idMiembro int, documentoMiembro, nombreMiembro string, numeroAsistencias int, correo, telefono, telefonoEmergencia string, idServicio int, fechaCaducidad, ultimaRenovacion time.Time, estado bool, imagen string) error {
	// Prepare the SQL statement to call the stored procedure
	// Execute the SQL statement
	_, err := db.Exec(`CALL ActualizarMiembro(? , ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, idMiembro, documentoMiembro, nombreMiembro, numeroAsistencias, correo, telefono, telefonoEmergencia, idServicio, fechaCaducidad, ultimaRenovacion, estado, imagen)
	if err != nil {
		return fmt.Errorf("error actualizando miembro: %v", err)
	}
	return nil
}

func EliminarMiembro(db *sql.DB, idMiembro int) error {
	_, err := db.Exec("CALL EliminarMiembro(?)", idMiembro)
	if err != nil {
		return fmt.Errorf("error eliminando miembro: %v", err)
	}
	return nil
}

func ObtenerMiembroPorDocumentoMovil(db *sql.DB, documento string) error {

	err := db.QueryRow("CALL ObtenerMiembroPorDocumentoMovil(?)", documento)
	if err != nil {
		return fmt.Errorf("error obteniendo miembro por documento: %v", err)
	}
	return nil
}
