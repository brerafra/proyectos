#!/bin/bash

REPO_DIR=~/desarrollo/aprendizaje/go/proyectos
cd "$REPO_DIR" || exit 1

FECHA=$(date "+%Y-%m-%d %H:%M:%S")

#Hacer pull antes de subir cambios
echo "📅 Haciendo pull del repositorio remoto..."
git pull --rebase origin main

#Verificar cambios
CHANGES=$(git status --porcelain)

if [[ -n "$CHANGES" ]]; then
    echo "📆 Cambios detectados:"
    echo "$CHANGES"

    git add .
    git commit -m "Actualización automática: $FECHA"
    git push origin main

    echo "✅ Cambios subidos correctamente."
else
    echo "🔄 No hay cambios para subir."
fi