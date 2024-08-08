# Maintainer: Mindia Edisherashvili <mindia.e@aol.com>
pkgname=ourbible
pkgver=0.1.0
pkgrel=1
epoch=
pkgdesc="bible viewer app with web interface uses MyBible android app's modules"
arch=('x86_64')
url=""
license=('LGPL' 'GPL')
groups=()
depends=('sqlite' 'go')
makedepends=('git')
checkdepends=()
optdepends=()
provides=()
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=(*)
noextract=()
sha256sums=('SKIP')
validpgpkeys=()

build() {
  cd "$srcdir"
  go build -o build/$pkgname
}

package() {
  cd "$srcdir"
  install -Dm755 "build/$pkgname" "$pkgdir/usr/local/bin/$pkgname"
  local target_dir="/usr/local/share/$pkgname"

  install -d "$pkgdir/$target_dir/assets"
  mv "assets/"* "$pkgdir/$target_dir/assets"
  install -d "$pkgdir/$target_dir/database"
  mv "database/"* "$pkgdir/$target_dir/database"
  install -d "$pkgdir/$target_dir/views"
  mv "views/"* "$pkgdir/$target_dir/views"

}
