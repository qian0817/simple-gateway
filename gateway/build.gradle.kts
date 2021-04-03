import com.github.jengelman.gradle.plugins.shadow.tasks.ShadowJar
import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.4.31"
    application
    id("com.github.johnrengelman.shadow") version "6.1.0"
}

group = "com.qianlei"
version = "0.0.1"

repositories {
    mavenCentral()
    jcenter()
}

dependencies {
    implementation(kotlin("reflect"))
    implementation(kotlin("stdlib-jdk8"))
    testImplementation(kotlin("test-junit5"))
    testImplementation("org.junit.jupiter", "junit-jupiter-engine", "5.7.0")

    implementation("com.fasterxml.jackson.module", "jackson-module-kotlin", "2.12.2")
    implementation("io.etcd", "jetcd-core", "0.5.4")

    implementation(platform("io.vertx:vertx-stack-depchain:4.0.3"))
    implementation("io.vertx", "vertx-web")
    implementation("io.vertx", "vertx-web-client")
    implementation("io.vertx", "vertx-lang-kotlin-coroutines")
    implementation("io.vertx", "vertx-lang-kotlin")
    testImplementation("io.vertx", "vertx-junit5")

    implementation("ch.qos.logback", "logback-classic", "1.2.3")
    implementation("io.github.microutils", "kotlin-logging", "2.0.6")
}

java {
    sourceCompatibility = JavaVersion.VERSION_11
    targetCompatibility = JavaVersion.VERSION_11
}

tasks.withType<KotlinCompile> {
    kotlinOptions {
        freeCompilerArgs = listOf("-Xjsr305=strict")
        jvmTarget = "11"
    }
}

tasks.withType<Test> {
    useJUnitPlatform()
}

val mainVerticleName = "com.qianlei.gateway.MainVerticle"
val launcherClassName = "io.vertx.core.Launcher"

val watchForChange = "src/**/*"
val doOnChange = "${projectDir}/gradlew classes"

application {
    mainClass.set("io.vertx.core.Launcher")
}

tasks.withType<ShadowJar> {
    archiveClassifier.set("fat")
    manifest {
        attributes(mapOf("Main-Verticle" to mainVerticleName))
    }
    mergeServiceFiles()
}


tasks.withType<JavaExec> {
    args = listOf(
        "run",
        mainVerticleName,
        "--redeploy=$watchForChange",
        "--launcher-class=$launcherClassName",
        "--on-redeploy=$doOnChange"
    )
}
