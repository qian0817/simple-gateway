package com.qianlei.gateway.service.loadbalance

import com.qianlei.gateway.config.Node
import io.vertx.core.http.HttpServerRequest

/**
 *
 * @author qianlei
 */
class RandomLoadBalance(private val nodes: List<Node>) : ILoadBalance {
    override fun getNodeAfterLoadBalance(request: HttpServerRequest): Node {
        return nodes.random()
    }
}