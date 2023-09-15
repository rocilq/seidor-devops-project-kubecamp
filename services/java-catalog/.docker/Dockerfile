FROM gradle:8.2-jdk17 AS build

WORKDIR /app
COPY . .
RUN gradle clean build -x test

FROM openjdk:21-ea-17-slim-bullseye

WORKDIR /app
COPY --from=build /app/build/libs/KubeCamp.Catalog-1.0.0.jar /app.jar
EXPOSE 8080

ENTRYPOINT ["java", "-jar", "/app.jar"]