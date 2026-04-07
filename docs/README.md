# QubeCore Launcher

[ENGLISH](https://github.com/Urban20/QubeCore-Launcher/blob/main/docs/README.en.md) 

Launcher minimalista y portable de línea de comandos para Minecraft Java Edition. Escrito en Go.

## Requisitos

- Go 1.21+
- Java (ruta configurable)

## Compilar

```bash
git clone https://github.com/Urban20/QubeCore-Launcher.git
cd QubeCore-Launcher/src/QubeCore/launcher/
go build qubecore .
```

## Uso

```bash
./qubecore
```

En el primer inicio, el listado de versiones se descarga automáticamente.

## Configuración

Se genera un `config.ini` en el primer inicio:


`Ruta` usa por defecto el binario de Java en el PATH. Edita manualmente y reinicia para aplicar cambios.
