# An Arch Linux Package bible viewer with web interface

## installation:

```bash
git clone https://github.com/mindiae/ourbible.git
mkdir ourbible-package
cp ourbible/PKGBUILD ourbible-package
cd ourbible-package
makepkg
sudo pacman -U ourbible-0.1.0-x86_64.pkg.tar.zst
```
