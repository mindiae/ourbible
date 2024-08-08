# Maintainer: Mindia Edisherashvili <mindia.e@aol.com>
pkgname=ourbible
pkgver=0.1.0
pkgrel=1
epoch=
pkgdesc="bible viewer app with web interface uses MyBible android app's modules"
arch=('x86_64')
url="https://github.com/mindiae/ourbible"
license=('LGPL' 'GPL')
depends=('sqlite' 'go')
makedepends=('git')
source=("https://github.com/mindiae/ourbible/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('23ddaf74196e440999fff004fbd6b7b697f40042afbd2edd7ae10a2a72eb043e')

build() {
  cd "$srcdir/$pkgname-$pkgver"
  go build -o build/$pkgname
}

package() {
  cd "$srcdir/$pkgname-$pkgver"
  install -Dm755 "build/$pkgname" "$pkgdir/usr/local/bin/$pkgname"
  local target_dir="/usr/local/share/$pkgname"

  install -d "$pkgdir/$target_dir/assets"
  mv "assets/"* "$pkgdir/$target_dir/assets"
  install -d "$pkgdir/$target_dir/database"
  mv "database/"* "$pkgdir/$target_dir/database"
  install -d "$pkgdir/$target_dir/views"
  mv "views/"* "$pkgdir/$target_dir/views"
}
