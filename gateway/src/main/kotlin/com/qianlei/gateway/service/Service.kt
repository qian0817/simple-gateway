package com.qianlei.gateway.service

import com.qianlei.gateway.config.Node
import com.qianlei.gateway.constant.LoadBalanceType
import com.qianlei.gateway.constant.LoadBalanceType.*
import com.qianlei.gateway.service.loadbalance.HashLoadBalance
import com.qianlei.gateway.service.loadbalance.ILoadBalance
import com.qianlei.gateway.service.loadbalance.RandomLoadBalance
import com.qianlei.gateway.service.loadbalance.RoundRibbonLoadBalance
import kotlinx.serialization.Serializable

/**
 *
 * @author qianlei
 */
@Serializable
data class Service(
    val type: LoadBalanceType = ROUND_RIBBON,
    val nodes: List<Node> = emptyList(),
    val hashOn: String = ""
) {
    fun loadBalance() = createLoadBalance()

    private fun createLoadBalance(): ILoadBalance {
        return when (type) {
            ROUND_RIBBON -> RoundRibbonLoadBalance(nodes)
            HASH -> HashLoadBalance(nodes, hashOn)
            RANDOM -> RandomLoadBalance(nodes)
        }
    }
}