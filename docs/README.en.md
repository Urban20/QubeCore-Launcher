# QubeCore Launcher

Minimal, portable command-line launcher for Minecraft Java Edition. Written in Go.

## Requirements

- Go 1.21+
- Java (configurable path)

## Build

```bash
git clone https://github.com/Urban20/QubeCore-Launcher.git
cd QubeCore-Launcher/src/QubeCore/launcher/
go build qubecore .
```

## Usage

```bash
./qubecore
```

On first run, the version manifest is downloaded automatically.

## Configuration

A `config.ini` is generated on first run:


`Ruta` defaults to the Java binary in PATH. Edit manually and restart to apply changes.
