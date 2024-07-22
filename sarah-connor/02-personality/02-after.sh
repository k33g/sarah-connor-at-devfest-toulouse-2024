#!/bin/bash
set -o allexport; source 02-settings.env; set +o allexport
go run main.go

: <<'COMMENT'
Avec ces paramètres, le modèle utilise son contexte
Et délire beaucoup moins
si j'utilise qwen c'est bof
COMMENT
