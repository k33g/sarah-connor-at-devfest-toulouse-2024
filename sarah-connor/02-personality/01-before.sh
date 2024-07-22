#!/bin/bash
set -o allexport; source 01-settings.env; set +o allexport
go run main.go

: <<'COMMENT'
Avec ces paramètres, le modèle raconte un peu n'importe quoi
COMMENT
