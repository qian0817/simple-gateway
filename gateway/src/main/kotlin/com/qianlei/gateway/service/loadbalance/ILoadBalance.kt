package com.qianlei.gateway.service.loadbalance

import com.qianlei.gateway.config.Node
import io.vertx.core.http.HttpServerRequest

/**
 *
 * @author qianlei
 */
interface ILoadBalance {
    fun getNodeAfterLoadBalance(request: HttpServerRequest): Node
}