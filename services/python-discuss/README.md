# 💻 Lesson's Discussion Service

### Sección entidades

Este micro servicio solo tiene un recurso para listar las discusiones de un curso. Este recurso se almacena en la base de datos MongoDB en la colección entidades.

#### Donde el recurso tiene la siguiente estructura:

- discussionId: identificador unico de la discusión.
- lessionId: identificador unico de la lección.
- title (string): titulo de la discusión.
- content (string): contenido de la discusión.
- timestamp (datetime): fecha de creación de la discusión.

#### Endpoint

> GET /discussions/{lessonId} : Listar las discusiones de un curso para una lección
> POST /discussions/{lessonId} : Crear una discusión

---
