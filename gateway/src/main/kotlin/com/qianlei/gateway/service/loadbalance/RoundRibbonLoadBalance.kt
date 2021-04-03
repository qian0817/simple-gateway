package com.qianlei.gateway.service.loadbalance

import com.qianlei.gateway.config.Node
import io.vertx.core.http.HttpServerRequest
import java.util.concurrent.atomic.AtomicInteger

/**
 * @author qianlei
 */
class RoundRibbonLoadBalance(private val nodes: List<Node>) : ILoadBalance {
    private val index = AtomicInteger(0)

    override fun getNodeAfterLoadBalance(request: HttpServerRequest): Node {
        val i = index.getAndUpdate { index -> (index + 1) % nodes.size }
        return nodes.getOrNull(i) ?: nodes.last()
    }
}