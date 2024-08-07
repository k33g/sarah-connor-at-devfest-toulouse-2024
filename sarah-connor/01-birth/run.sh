#!/bin/bash
set -o allexport; source settings.env; set +o allexport
go run main.go


: <<'COMMENT'
Ce projet sert d'initialisation
COMMENT
