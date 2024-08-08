# Maintainer: Mindia Edisherashvili <mindia.e@aol.com>
pkgname=ourbible-git
pkgver() {
  cd "$srcdir/$pkgname"
  # Get the latest tag and the number of commits since that tag
  git describe --tags --long --dirty | sed 's/^v//;s/-/+/g'
}
pkgrel=1
epoch=
pkgdesc="bible viewer app with web interface uses MyBible android app's modules"
arch=('x86_64')
url="https://github.com/mindiae/ourbible" # Replace with the actual repository URL
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
source=("git+https://github.com/mindiae/ourbible.git") # Replace with the actual repository URL
noextract=()
sha256sums=('SKIP')

pkgver() {
  cd "$srcdir/$pkgname"
  git describe --tags --long --dirty | sed 's/^v//;s/-/+/g'
}

build() {
  cd "$srcdir/$pkgname"
  go build -o build/$pkgname
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm755 "build/$pkgname" "$pkgdir/usr/local/bin/$pkgname"
  local target_dir="/usr/local/share/$pkgname"

  install -d "$pkgdir/$target_dir/assets"
  mv "assets/"* "$pkgdir/$target_dir/assets"
  install -d "$pkgdir/$target_dir/database"
  mv "database/"* "$pkgdir/$target_dir/database"
  install -d "$pkgdir/$target_dir/views"
  mv "views/"* "$pkgdir/$target_dir/views"
}
