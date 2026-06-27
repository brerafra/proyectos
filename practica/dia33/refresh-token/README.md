# Refresh token

Permite mantener vivas las sesiones de usuario sin pedir credenciales
constantemente. el flujo consiste en usar un "access token" de corta 
duración (ej 15 min) para llamadas a API y un refresh token de larga duración (ej. 30 dias) guardado de forma segura, el cual se envia al servidor para generar un nuevo par de tokens cuando el primero expira.

