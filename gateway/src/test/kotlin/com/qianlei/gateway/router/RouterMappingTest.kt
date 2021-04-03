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
            routerMapping.addRouter(Router("", path = ""))
            routerMapping.addRouter(Router("", path = "/a"))
            routerMapping.addRouter(Router("", path = "/a/b"))
            routerMapping.addRouter(Router("", path = "/*"))
            routerMapping.addRouter(Router("", path = "/a/*"))

            assertThrows<IllegalArgumentException> { routerMapping.addRouter(Router("", path = "//a")) }
            assertThrows<IllegalArgumentException> { routerMapping.addRouter(Router("", path = "/*/a")) }
            assertThrows<IllegalArgumentException> { routerMapping.addRouter(Router("", path = "/a/**")) }
        }
    }
}