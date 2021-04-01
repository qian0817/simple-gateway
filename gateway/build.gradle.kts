import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.4.31"
}

group = "com.qianlei"
version = "0.0.1"

repositories {
    mavenCentral()
    jcenter()
}

dependencies {
    implementation("org.jetbrains.kotlin", "kotlin-reflect")
    implementation("org.jetbrains.kotlin", "kotlin-stdlib-jdk8")
    testImplementation("org.jetbrains.kotlin", "kotlin-test-junit5")
    testImplementation("org.junit.jupiter", "junit-jupiter-engine", "5.7.0")
    implementation("com.fasterxml.jackson.module", "jackson-module-kotlin", "2.12.2")
    implementation("io.etcd", "jetcd-core", "0.5.4")
    implementation("io.ktor", "ktor-server-core", "1.5.2")
    implementation("io.ktor", "ktor-server-netty", "1.5.2")
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
