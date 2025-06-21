#!/bin/bash
# NOTE! DO NOT RUN THIS SCRIPT MORE THAN ONCE A DAY
# See limitation notice on
# https://regauth.standards.ieee.org/standards-ra-web/pub/view.html

cd source
if [ $(find ma-l.csv -mtime +1) ]; then
    echo "update ma-l.csv"
    wget -o source/ma-l.csv https://standards-oui.ieee.org/oui/oui.csv
fi

if [ $(find ma-m.csv -mtime +1) ]; then
    echo "update ma-m.csv"
    wget -o source/ma-m.csv https://standards-oui.ieee.org/oui28/mam.csv
fi

if [ $(find ma-s.csv -mtime +1) ]; then
    echo "update ma-s.csv"
    wget -o source/ma-s.csv https://standards-oui.ieee.org/oui36/oui36.csv
fi

