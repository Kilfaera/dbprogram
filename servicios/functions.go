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

type Miembro struct {
	IdMiembro          int
	DocumentoMiembro   string
	NombreMiembro      string
	NumeroAsistencias  int
	Correo             string
	Telefono           string
	TelefonoEmergencia string
	IdServicio         int
	FechaCaducidad     time.Time
	UltimaRenovacion   time.Time
	Estado             bool
	Imagen             string
}

func ObtenerMiembroPorDocumentoMovil(db *sql.DB, documentoMiembro string) (Miembro, error) {
	var miembro Miembro
	var estadoByte []byte
	err := db.QueryRow("CALL ObtenerMiembroPorDocumentoMovil(?)", documentoMiembro).Scan(
		&miembro.IdMiembro,
		&miembro.DocumentoMiembro,
		&miembro.NombreMiembro,
		&miembro.NumeroAsistencias,
		&miembro.Correo,
		&miembro.Telefono,
		&miembro.TelefonoEmergencia,
		&miembro.IdServicio,
		&miembro.FechaCaducidad,
		&miembro.UltimaRenovacion,
		&estadoByte,
		&miembro.Imagen,
	)
	if err != nil {
		return Miembro{}, fmt.Errorf("error obteniendo miembro por documento: %v", err)
	}
	// Convertir el valor del BIT a bool
	miembro.Estado = estadoByte[0] == 1
	return miembro, nil
}

func ObtenerMiembros(db *sql.DB) ([]Miembro, error) {
	rows, err := db.Query("CALL ObtenerMiembros()")
	if err != nil {
		return nil, fmt.Errorf("error ejecutando ObtenerMiembros: %v", err)
	}
	defer rows.Close()
	//Slice para almacenar los resultados
	var miembros []Miembro
	for rows.Next() {
		var miembro Miembro
		var estadoByte []byte
		err := rows.Scan(
			&miembro.IdMiembro,
			&miembro.DocumentoMiembro,
			&miembro.NombreMiembro,
			&miembro.NumeroAsistencias,
			&miembro.Correo,
			&miembro.Telefono,
			&miembro.TelefonoEmergencia,
			&miembro.IdServicio,
			&miembro.FechaCaducidad,
			&miembro.UltimaRenovacion,
			&estadoByte,
			&miembro.Imagen,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		// Convertir el valor del BIT a bool
		miembro.Estado = estadoByte[0] == 1
		miembros = append(miembros, miembro)
	}
	// Verificar errores de iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
	}
	return miembros, nil
}

func InsertarGasto(db *sql.DB, idProveedor int, idUsuario int, documento string, descripcion string, montoTotal float32, formaPago string) error {
	_, err := db.Exec("CALL InsertarGasto(?, ?, ?, ?, ?, ?, ?, ?, ?)", idProveedor, idUsuario, documento, descripcion, montoTotal, formaPago
	if err != nil {
		return fmt.Errorf("error insertando gasto: %v", err)
	}
	return nil
}

func ActualizarGasto(db *sql.DB, idGasto int, idProveedor int, idUsuario int, docudocumento string, descripcion string, montoTotal float32, formaPago string) error {
  _, err := db.Exec("CALL ActualizarGasto(?, ?, ?, ?, ?, ?, ?, ?, ?)", idGasto, idProveedor, idUsuario, documento, descripcion, montoTotal, formaPago
  if err != nil {
    return fmt.Errorf("error actualizando gasto: %v", err)
  }
  return nil
}

func EliminarGasto(db *sql.DB, idGasto int) error {
  _, err := db.Exec("CALL EliminarGasto(?)", idGasto)
  if err != nil {
    return fmt.Errorf("error eliminando gasto: %v", err)
  }
  return nil
}

type Gasto struct {
  IdGasto     int
  proveedor   string
  usuario     string
  documento   string
  descripcion string
  montoTotal  float32
  formaPago   string
}

func ObtenerGastosConProveedorYUsuario(db *sql.DB) ([]Gasto, error) {
  rows, err := db.Query("CALL ObtenerGastosConProveedorYUsuario()")
  if err != nil {
    return nil, fmt.Errorf("error ejecutando ObtenerGastosConProveedorYUsuario: %v", err)
  }
  defer rows.Close()
  //Slice para almacenar los resultados
  var gastos []Gasto
  for rows.Next() {
    var gasto Gasto
    err := rows.Scan(
      &gasto.IdGasto,
      &gasto.proveedor,
      &gasto.usuario,
      &gasto.documento,
      &gasto.descripcion,
      &gasto.montoTotal,
      &gasto.formaPago,
    )
    if err != nil {
      return nil, fmt.Errorf("error escaneando fila: %v", err)
    }
    gastos = append(gastos, gasto)
  }
  // Verificar errores de iteración
  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
  }
  return gastos, nil
}

func InsertarVenta(db *sql.DB, idUsuario int, idMiembro int, productos string, cantidad int, descripcion string, formaPago string, montoPago float32, montoCambio float32, montoTotal float32) error {
  _, err := db.Exec("CALL InsertarVenta(?, ?, ?, ?, ?, ?, ?)", idUsuario, idMiembro, productos, cantidad, descripcion, formaPago, montoPago, montoCambio, montoTotal)
  if err != nil {
    return fmt.Errorf("error insertando venta: %v", err)
  }
  return nil
}

func ActualizarVenta(db *sql.DB, idVenta int, idUsuario int, idMiembro int, productos string, cantidad int, descripcion string, formaPago string, montoPago float32, montoCambio float32, montoTotal float32) error {
  _, err := db.Exec("CALL ActualizarVenta(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", idVenta, idUsuario, idMiembro, productos, cantidad, descripcion, formaPago, montoPago, montoCambio, montoTotal)
  if err != nil {
    return fmt.Errorf("error actualizando venta: %v", err)
  }
  return nil
  }

func EliminarVenta(db *sql.DB, idVenta int) error {
  _, err := db.Exec("CALL EliminarVenta(?)", idVenta)
  if err != nil {
    return fmt.Errorf("error eliminando venta: %v", err)
  }
  return nil
  }

type Venta struct {
  IdVenta     int
  usuario     string
  miembro     string
  productos   string
  cantidad    int
  descripcion string
  formaPago   string
  montoPago   float32
  montoCambio float32
  montoTotal  float32
}

func ObtenerVentasConUsuarioYMiembro(db *sql.DB) ([]Venta, error) {
  rows, err := db.Query("CALL ObtenerVentasConUsuarioYMiembro()")
  if err != nil {
    return nil, fmt.Errorf("error ejecutando ObtenerVentasConUsuarioYMiembro: %v", err)
  }
  defer rows.Close()
  //Slice para almacenar los resultados
  var ventas []Venta
  for rows.Next() {
    var venta Venta
    err := rows.Scan(
      &venta.IdVenta,
      &venta.usuario,
      &venta.miembro,
      &venta.productos,
      &venta.cantidad,
      &venta.descripcion,
      &venta.formaPago,
      &venta.montoPago,
      &venta.montoCambio,
      &venta.montoTotal,
    )
    if err != nil {
      return nil, fmt.Errorf("error escaneando fila: %v", err)
    }
    ventas = append(ventas, venta)
  }
  // Verificar errores de iteración
  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
  }
  return ventas, nil
  }

  func ObtenerComprasMiembro(db *sql.DB, idMiembro int) ([]Venta, error) {
    rows, err := db.Query("CALL ObtenerComprasMiembro(?)", idMiembro)
    if err != nil {
      return nil, fmt.Errorf("error ejecutando ObtenerComprasMiembro: %v", err)
    }
    defer rows.Close()
    //Slice para almacenar los resultados
    var ventas []Venta
    for rows.Next() {
      var venta Venta
      err := rows.Scan(
        &venta.IdVenta,
        &venta.usuario,
        &venta.miembro,
        &venta.productos,
        &venta.cantidad,
        &venta.descripcion,
        &venta.formaPago,
        &venta.montoPago,
        &venta.montoCambio,
        &venta.montoTotal,
      )
      if err != nil {
        return nil, fmt.Errorf("error escaneando fila: %v", err)
      }
      ventas = append(ventas, venta)
    }
    // Verificar errores de iteración
    if err = rows.Err(); err != nil {
      return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
    }
    return ventas, nil
  }

func InsertarAsistenciaMiembro(db *sql.DB, idMiembro int, idUsuario int, nota string, idServicio int) error {
  _, err := db.Exec("CALL InsertarAsistenciaMiembro(?, ?, ?, ?, ?)", idMiembro, idUsuario, nota, idServicio)
  if err != nil {
    return fmt.Errorf("error insertando asistencia miembro: %v", err)
  }
  return nil
}

func ActualizarAsistenciaMiembro(db *sql.DB, idAsistenciaMiembro int, idMiembro int, idUsuario int, nota string, idServicio int) error {
  _, err := db.Exec("CALL ActualizarAsistenciaMiembro(?, ?, ?, ?, ?, ?)", idAsistenciaMiembro, idMiembro, idUsuario, nota, idServicio)
  if err != nil {
    return fmt.Errorf("error actualizando asistencia miembro: %v", err)
  }
  return nil
}

func EliminarAsistenciaMiembro(db *sql.DB, idAsistenciaMiembro int) error {
  _, err := db.Exec("CALL EliminarAsistenciaMiembro(?)", idAsistenciaMiembro)
  if err != nil {
    return fmt.Errorf("error eliminando asistencia miembro: %v", err)
  }
  return nil
  }

type AsistenciaMiembro struct {
  IdAsistenciaMiembro int
  miembro             string
  usuario             string
  servicio            string
  nota                string
  fechaRegistro       time.Time
  }

func ObtenerAsistenciasMiembroConUsuarioYServicio(db *sql.DB) ([]AsistenciaMiembro, error) {
  rows, err := db.Query("CALL ObtenerAsistenciasMiembroConUsuarioYServicio()")
  if err != nil {
    return nil, fmt.Errorf("error ejecutando ObtenerAsistenciasMiembroConUsuarioYServicio: %v", err)
  }
  defer rows.Close()
  //Slice para almacenar los resultados
  var asistencias []AsistenciaMiembro
  for rows.Next() {
    var asistencia AsistenciaMiembro
    err := rows.Scan(
      &asistencia.IdAsistenciaMiembro,
      &asistencia.miembro,
      &asistencia.usuario,
      &asistencia.servicio,
      &asistencia.nota,
      &asistencia.fechaRegistro,
    )
    if err != nil {
      return nil, fmt.Errorf("error escaneando fila: %v", err)
    }
    asistencias = append(asistencias, asistencia)
  }
  // Verificar errores de iteración
  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
  }
  return asistencias, nil
  }

  func ObtenerAsistenciaMiembro(db *sql.DB, idMiembro int) ([]AsistenciaMiembro, error) {
    rows, err := db.Query("CALL ObtenerAsistenciasMiembroPorMiembro(?)", idMiembro)
    if err != nil {
      return nil, fmt.Errorf("error ejecutando ObtenerAsistenciasMiembroPorMiembro: %v", err)
    }
    defer rows.Close()
    //Slice para almacenar los resultados
    var asistencias []AsistenciaMiembro
    for rows.Next() {
      var asistencia AsistenciaMiembro
      err := rows.Scan(
      &asistencia.IdAsistenciaMiembro,
      &asistencia.miembro,
      &asistencia.usuario,
      &asistencia.servicio,
      &asistencia.nota,
      &asistencia.fechaRegistro,
      )
      if err != nil {
      return nil, fmt.Errorf("error escaneando fila: %v", err)
      }
      asistencias = append(asistencias, asistencia)
    }
    // Verificar errores de iteración
    if err = rows.Err(); err != nil {
      return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
    }
    return asistencias, nil
}

fun InsertarNovedad(db *sql.DB, titulo string, texto, string, fechaExpiracion time.Time, idUsuario int) error {
  _, err := db.Exec("CALL InsertarNovedad(?, ?, ?, ?)", titulo, texto, fechaExpiracion, idUsuario)
  if err != nil {
    return fmt.Errorf("error insertando novedad: %v", err)
      }
    return nil
  }

func ActualizarNovedad(db *sql.DB, idNovedad int, titulo string, texto string, fechaExpiracion time.Time, idUsuario int) error {
  _, err := db.Exec("CALL ActualizarNovedad(?, ?, ?, ?, ?)", idNovedad, titulo, texto, fechaExpiracion, idUsuario)
  if err != nil {
    return fmt.Errorf("error actualizando novedad: %v", err)
  }
  return nil
}

func EliminarNovedad(db *sql.DB, idNovedad int) error {
  _, err := db.Exec("CALL EliminarNovedad(?)", idNovedad)
  if err != nil {
    return fmt.Errorf("error eliminando novedad: %v", err)
  }
  return nil
  }

type Novedad struct {
  IdNovedad       int
  titulo          string
  texto           string
  fechaRegistro   time.Time
  fechaExpiracion time.Time
  usuario         string
}

func ObtenerNovedadesConUsuario(db *sql.DB) ([]Novedad, error) {
  rows, err := db.Query("CALL ObtenerNovedadesConUsuario()")
  if err != nil {
    return nil, fmt.Errorf("error ejecutando ObtenerNovedadesConUsuario: %v", err)
  }
  defer rows.Close()
  //Slice para almacenar los resultados
  var novedades []Novedad
  for rows.Next() {
    var novedad Novedad
    err := rows.Scan(
      &novedad.IdNovedad,
      &novedad.titulo,
      &novedad.texto,
      &novedad.fechaRegistro,
      &novedad.fechaExpiracion,
      &novedad.usuario,
    )
    if err != nil {
      return nil, fmt.Errorf("error escaneando fila: %v", err)
    }
    novedades = append(novedades, novedad)
  }
  // Verificar errores de iteración
  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
  }
  return novedades, nil
  }
  
func InsertarNotificacion(db *sql.DB, titutlo string, texto string, destinatario int, estado bool){
  _, err := db.Exec("CALL InsertarNotificacion(?, ?, ?, ?)", titulo, texto, destinatario, estado)
  if err != nil {
    return fmt.Errorf("error insertando notificacion: %v", err)
  }
  return nil
}

func ActualizarNotificacion(db *sql.DB, idNotificacion int, titulo string, texto string, destinatario int, estado bool){
  _, err := db.Exec("CALL ActualizarNotificacion(?, ?, ?, ?, ?)", idNotificacion, titulo, texto, destinatario, estado)
  if err != nil {
    return fmt.Errorf("error actualizando notificacion: %v", err)
  }
  return nil
}

func EliminarNotificacion(db *sql.DB, idNotificacion int){
  _, err := db.Exec("CALL EliminarNotificacion(?)", idNotificacion)
  if err != nil {
    return fmt.Errorf("error eliminando notificacion: %v", err)
  }
  return nil
  }

type Notificacion struct {
  IdNotificacion  int
  titulo          string
  texto           string
  fechaRegistro   time.Time
  destinatario    string
  estado          bool
}

func ObtenerNotificacionMiembro(db *sql.DB) ([]Notificacion, error) {
  rows, err := db.Query("CALL ObtenerNotificacionesConDestinatario()")
  if err != nil {
    return nil, fmt.Errorf("error ejecutando ObtenerNotificacionesConDestinatario: %v", err)
  }
  defer rows.Close()
  //Slice para almacenar los resultados
  var notificaciones []Notificacion
  for rows.Next() {
    var notificacion Notificacion
    var estadoByte []byte
    err := rows.Scan(
      &notificacion.IdNotificacion,
      &notificacion.titulo,
      &notificacion.texto,
      &notificacion.fechaRegistro,
      &notificacion.destinatario,
      &estadoByte,
    )
    if err != nil {
      return nil, fmt.Errorf("error escaneando fila: %v", err)
    }
    // Convertir el valor del BIT a bool
    notificacion.estado = estadoByte[0] == 1
    notificaciones = append(notificaciones, notificacion)
  }
  // Verificar errores de iteración
  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("error durante la iteración de resultados: %v", err)
  }
  return notificaciones, nil
  }
