# OurBible GUI Bible viewer Written in go and web technologies

![ourbible-kjv-elb-doubleview-dark.png](https://raw.githubusercontent.com/mindiae/ourbible-pictures/refs/heads/main/ourbible-kjv-elb-doubleview-dark.png "Doubleview Dark KJV")

![ourbible-kjv-singleview-dark.png](https://github.com/mindiae/ourbible-pictures/blob/main/ourbible-kjv-singleview-dark.png?raw=true "Singleview Dark KJV")

![ourbible-kjv-singleview-light-booksview.png](https://github.com/mindiae/ourbible-pictures/blob/main/ourbible-kjv-singleview-light-booksview.png?raw=true "Singleview Light BooksView")


If you want to test application without installing it, first you need to install provided font-awesome fonts to your system


## Installation:

On ArchLinux:

```bash
wget https://github.com/mindiae/ourbible/releases/download/version/ourbible-version-x86_64.pkg.tar.zst
sudo pacman -U ourbible-version-x86_64.pkg.tar.zst
```
exact versions are in releases section



On Ubuntu and Debian based systems:
```bash
wget https://github.com/mindiae/ourbible/releases/download/version/ourbible_version_amd64.deb
sudo dpkg -i ourbible_version_amd64.deb
```

On windows:  
download ourbible-vVERSION.exe from releases and install it


Or install from aur
```bash
mkdir ourbible
cd ourbible
wget -O PKGBUILD "https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=ourbible"
makepkg
sudo pacman -U ourbible-version-x86_64.pkg.tar.zst
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
go run ./cmd/webview #or ./cmd if you want with web interface
```


## Running from web interface

if you run:
```bash
go run ./cmd
```
Just open http://localhost:42069 on browser after running the application
