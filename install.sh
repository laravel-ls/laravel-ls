#!/bin/bash

REPOOWNER="laravel-ls"
REPONAME="laravel-ls"
RELEASETAG=$(curl -s "https://api.github.com/repos/$REPOOWNER/$REPONAME/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

KERNEL=$(uname -s 2>/dev/null || /usr/bin/uname -s)
case ${KERNEL} in
    "Linux"|"linux")
        KERNEL="linux"
        ;;
    *)
        echo "Error: OS '${KERNEL}' not supported" > /dev/stderr
        exit 1
        ;;
esac

MACHINE=$(uname -m 2>/dev/null || /usr/bin/uname -m)
case ${MACHINE} in
    x86_64)
        MACHINE="amd64"
        ;;
    *)
        echo "Error: Your architecture (${MACHINE}) is not currently supported" > /dev/stderr
        exit 1
        ;;
esac

BINNAME="${BINNAME:-laravel-ls}"
BINDIR="${BINDIR:-/usr/local/bin}"
ASSETNAME=laravel-ls-${RELEASETAG}-${KERNEL}-${MACHINE}
URL="https://github.com/$REPOOWNER/$REPONAME/releases/download/${RELEASETAG}/${ASSETNAME}"

echo -e "Downloading version $RELEASETAG from $URL\n"
curl -q --fail --location --progress-bar --output "${ASSETNAME}" "$URL"
sudo install $ASSETNAME $BINDIR/$BINNAME
rm $ASSETNAME
echo "Installation of version $RELEASETAG complete!"
