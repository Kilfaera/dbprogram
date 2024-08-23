#==============================================================================
# Rol ///////////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Rol(
    IdRol INT AUTO_INCREMENT PRIMARY KEY,
    Descripcion VARCHAR(50),
    MenuInicio BIT,
    MenuMiembro BIT,
    MenuReporte BIT,
    MenuConfiguracion BIT,
    MenuInformacion BIT, 
    MenuUsuario BIT,
    MenuReset BIT,
    FechaRegistro DATETIME
);

#==============================================================================
# Usuario ///////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Usuario (
    IdUsuario INT AUTO_INCREMENT PRIMARY KEY,
    DocumentoUsuario VARCHAR(50),
    Nombre VARCHAR(100),
    Salario DECIMAL(10,2),
    IdRol INT,
    Correo VARCHAR(50),
    Contraseña VARCHAR(50),
    Telefono VARCHAR(50),
    TelefonoEmergencia VARCHAR(50),
    Tema VARCHAR(50),
    Colorenfacis VARCHAR(50),
    Estado BIT,
    FechaRegistro DATETIME,
    FOREIGN KEY (IdRol) REFERENCES Rol(IdRol)
);

#==============================================================================
# Proveedor /////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Proveedor(
    IdProveedor INT AUTO_INCREMENT PRIMARY KEY,
    Documento VARCHAR(50),
    RazonSocial VARCHAR(50),
    Correo VARCHAR(50),
    Telefono VARCHAR(50),
    Estado BIT,
    FechaRegistro DATETIME 
);

#==============================================================================
# TipoServicio //////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE TipoServicio(
    IdServicio INT AUTO_INCREMENT PRIMARY KEY,
    NombreServicio VARCHAR(100),
    TiempoServicio VARCHAR(50)
);

#==============================================================================
# Categoria /////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Categoria(
    IdCategoria INT AUTO_INCREMENT PRIMARY KEY, 
    NombreCategoria VARCHAR(50),
    Descripcion VARCHAR(50)
);

#==============================================================================
# Producto //////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Producto (
    IdProducto INT AUTO_INCREMENT PRIMARY KEY,
    IdCategoria INT,
    NombreProducto VARCHAR(50),
    Descripcion VARCHAR(50),
    Stock INT,
    StockMin INT,
    ValorProductoCompra DECIMAL(10,2),
    ValorProductoVenta DECIMAL(10,2),
    Estado BIT,
    FOREIGN KEY (IdCategoria) REFERENCES Categoria(IdCategoria)
);

#==============================================================================
# Miembro ///////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Miembro (
    IdMiembro INT AUTO_INCREMENT PRIMARY KEY,
    DocumentoMiembro VARCHAR(50),
    NombreMiembro VARCHAR(100),
    FechaRegistro DATETIME,
    NumeroAsistencias INT,
    Correo VARCHAR(50),
    Telefono VARCHAR(50),
    TelefonoEmergencia VARCHAR(50),
    IdServicio INT,
    FechaCaducidad DATETIME,
    UltimaRenovacion DATETIME,
    Estado BIT,
    Imagen VARCHAR(255),
    FOREIGN KEY (IdServicio) REFERENCES TipoServicio(IdServicio)
);

#==============================================================================
# Gasto /////////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Gasto (
    IdGasto INT AUTO_INCREMENT PRIMARY KEY,
    IdProveedor INT,
    IdUsuario INT,
    Documento VARCHAR(100),
    Descripcion VARCHAR(100),
    MontoTotal DECIMAL(10,2),
    FormaPago VARCHAR(50),
    FechaRegistro DATETIME,
    FOREIGN KEY (IdProveedor) REFERENCES Proveedor(IdProveedor),
    FOREIGN KEY (IdUsuario) REFERENCES Usuario(IdUsuario)
);

#==============================================================================
# Venta /////////////////////////////////////////////////////////////////////////
#==============================================================================

CREATE TABLE Venta (
    IdVenta INT AUTO_INCREMENT PRIMARY KEY,
    IdUsuario INT,
    IdMiembro INT,
    Productos TEXT,
    Cantidad INT,
    Descripcion VARCHAR(100),
    FormaPago VARCHAR(50),
    MontoPago DECIMAL(10,2),
    MontoCambio DECIMAL(10,2),
    MontoTotal DECIMAL(10,2),
    FechaRegistro DATETIME,
    FOREIGN KEY (IdUsuario) REFERENCES Usuario(IdUsuario),
    FOREIGN KEY (IdMiembro) REFERENCES Miembro(IdMiembro)
);


#==============================================================================
#Rol procedures///////////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarRol(
    IN Descripcion VARCHAR(50),
    IN MenuInicio BIT,
    IN MenuMiembro BIT,
    IN MenuReporte BIT,
    IN MenuConfiguracion BIT,
    IN MenuInformacion BIT,
    IN MenuUsuario BIT,
    IN MenuReset BIT
)
BEGIN
    INSERT INTO Rol(Descripcion, MenuInicio, MenuMiembro, MenuReporte, MenuConfiguracion, MenuInformacion, MenuUsuario, MenuReset, FechaRegistro)
    VALUES (Descripcion, MenuInicio, MenuMiembro, MenuReporte, MenuConfiguracion, MenuInformacion, MenuUsuario, MenuReset, NOW());
END //

CREATE PROCEDURE ActualizarRol(
    IN IdRolToUpdate INT,
    IN NewDescripcion VARCHAR(50),
    IN NewMenuInicio BIT,
    IN NewMenuMiembro BIT,
    IN NewMenuReporte BIT,
    IN NewMenuConfiguracion BIT,
    IN NewMenuInformacion BIT,
    IN NewMenuUsuario BIT,
    IN NewMenuReset BIT
)
BEGIN
    UPDATE Rol
    SET Descripcion = NewDescripcion,
        MenuInicio = NewMenuInicio,
        MenuMiembro = NewMenuMiembro,
        MenuReporte = NewMenuReporte,
        MenuConfiguracion = NewMenuConfiguracion,
        MenuInformacion = NewMenuInformacion,
        MenuUsuario = NewMenuUsuario,
        MenuReset = NewMenuReset
    WHERE IdRol = IdRolToUpdate;
END //

CREATE PROCEDURE EliminarRol(
    IN IdRolToDelete INT
)
BEGIN
    DELETE FROM Rol
    WHERE IdRol = IdRolToDelete;
END //

DELIMITER ;


#==============================================================================
#Usuario procedures///////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarUsuario(
    IN DocumentoUsuario VARCHAR(50),
    IN Nombre VARCHAR(100),
    IN Salario DECIMAL(10,2),
    IN IdRol INT,
    IN Correo VARCHAR(50),
    IN Contraseña VARCHAR(50),
    IN Telefono VARCHAR(50),
    IN TelefonoEmergencia VARCHAR(50),
    IN Tema VARCHAR(50),
    IN Colorenfacis VARCHAR(50),
    IN Estado BIT
)
BEGIN
    INSERT INTO Usuario(DocumentoUsuario, Nombre, Salario, IdRol, Correo, Contraseña, Telefono, TelefonoEmergencia, Tema, Colorenfacis, Estado, FechaRegistro)
    VALUES (DocumentoUsuario, Nombre, Salario, IdRol, Correo, Contraseña, Telefono, TelefonoEmergencia, Tema, Colorenfacis, Estado, NOW());
END //

CREATE PROCEDURE ActualizarUsuario(
    IN IdUsuarioToUpdate INT,
    IN NewDocumentoUsuario VARCHAR(50),
    IN NewNombre VARCHAR(100),
    IN NewSalario DECIMAL(10,2),
    IN NewIdRol INT,
    IN NewCorreo VARCHAR(50),
    IN NewContraseña VARCHAR(50),
    IN NewTelefono VARCHAR(50),
    IN NewTelefonoEmergencia VARCHAR(50),
    IN NewTema VARCHAR(50),
    IN NewColorenfacis VARCHAR(50),
    IN NewEstado BIT
)
BEGIN
    UPDATE Usuario
    SET DocumentoUsuario = NewDocumentoUsuario,
        Nombre = NewNombre,
        Salario = NewSalario,
        IdRol = NewIdRol,
        Correo = NewCorreo,
        Contraseña = NewContraseña,
        Telefono = NewTelefono,
        TelefonoEmergencia = NewTelefonoEmergencia,
        Tema = NewTema,
        Colorenfacis = NewColorenfacis,
        Estado = NewEstado
    WHERE IdUsuario = IdUsuarioToUpdate;
END //

CREATE PROCEDURE EliminarUsuario(
    IN IdUsuarioToDelete INT
)
BEGIN
    DELETE FROM Usuario
    WHERE IdUsuario = IdUsuarioToDelete;
END //

CREATE PROCEDURE ObtenerUsuariosConRoles()
BEGIN
    SELECT U.IdUsuario, U.DocumentoUsuario, U.Nombre, U.Salario, R.Descripcion as 'Rol', U.Correo, U.Contraseña, U.Telefono, U.TelefonoEmergencia, U.Tema, U.Colorenfacis, U.Estado, U.FechaRegistro
    FROM Usuario U
    INNER JOIN Rol R ON U.IdRol = R.IdRol;
END //

DELIMITER ;


#==============================================================================
#Proveedor procedures/////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarProveedor(
    IN Documento VARCHAR(50),
    IN RazonSocial VARCHAR(50),
    IN Correo VARCHAR(50),
    IN Telefono VARCHAR(50),
    IN Estado BIT
)
BEGIN
    INSERT INTO Proveedor(Documento, RazonSocial, Correo, Telefono, Estado, FechaRegistro)
    VALUES (Documento, RazonSocial, Correo, Telefono, Estado, NOW());
END //

CREATE PROCEDURE ActualizarProveedor(
    IN IdProveedorToUpdate INT,
    IN NewDocumento VARCHAR(50),
    IN NewRazonSocial VARCHAR(50),
    IN NewCorreo VARCHAR(50),
    IN NewTelefono VARCHAR(50),
    IN NewEstado BIT
)
BEGIN
    UPDATE Proveedor
    SET Documento = NewDocumento,
        RazonSocial = NewRazonSocial,
        Correo = NewCorreo,
        Telefono = NewTelefono,
        Estado = NewEstado
    WHERE IdProveedor = IdProveedorToUpdate;
END //

CREATE PROCEDURE EliminarProveedor(
    IN IdProveedorToDelete INT
)
BEGIN
    DELETE FROM Proveedor
    WHERE IdProveedor = IdProveedorToDelete;
END //

DELIMITER ;

#==============================================================================
#TipoServicio procedures//////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarTipoServicio(
    IN NombreServicio VARCHAR(100),
    IN TiempoServicio VARCHAR(50)
)
BEGIN
    INSERT INTO TipoServicio(NombreServicio, TiempoServicio)
    VALUES (NombreServicio, TiempoServicio);
END //

CREATE PROCEDURE ActualizarTipoServicio(
    IN IdServicioToUpdate INT,
    IN NewNombreServicio VARCHAR(100),
    IN NewTiempoServicio VARCHAR(50)
)
BEGIN
    UPDATE TipoServicio
    SET NombreServicio = NewNombreServicio,
        TiempoServicio = NewTiempoServicio
    WHERE IdServicio = IdServicioToUpdate;
END //

CREATE PROCEDURE EliminarTipoServicio(
    IN IdServicioToDelete INT
)
BEGIN
    DELETE FROM TipoServicio
    WHERE IdServicio = IdServicioToDelete;
END //

DELIMITER ;


#==============================================================================
#Categoria procedures/////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarCategoria(
    IN NombreCategoria VARCHAR(50),
    IN Descripcion VARCHAR(50)
)
BEGIN
    INSERT INTO Categoria(NombreCategoria, Descripcion)
    VALUES (NombreCategoria, Descripcion);
END //

CREATE PROCEDURE ActualizarCategoria(
    IN IdCategoriaToUpdate INT,
    IN NewNombreCategoria VARCHAR(50),
    IN NewDescripcion VARCHAR(50)
)
BEGIN
    UPDATE Categoria
    SET NombreCategoria = NewNombreCategoria,
        Descripcion = NewDescripcion
    WHERE IdCategoria = IdCategoriaToUpdate;
END //

CREATE PROCEDURE EliminarCategoria(
    IN IdCategoriaToDelete INT
)
BEGIN
    DELETE FROM Categoria
    WHERE IdCategoria = IdCategoriaToDelete;
END //

DELIMITER ;

#==============================================================================
#Producto procedures//////////////////////////////////////////////////////////////////////
#==============================================================================

DELIMITER //

CREATE PROCEDURE InsertarProducto(
    IN IdCategoria INT,
    IN NombreProducto VARCHAR(50),
    IN Descripcion VARCHAR(50),
    IN Stock INT,
    IN StockMin INT,
    IN ValorProductoCompra DECIMAL(10,2),
    IN ValorProductoVenta DECIMAL(10,2),
    IN Estado BIT
)
BEGIN
    INSERT INTO Producto(IdCategoria, NombreProducto, Descripcion, Stock, StockMin, ValorProductoCompra, ValorProductoVenta, Estado)
    VALUES (IdCategoria, NombreProducto, Descripcion, Stock, StockMin, ValorProductoCompra, ValorProductoVenta, Estado);
END //

CREATE PROCEDURE ActualizarProducto(
    IN IdProductoToUpdate INT,
    IN NewIdCategoria INT,
    IN NewNombreProducto VARCHAR(50),
    IN NewDescripcion VARCHAR(50),
    IN NewStock INT,
    IN NewStockMin INT,
    IN NewValorProductoCompra DECIMAL(10,2),
    IN NewValorProductoVenta DECIMAL(10,2),
    IN NewEstado BIT
)
BEGIN
    UPDATE Producto
    SET IdCategoria = NewIdCategoria,
        NombreProducto = NewNombreProducto,
        Descripcion = NewDescripcion,
        Stock = NewStock,
        StockMin = NewStockMin,
        ValorProductoCompra = NewValorProductoCompra,
        ValorProductoVenta = NewValorProductoVenta,
        Estado = NewEstado
    WHERE IdProducto = IdProductoToUpdate;
END //

CREATE PROCEDURE EliminarProducto(
    IN IdProductoToDelete INT
)
BEGIN
    DELETE FROM Producto
    WHERE IdProducto = IdProductoToDelete;
END //

CREATE PROCEDURE ObtenerProductosConCategoria()
BEGIN
    SELECT P.IdProducto, C.NombreCategoria, P.NombreProducto, P.Descripcion, P.Stock, P.StockMin, P.ValorProductoCompra, P.ValorProductoVenta, P.Estado
    FROM Producto P
    INNER JOIN Categoria C ON P.IdCategoria = C.IdCategoria;
END //

DELIMITER ;


#==============================================================================
#Miembro procedures///////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarMiembro(
    IN DocumentoMiembro VARCHAR(50),
    IN NombreMiembro VARCHAR(100),
    IN NumeroAsistencias INT,
    IN Correo VARCHAR(50),
    IN Telefono VARCHAR(50),
    IN TelefonoEmergencia VARCHAR(50),
    IN IdServicio INT,
    IN FechaCaducidad DATETIME,
    IN UltimaRenovacion DATETIME,
    IN Estado BIT,
    IN Imagen VARCHAR(255)
)
BEGIN
    INSERT INTO Miembro(DocumentoMiembro, NombreMiembro, NumeroAsistencias, Correo, Telefono, TelefonoEmergencia, IdServicio, FechaCaducidad, UltimaRenovacion, Estado, Imagen, FechaRegistro)
    VALUES (DocumentoMiembro, NombreMiembro, NumeroAsistencias, Correo, Telefono, TelefonoEmergencia, IdServicio, FechaCaducidad, UltimaRenovacion, Estado, Imagen, NOW());
END //

CREATE PROCEDURE ActualizarMiembro(
    IN IdMiembroToUpdate INT,
    IN NewDocumentoMiembro VARCHAR(50),
    IN NewNombreMiembro VARCHAR(100),
    IN NewNumeroAsistencias INT,
    IN NewCorreo VARCHAR(50),
    IN NewTelefono VARCHAR(50),
    IN NewTelefonoEmergencia VARCHAR(50),
    IN NewIdServicio INT,
    IN NewFechaCaducidad DATETIME,
    IN NewUltimaRenovacion DATETIME,
    IN NewEstado BIT,
    IN NewImagen VARCHAR(255)
)
BEGIN
    UPDATE Miembro
    SET DocumentoMiembro = NewDocumentoMiembro,
        NombreMiembro = NewNombreMiembro,
        NumeroAsistencias = NewNumeroAsistencias,
        Correo = NewCorreo,
        Telefono = NewTelefono,
        TelefonoEmergencia = NewTelefonoEmergencia,
        IdServicio = NewIdServicio,
        FechaCaducidad = NewFechaCaducidad,
        UltimaRenovacion = NewUltimaRenovacion,
        Estado = NewEstado,
        Imagen = NewImagen
    WHERE IdMiembro = IdMiembroToUpdate;
END //

CREATE PROCEDURE EliminarMiembro(
    IN IdMiembroToDelete INT
)
BEGIN
    DELETE FROM Miembro
    WHERE IdMiembro = IdMiembroToDelete;
END //

CREATE PROCEDURE ObtenerMiembroPorDocumentoMovil(
    IN DocumentoMiembro VARCHAR(50)
)
BEGIN
    SELECT M.IdMiembro, M.DocumentoMiembro, M.NombreMiembro, M.NumeroAsistencias, M.Correo, M.Telefono, M.TelefonoEmergencia, TS.NombreServicio AS Servicio, M.FechaCaducidad, M.UltimaRenovacion, M.Estado, M.Imagen
    FROM Miembro M
    INNER JOIN TipoServicio TS ON M.IdServicio = TS.IdServicio
    WHERE M.DocumentoMiembro = DocumentoMiembro;
END //

CREATE PROCEDURE ObtenerMiembro(
)
BEGIN
    SELECT M.IdMiembro, M.DocumentoMiembro, M.NombreMiembro, M.NumeroAsistencias, M.Correo, M.Telefono, M.TelefonoEmergencia, TS.NombreServicio AS Servicio, M.FechaCaducidad, M.UltimaRenovacion, M.Estado, M.Imagen
    FROM Miembro M
    INNER JOIN TipoServicio TS ON M.IdServicio = TS.IdServicio;
END //

DELIMITER ;

#==============================================================================
#Gasto procedures/////////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarGasto(
    IN IdProveedor INT,
    IN IdUsuario INT,
    IN Documento VARCHAR(100),
    IN Descripcion VARCHAR(100),
    IN MontoTotal DECIMAL(10,2),
    IN FormaPago VARCHAR(50)
)
BEGIN
    INSERT INTO Gasto(IdProveedor, IdUsuario, Documento, Descripcion, MontoTotal, FormaPago, FechaRegistro)
    VALUES (IdProveedor, IdUsuario, Documento, Descripcion, MontoTotal, FormaPago, NOW());
END //

CREATE PROCEDURE ActualizarGasto(
    IN IdGastoToUpdate INT,
    IN NewIdProveedor INT,
    IN NewIdUsuario INT,
    IN NewDocumento VARCHAR(100),
    IN NewDescripcion VARCHAR(100),
    IN NewMontoTotal DECIMAL(10,2),
    IN NewFormaPago VARCHAR(50)
)
BEGIN
    UPDATE Gasto
    SET IdProveedor = NewIdProveedor,
        IdUsuario = NewIdUsuario,
        Documento = NewDocumento,
        Descripcion = NewDescripcion,
        MontoTotal = NewMontoTotal,
        FormaPago = NewFormaPago
    WHERE IdGasto = IdGastoToUpdate;
END //

CREATE PROCEDURE EliminarGasto(
    IN IdGastoToDelete INT
)
BEGIN
    DELETE FROM Gasto
    WHERE IdGasto = IdGastoToDelete;
END //

CREATE PROCEDURE ObtenerGastosConProveedorYUsuario()
BEGIN
    SELECT G.IdGasto, P.RazonSocial AS Proveedor, U.Nombre AS Usuario, G.Documento, G.Descripcion, G.MontoTotal, G.FormaPago, G.FechaRegistro
    FROM Gasto G
    INNER JOIN Proveedor P ON G.IdProveedor = P.IdProveedor
    INNER JOIN Usuario U ON G.IdUsuario = U.IdUsuario;
END //

DELIMITER ;


#==============================================================================
#Venta procedures/////////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarVenta(
    IN IdUsuario INT,
    IN IdMiembro INT,
    IN Productos TEXT,
    IN Cantidad INT,
    IN Descripcion VARCHAR(100),
    IN FormaPago VARCHAR(50),
    IN MontoPago DECIMAL(10,2),
    IN MontoCambio DECIMAL(10,2),
    IN MontoTotal DECIMAL(10,2)
)
BEGIN
    INSERT INTO Venta(IdUsuario, IdMiembro, Productos, Cantidad, Descripcion, FormaPago, MontoPago, MontoCambio, MontoTotal, FechaRegistro)
    VALUES (IdUsuario, IdMiembro, Productos, Cantidad, Descripcion, FormaPago, MontoPago, MontoCambio, MontoTotal, NOW());
END //

CREATE PROCEDURE ActualizarVenta(
    IN IdVentaToUpdate INT,
    IN NewIdUsuario INT,
    IN NewIdMiembro INT,
    IN NewProductos TEXT,
    IN NewCantidad INT,
    IN NewDescripcion VARCHAR(100),
    IN NewFormaPago VARCHAR(50),
    IN NewMontoPago DECIMAL(10,2),
    IN NewMontoCambio DECIMAL(10,2),
    IN NewMontoTotal DECIMAL(10,2)
)
BEGIN
    UPDATE Venta
    SET IdUsuario = NewIdUsuario,
        IdMiembro = NewIdMiembro,
        Productos = NewProductos,
        Cantidad = NewCantidad,
        Descripcion = NewDescripcion,
        FormaPago = NewFormaPago,
        MontoPago = NewMontoPago,
        MontoCambio = NewMontoCambio,
        MontoTotal = NewMontoTotal
    WHERE IdVenta = IdVentaToUpdate;
END //

CREATE PROCEDURE EliminarVenta(
    IN IdVentaToDelete INT
)
BEGIN
    DELETE FROM Venta
    WHERE IdVenta = IdVentaToDelete;
END //

CREATE PROCEDURE ObtenerVentasConUsuarioYMiembro()
BEGIN
    SELECT V.IdVenta, U.Nombre AS Usuario, M.NombreMiembro AS Miembro, V.Productos, V.Cantidad, V.Descripcion, V.FormaPago, V.MontoPago, V.MontoCambio, V.MontoTotal, V.FechaRegistro
    FROM Venta V
    INNER JOIN Usuario U ON V.IdUsuario = U.IdUsuario
    INNER JOIN Miembro M ON V.IdMiembro = M.IdMiembro;

END //

CREATE PROCEDURE ObtenerComprasMiembro(
    in p_IdMiembro int
    )
BEGIN
    SELECT V.IdVenta, U.Nombre AS Usuario, M.NombreMiembro AS Miembro, V.Productos, V.Cantidad, V.Descripcion, V.FormaPago, V.MontoPago, V.MontoCambio, V.MontoTotal, V.FechaRegistro
    FROM Venta V
    INNER JOIN Usuario U ON V.IdUsuario = U.IdUsuario
    INNER JOIN Miembro M ON V.IdMiembro = M.IdMiembro
    WHERE IdMiembro = p_IdMiembro;
END //
DELIMITER ;


#==============================================================================
#AsistenciaMiembro procedures/////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarAsistenciaMiembro(
    IN IdMiembro INT,
    IN IdUsuario INT,
    IN Nota TEXT,
    IN IdServicio INT
)
BEGIN
    INSERT INTO AsistenciaMiembro(IdMiembro, IdUsuario, Nota, IdServicio, FechaRegistro)
    VALUES (IdMiembro, IdUsuario, Nota, IdServicio, NOW());
END //

CREATE PROCEDURE ActualizarAsistenciaMiembro(
    IN IdAsistenciaMiembroToUpdate INT,
    IN NewIdMiembro INT,
    IN NewIdUsuario INT,
    IN NewNota TEXT,
    IN NewIdServicio INT
)
BEGIN
    UPDATE AsistenciaMiembro
    SET IdMiembro = NewIdMiembro,
        IdUsuario = NewIdUsuario,
        Nota = NewNota,
        IdServicio = NewIdServicio
    WHERE IdAsistenciaMiembro = IdAsistenciaMiembroToUpdate;
END //

CREATE PROCEDURE EliminarAsistenciaMiembro(
    IN IdAsistenciaMiembroToDelete INT
)
BEGIN
    DELETE FROM AsistenciaMiembro
    WHERE IdAsistenciaMiembro = IdAsistenciaMiembroToDelete;
END //

CREATE PROCEDURE ObtenerAsistenciasConMiembroUsuarioYServicio()
BEGIN
    SELECT A.IdAsistenciaMiembro, M.NombreMiembro AS Miembro, U.Nombre AS Usuario, TS.NombreServicio AS Servicio, A.Nota, A.FechaRegistro
    FROM AsistenciaMiembro A
    INNER JOIN Miembro M ON A.IdMiembro = M.IdMiembro
    INNER JOIN Usuario U ON A.IdUsuario = U.IdUsuario
    INNER JOIN TipoServicio TS ON A.IdServicio = TS.IdServicio;
END //

DELIMITER ;
CREATE PROCEDURE ObtenerAsistenciasMiembro(
    in p_IdMiembro int
)
BEGIN
    SELECT A.IdAsistenciaMiembro, M.NombreMiembro AS Miembro, U.Nombre AS Usuario, TS.NombreServicio AS Servicio, A.Nota, A.FechaRegistro
    FROM AsistenciaMiembro A
    INNER JOIN Miembro M ON A.IdMiembro = M.IdMiembro
    INNER JOIN Usuario U ON A.IdUsuario = U.IdUsuario
    INNER JOIN TipoServicio TS ON A.IdServicio = TS.IdServicio
    where IdMiembro = p_p_IdMiembro
END //

DELIMITER ;

#==============================================================================
#Novedad procedures///////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarNovedad(
    IN Titulo TEXT,
    IN Texto TEXT,
    IN FechaExpiracion DATETIME,
    IN IdUsuario INT
)
BEGIN
    INSERT INTO Novedad(Titulo, Texto, FechaExpiracion, IdUsuario, FechaRegistro)
    VALUES (Titulo, Texto, FechaExpiracion, IdUsuario, NOW());
END //

CREATE PROCEDURE ActualizarNovedad(
    IN IdNovedadToUpdate INT,
    IN NewTitulo TEXT,
    IN NewTexto TEXT,
    IN NewFechaExpiracion DATETIME,
    IN NewIdUsuario INT
)
BEGIN
    UPDATE Novedad
    SET Titulo = NewTitulo,
        Texto = NewTexto,
        FechaExpiracion = NewFechaExpiracion,
        IdUsuario = NewIdUsuario
    WHERE IdNovedad = IdNovedadToUpdate;
END //

CREATE PROCEDURE EliminarNovedad(
    IN IdNovedadToDelete INT
)
BEGIN
    DELETE FROM Novedad
    WHERE IdNovedad = IdNovedadToDelete;
END //

CREATE PROCEDURE ObtenerNovedadesConUsuario()
BEGIN
    SELECT N.IdNovedad, N.Titulo, N.Texto, N.FechaRegistro, N.FechaExpiracion, U.Nombre AS Usuario
    FROM Novedad N
    INNER JOIN Usuario U ON N.IdUsuario = U.IdUsuario;
END //

DELIMITER ;


#==============================================================================
#Notificacion procedures///////////////////////////////////////////////////////////////////////
#==============================================================================
DELIMITER //

CREATE PROCEDURE InsertarNotificacion(
    IN Titulo TEXT,
    IN Texto TEXT,
    IN Destinatario int,
    IN Estado BIT
)
BEGIN
    INSERT INTO Notificacion(Titulo, Texto, FechaRegistro, Destinatario, Estado)
    VALUES (Titulo, Texto, NOW(), Destinatario, Estado);
END //

CREATE PROCEDURE ActualizarNotificacion(
    IN IdNotificacionToUpdate INT,
    IN NewTitulo TEXT,
    IN NewTexto TEXT,
    IN NewDestinatario int,
    IN NewEstado BIT
)
BEGIN
    UPDATE Notificacion
    SET Titulo = NewTitulo,
        Texto = NewTexto,
        Destinatario = NewDestinatario,
        Estado = NewEstado
    WHERE IdNotificacion = IdNotificacionToUpdate;
END //

CREATE PROCEDURE EliminarNotificacion(
    IN IdNotificacionToDelete INT
)
BEGIN
    DELETE FROM Notificacion
    WHERE IdNotificacion = IdNotificacionToDelete;
END //

CREATE Procedure obtenerNotificacionMiembro(
    in p_Destinatario int
)
BEGIN
    SELECT * FROM Notificacion 
    WHERE `Destinatario`= p_Destinatario;
END//

DELIMITER ;









CALL InsertarRol('Administrador', 1, 1, 1, 1, 1, 1, 1);
CALL InsertarRol('Empleado', 1, 1, 1, 0, 0, 1, 0);
CALL InsertarUsuario('12345678', 'Juan Perez', 2000.00, 1, 'juan@example.com', 'contraseña', '123456789', '987654321', 'Tema', 'Color', 1);
CALL InsertarUsuario('87654321', 'Maria Garcia', 1500.00, 2, 'maria@example.com', 'contraseña', '987654321', '123456789', 'Tema', 'Color', 1);
CALL InsertarProveedor('123456789', 'Proveedor A', 'proveedorA@example.com', '123456789', 1);
CALL InsertarProveedor('987654321', 'Proveedor B', 'proveedorB@example.com', '987654321', 1);
CALL InsertarTipoServicio('Servicio A', '10 horas');
CALL InsertarTipoServicio('Servicio B', '5 horas');
CALL InsertarCategoria('Categoria A', 'Descripción A');
CALL InsertarCategoria('Categoria B', 'Descripción B');
CALL InsertarProducto(1, 'Producto A', 'Descripción A', 100, 10, 10.00, 20.00, 1);
CALL InsertarProducto(2, 'Producto B', 'Descripción B', 50, 5, 20.00, 30.00, 1);
CALL InsertarMiembro('12345678', 'Miembro A', 5, 'miembroA@example.com', '123456789', '987654321', 1, '2024-12-31 00:00:00', '2024-04-25 00:00:00', 1, 'imagen.jpg');
CALL InsertarMiembro('1001175706', 'Miembro c', 10, 'miembroB@example.com', '987654321', '123456789', 2, '2024-12-31 00:00:00', '2024-04-25 00:00:00', 1, 'https://drive.google.com/file/d/1vBwSGhbv_PEhMpleq_6Ac16fOWgLTDB_/view');
CALL InsertarGasto(1, 1, 'Factura A', 'Descripción A', 100.00, 'Efectivo');
CALL InsertarGasto(2, 2, 'Factura B', 'Descripción B', 150.00, 'Tarjeta');
CALL InsertarVenta(1, 1, '1,2', 2, 'Venta A', 'Efectivo', 50.00, 0.00, 100.00);
CALL InsertarVenta(2, 2, '1', 1, 'Venta B', 'Tarjeta', 30.00, 0.00, 30.00);
CALL InsertarAsistenciaMiembro(1, 1, 'Nota A', 1);
CALL InsertarAsistenciaMiembro(2, 2, 'Nota B', 2);
CALL InsertarNovedad('Novedad A', 'Texto A', '2024-05-01 00:00:00', 1);
CALL InsertarNovedad('Novedad B', 'Texto B', '2024-05-01 00:00:00', 2);
CALL InsertarNotificacion('Notificación A', 'Texto A', '12345678', 1);
CALL InsertarNotificacion('Notificación B', 'Texto B', '87654321', 1);

call ObtenerUsuariosConRoles();
call ObtenerMiembroPorDocumentoMovil('12345678');
call ObtenerAsistenciasConMiembroUsuarioYServicio();

