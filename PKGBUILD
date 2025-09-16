# Maintainer: Doridian <git at doridian dot net>

# This should ideally be inside a pkgver() subroutine, but that is not possible
# as part of the version comes from the commit count since the latest tag
# so if you commit your current changes the PKGBUILD that would push it one tag further
# than it just calculated, so it would cause a perma-diff in git which is very suboptimal
latest_tag="$(git describe --tags --abbrev=0)"
commits_since_tag="$(git rev-list --count "${latest_tag}..HEAD")"
tag_suffix=''
if [ -n "$(git status --porcelain)" ]; then
  tag_suffix='.dev'
  commits_since_tag=$((commits_since_tag + 1))
fi

pkgname=backupmgr
pkgver="${latest_tag}.${commits_since_tag}${tag_suffix}"
pkgrel="1"
pkgdesc='Restic backup manager'
arch=('x86_64' 'arm64')
url='https://github.com/FoxDenHome/backupmgr.git'
license=('GPL-3.0-or-later')
makedepends=('git' 'go')
depends=()
source=(
  'config.example.json'
  'systemd/backupmgr-backup.service'
  'systemd/backupmgr-backup.timer'
  'systemd/backupmgr-backup.slice'
  'systemd/backupmgr-prune.service'
  'systemd/backupmgr-prune.timer'
  'systemd/backupmgr-prune.slice'
)
sha256sums=(
  'SKIP'
  'SKIP'
  'SKIP'
  'SKIP'
  'SKIP'
  'SKIP'
  'SKIP'
)

goldflags='' # Hidden tweak for source-ing this file

build() {
  cd "${startdir}"
  go build -trimpath -ldflags "${goldflags} -X github.com/FoxDenHome/backupmgr/util.version=${pkgver} -X github.com/FoxDenHome/backupmgr/util.gitrev=$(git rev-parse HEAD)" -o "${srcdir}/backupmgr" ./cmd/backupmgr
}

package() {
  backup=('etc/backupmgr/config.json')
  cd "${srcdir}"
  mkdir -p "${pkgdir}/etc/backupmgr"
  install -Dm755 ./backupmgr "${pkgdir}/usr/bin/backupmgr"
  install -Dm600 ./config.example.json "${pkgdir}/etc/backupmgr/config.json"

  install -Dm644 ./backupmgr-backup.service "${pkgdir}/usr/lib/systemd/system/backupmgr-backup.service"
  install -Dm644 ./backupmgr-backup.timer "${pkgdir}/usr/lib/systemd/system/backupmgr-backup.timer"
  install -Dm644 ./backupmgr-backup.slice "${pkgdir}/usr/lib/systemd/system/backupmgr-backup.slice"
  install -Dm644 ./backupmgr-prune.service "${pkgdir}/usr/lib/systemd/system/backupmgr-prune.service"
  install -Dm644 ./backupmgr-prune.timer "${pkgdir}/usr/lib/systemd/system/backupmgr-prune.timer"
  install -Dm644 ./backupmgr-prune.slice "${pkgdir}/usr/lib/systemd/system/backupmgr-prune.slice"
}
