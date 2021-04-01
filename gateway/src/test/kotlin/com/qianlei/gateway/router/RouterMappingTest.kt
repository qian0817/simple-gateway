package com.qianlei.gateway.router

import kotlinx.coroutines.runBlocking
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.assertThrows

/**
 *
 * @author qianlei
 */
class RouterMappingTest {
    @Test
    fun testIllegalPath() {
        runBlocking {
            val routerMapping = RouterMapping()
            routerMapping.addRouter("" to Router.createEmptyRouter())
            routerMapping.addRouter("/" to Router.createEmptyRouter())
            routerMapping.addRouter("/a" to Router.createEmptyRouter())
            routerMapping.addRouter("/a/b" to Router.createEmptyRouter())
            routerMapping.addRouter("/*" to Router.createEmptyRouter())
            routerMapping.addRouter("/a/*" to Router.createEmptyRouter())

            assertThrows<IllegalArgumentException> { routerMapping.addRouter("//a" to Router.createEmptyRouter()) }
            assertThrows<IllegalArgumentException> { routerMapping.addRouter("/*/a" to Router.createEmptyRouter()) }
            assertThrows<IllegalArgumentException> { routerMapping.addRouter("/a/**" to Router.createEmptyRouter()) }
        }
    }
}