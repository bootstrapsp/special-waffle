### Development setup

0. Java is required.
1. Install [Maven](https://maven.apache.org/install.html).
2. Clone repository
3. Launch below code in command line (Windows command line 'cmd') to add jar library to maven local repository
```sh
mvn install:install-file -Dfile=libs/thingworx-java-sdk-6.2.0-javadoc.jar -DgroupId=com.thingworx  -DartifactId=thingworx-java-sdk -Dversion=6.2.0 -Dpackaging=jar
mvn install:install-file -Dfile=libs/thingworx-java-sdk-6.2.0.jar -DgroupId=com.thingworx  -DartifactId=thingworx-java-sdk -Dversion=6.2.0 -Dpackaging=jar
```

4. Open & Develop! In order to autogenerate proto files **compile** the project.