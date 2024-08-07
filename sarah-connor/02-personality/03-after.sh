#!/bin/bash
set -o allexport; source 03-settings.env; set +o allexport
go run main.go

: <<'COMMENT'
En utilisant qwen2 et le même context que pour Ollama
Et en modifiant TOP_K et TOP_P (pour le focus)
je continue à délirer
il faut donc changer le contexte
COMMENT
