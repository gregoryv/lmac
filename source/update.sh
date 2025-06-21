#!/bin/bash
#
# NOTE! DO NOT RUN THIS SCRIPT MORE THAN ONCE A DAY
#
# See limitation notice on
# https://regauth.standards.ieee.org/standards-ra-web/pub/view.html

if [ $(find mal.csv -mtime +1) ]; then
    echo "update mal.csv"
    wget -o mal.csv https://standards-oui.ieee.org/oui/oui.csv
fi

if [ $(find mam.csv -mtime +1) ]; then
    echo "update mam.csv"
    wget -o mam.csv https://standards-oui.ieee.org/oui28/mam.csv
fi

if [ $(find mas.csv -mtime +1) ]; then
    echo "update mas.csv"
    wget -o mas.csv https://standards-oui.ieee.org/oui36/oui36.csv
fi

