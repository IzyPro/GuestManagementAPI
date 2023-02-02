FROM maven:3.6.3-jdk-11-slim AS BUILD

COPY . /opt/app
WORKDIR /opt/app
RUN mvn install

FROM openjdk:11-jre-slim
ENV JAR='guestlist-1.0-SNAPSHOT-jar-with-dependencies.jar'
COPY --from=BUILD /opt/app/target/classes /opt/app/classes
COPY --from=BUILD /opt/app/target/${JAR} /opt/app/${JAR}
WORKDIR /opt/app

ENTRYPOINT ["/bin/sh", "-c", "java -jar ${JAR}"]
