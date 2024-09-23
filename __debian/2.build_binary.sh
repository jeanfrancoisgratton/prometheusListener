#!/usr/bin/env bash

PKGDIR="prometheusListener-1.04.01-0_amd64"

mkdir -p ${PKGDIR}/opt/sbin ${PKGDIR}/DEBIAN ${PKGDIR}/etc/systemd/system ${PKGDIR}/etc
for i in control preinst prerm postinst postrm;do
  mv $i ${PKGDIR}/DEBIAN/
done

echo "Building binary from source"
cd ../src
go build -o ../__debian/${PKGDIR}/opt/sbin/prometheusSDlistener .
strip ../__debian/${PKGDIR}/opt/sbin/prometheusSDlistener
sudo chown 0:0 ../__debian/${PKGDIR}/opt/sbin/prometheusSDlistener

echo "Binary built. Now packaging..."
cd ../__debian/
cp ../prometheusSDlistener.service ${PKGDIR}/etc/systemd/system/prometheusSDlistener.service
chmod 644 ${PKGDIR}/etc/systemd/system/prometheusSDlistener.service
touch ${PKGDIR}/etc/prometheusSDlistener.json
dpkg-deb -b ${PKGDIR}
