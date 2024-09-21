# OurBible GUI Bible viewer which also comes with web interface

## Installation:

On ArchLinux

```bash
mkdir ourbible
cd ourbible
wget https://github.com/mindiae/ourbible/raw/main/PKGBUILD
makepkg
sudo pacman -U ourbible-0.10.0-0-x86_64.pkg.tar.zst
```

On Windows:

You have to have mingw64 installed for gcc and put into PATH variable

For creating windows installation file you have to have innosetup insalled

```bash
git clone https://github.com/mindiae/ourbible.git
cd ourbible
go build -o build/ourbible.exe -ldflags "-H=windowsgui" ./cmd/webview/
./build/ourbible.exe
```

open innosetup.iss with innosetup compiler and compile the program
installation file will be placed inside output folder

## Running without installation

If you want to just try it out or you are not on ArchLinux

```bash
git clone https://github.com/mindiae/ourbible.git
cd ourbible
go run ./cmd
```

## Running from web interface

Just open http://localhost:42069 on browser after running the application
