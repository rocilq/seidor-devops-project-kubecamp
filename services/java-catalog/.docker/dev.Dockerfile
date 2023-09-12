FROM mcr.microsoft.com/vscode/devcontainers/java:17-bullseye

# Install Gradle
RUN curl -s "https://get.sdkman.io" | bash
RUN bash -c "source /usr/local/sdkman/bin/sdkman-init.sh && sdk install gradle 7.2"

WORKDIR /app

LABEL lang="java"
LABEL version="17"
LABEL environment="dev"

CMD ["./gradlew", "bootRun"]


