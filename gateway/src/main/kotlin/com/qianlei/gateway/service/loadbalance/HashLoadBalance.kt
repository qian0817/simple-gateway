package com.qianlei.gateway.service.loadbalance

import com.qianlei.gateway.config.Node
import io.vertx.core.http.HttpServerRequest
import kotlin.math.abs

/**
 * @author qianlei
 */
class HashLoadBalance(private val nodes: List<Node>, private val hashOn: String) : ILoadBalance {

    override fun getNodeAfterLoadBalance(request: HttpServerRequest): Node {
        val hash = when {
            hashOn == "uri" -> request.uri().hashCode()
            hashOn == "server_addr" -> request.remoteAddress().hashCode()
            hashOn == "query_string" -> request.query()?.hashCode()
            hashOn == "cookie" -> request.cookieMap()?.hashCode()
            hashOn.startsWith("header_") -> request.getHeader(hashOn.substring(7))?.hashCode()
            else -> throw IllegalArgumentException("unsupported hash_on")
        } ?: return nodes.random()
        return nodes[abs(hash) % nodes.size]
    }
}