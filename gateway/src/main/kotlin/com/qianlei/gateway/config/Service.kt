package com.qianlei.gateway.config

import com.qianlei.gateway.constant.LoadBalanceType

/**
 *
 * @author qianlei
 */
data class Service(
    val type: LoadBalanceType = LoadBalanceType.ROUND_RIBBON,
    val nodes: List<Node> = emptyList(),
) {
    fun getNode(): Node {
        return nodes[0]
    }
}