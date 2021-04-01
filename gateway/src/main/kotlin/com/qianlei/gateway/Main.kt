package com.qianlei.gateway

import com.qianlei.gateway.config.ServerConfig
import io.ktor.application.*
import io.ktor.request.*
import io.ktor.response.*
import io.ktor.routing.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import java.net.URLDecoder
import kotlin.text.Charsets.UTF_8

/**
 *
 * @author qianlei
 */
fun main() {
    val config = ServerConfig()
    embeddedServer(Netty, config.port, config.host) {
        routing {
            route("{...}") {
                handle {
                    val path = URLDecoder.decode(call.request.path(), UTF_8)
                    val uri = URLDecoder.decode(call.request.uri, UTF_8)
                    call.respondText { "path:$path\n uri:$uri" }
                }
            }
        }
    }.start()
}