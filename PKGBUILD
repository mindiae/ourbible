# Maintainer: Mindia Edisherashvili <mindia.e@aol.com>
pkgname=ourbible
pkgver=0.15.0
pkgrel=0
epoch=
pkgdesc="Bible viewer app with GUI and web interface which uses MyBible Android app's modules"
arch=('x86_64')
url="https://github.com/mindiae/ourbible"
license=('GPL-3.0')
depends=('sqlite' 'go' 'gtk4' 'webkitgtk-6.0' 'ttf-font-awesome')
makedepends=('git')
source=("https://github.com/mindiae/ourbible/archive/refs/tags/$pkgver.tar.gz")
sha256sums=("SKIP")

build() {
  cd "$srcdir/$pkgname-$pkgver"
  go build -o build/$pkgname ./cmd/webview/
}

package() {
  cd "$srcdir/$pkgname-$pkgver"
  install -Dm755 "build/$pkgname" "$pkgdir/usr/local/bin/$pkgname"
  local target_dir="/usr/local/share/$pkgname"

  install -d "$pkgdir/usr/share/applications"
  mv "ourbible.desktop" "$pkgdir/usr/share/applications"
  install -d "$pkgdir/$target_dir/static"
  mv "static/"* "$pkgdir/$target_dir/static"
  install -d "$pkgdir/$target_dir/database"
  mv "database/"* "$pkgdir/$target_dir/database"
  mv "storage.sqlite3" "$pkgdir/$target_dir/"
}
