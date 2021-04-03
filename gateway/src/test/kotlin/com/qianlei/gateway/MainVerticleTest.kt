package com.qianlei.gateway

import io.vertx.core.Vertx
import io.vertx.junit5.VertxExtension
import io.vertx.junit5.VertxTestContext
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.extension.ExtendWith

/**
 *
 * @author qianlei
 */
@ExtendWith(VertxExtension::class)
class MainVerticleTest {
    @BeforeEach
    fun deployVerticle(vertx: Vertx, testContext: VertxTestContext) {
        vertx.deployVerticle(MainVerticle(), testContext.succeeding { testContext.completeNow() })
    }

    @Test
    fun verticleDeployed(vertx: Vertx, testContext: VertxTestContext) {
        testContext.completeNow()
    }
}