# Despliegue y Orquestación de Microservicios de una Plataforma Educativa

## 🎯 Objetivos

- Comprender la arquitectura de microservicios y cómo se comunican entre sí.
- Aprender a contenerizar aplicaciones desarrolladas en diferentes lenguajes mediante Docker.
- Familiarizarse con la automatización de procesos CI/CD usando Jenkins y SonarQube.
- Diseñar e implementar un chart de Helm para administrar el despliegue de la aplicación en un entorno de Kubernetes.

## ⛺️ Preparación

Antes de comenzar con los ejercicios, asegúrate de tener instalado y configurado:

- Docker y Docker Compose
- Jenkins y SonarQube (Tendrás acceso a `jenkins.aroldev.com` y `sonarqube.aroldev.com` respectivamente)
- Helm y Kubernetes

## 🥽 Ejercicios

1. **Dockerización de Microservicios**:
   Cada uno de los microservicios tiene una funcionalidad específica dentro de la aplicación:

   - Go: Autenticación y autorización.
   - Java: Catálogo de cursos y lecciones.
   - PHP: Control de progreso de cada usuario con respecto a las lecciones.
   - Python: Foro de discusión sobre las lecciones.

   Tu tarea es crear un `Dockerfile` para cada uno de estos microservicios. Asegúrate de que estén optimizados para producción.

2. **Integración Continua con Jenkins y SonarQube**:

   - Para cada microservicio, crea un `Jenkinsfile` que defina los pasos para correr tests (si los hay), verificar la calidad del código usando SonarQube y crear una imagen Docker para finalmente subirla a DockerHub.
   - Configura un webhook en tu repositorio de GitHub para que Jenkins inicie este proceso automáticamente cada vez que se haga un push a la rama `main`.

3. **Despliegue con Helm**:

   - Diseña y crea un chart de Helm que se acomode a la infraestructura mostrada en el `docker-compose.yml` que se encuentra en la carpeta `services`.
   - Utiliza Secrets y ConfigMaps según lo consideres necesario para manejar información sensible o de configuración.

4. **Documentación**:

   - **Dibuja un diagrama** que muestre cómo se relacionan y comunican los diferentes microservicios.
   - Describe el flujo de trabajo de Jenkins, incluyendo cómo se activa con webhooks, la integración con SonarQube y cómo se crea y almacena la imagen en DockerHub.
   - **Detalla cualquier paso manual** necesario o consideraciones especiales.
   - Proporciona una **guía para configurar y desplegar los microservicios**, incluyendo la configuración de Docker, Helm y cualquier otra herramienta utilizada. **Cómo se actualizan los servicios?**
   - Asegúrate de **incluir cualquier consideración especial o configuración necesaria** para diferentes entornos (desarrollo, prueba, producción).

## 💻 Cómo empezar y entrega

Para empezar a trabajar y entregar el ejercicio, sigue los siguientes pasos:

1. **Hacer Fork del Repositorio**:  
   Antes de empezar a trabajar en las soluciones, haz un fork del repositorio original. Esta acción creará una copia del repositorio en tu cuenta personal de GitHub.

2. **Clonar el Repositorio**:  
   Luego de hacer el fork, clona el repositorio a tu máquina local. Asegúrate de clonar tu fork y no el repositorio original. Puedes hacerlo utilizando el siguiente comando:

   ```bash
   $ git clone https://github.com/[tu_nombre_de_usuario]/[nombre_del_repositorio].git
   ```

   > **Nota**: Reemplaza `[tu_nombre_de_usuario]` con tu nombre de usuario de GitHub y `[nombre_del_repositorio]` con el nombre del repositorio que acabas de copiar.

3. **Trabajar en las Soluciones**:  
   Una vez clonado el repositorio, navega hacia él y empieza a trabajar en las soluciones de los ejercicios. Realiza los cambios y mejoras que consideres necesarios.

4. **Hacer Push de tus Cambios**:  
   Luego de haber completado tus soluciones, guarda tus cambios y súbelos a tu repositorio de GitHub. Utiliza los siguientes comandos:

   ```bash
   git add .
   git commit -m "<tu_nombre>"
   git push origin main
   ```

   **Nota**: Si estás trabajando en una rama diferente a `main`, reemplaza `main` con el nombre de tu rama.

5. **Crear un Pull Request**:  
   Una vez hayas hecho `push` de tus cambios, ve a la página de tu fork en GitHub. Haz clic en el botón `New pull request`. Asegúrate de que estás solicitando el pull request al repositorio original y no a tu fork. Completa la información requerida y haz clic en `Create pull request`.

¡Happy devoping! 🚢
