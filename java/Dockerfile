# Download datadog java agent
FROM curlimages/curl as curl
RUN curl -L https://dtdg.co/latest-java-tracer --output /tmp/dd-java-agent.jar

# Build the jar
FROM openjdk:19-jdk as build
WORKDIR /workspace/app

COPY mvnw .
COPY .mvn .mvn
COPY pom.xml .
COPY src src

RUN ./mvnw package -DskipTests
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)

# Build the app container
FROM gcr.io/distroless/java17-debian11
WORKDIR /app

COPY --from=curl /tmp/dd-java-agent.jar dd-java-agent.jar
COPY --from=build /workspace/app/target/helloworld-0.0.1-SNAPSHOT.jar app.jar
COPY --from=datadog/serverless-init:beta3 /datadog-init /app/datadog-init

ENV DD_SERVICE=datadog-demo-run-java
ENV DD_ENV=datadog-demo
ENV DD_VERSION=1

ENTRYPOINT ["/app/datadog-init"]
CMD ["/usr/bin/java", "-javaagent:/app/dd-java-agent.jar", "-jar", "/app/app.jar"]
