@echo off

mkdir static\js
curl -o static\js\alpine.js https://cdn.jsdelivr.net/npm/alpinejs@3.14.3/dist/cdn.js
curl -o static\js\persist.js https://cdn.jsdelivr.net/npm/@alpinejs/persist@3.14.3/dist/cdn.js
curl -o static\js\alpinejs-swipe.js https://unpkg.com/alpinejs-swipe@1.0.2/dist/cjs.js
curl -L -O https://github.com/mindiae/ourbible/releases/download/0.22.5/database.tar.gz
7z.exe x database.tar.gz
7z.exe x database.tar
del database.tar.gz
