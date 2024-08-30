# OurBible GUI bible viewer which also comes with Web interface

## Installation:

On ArchLinux

```bash
mkdir ourbible
cd ourbible
wget https://github.com/mindiae/ourbible/raw/main/PKGBUILD
makepkg
sudo pacman -U ourbible-0.10.0-0-x86_64.pkg.tar.zst
```

## Running without installation

If you want to just try it out or you are not on ArchLinux

```bash
git clone https://github.com/mindiae/ourbible.git
cd ourbible
go run ./cmd
```

## Running from web interface

Just open http://localhost:42069 on browser after running the application
