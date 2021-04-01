package com.qianlei.gateway.config

import com.qianlei.gateway.constant.LoadBalanceType

/**
 *
 * @author qianlei
 */
data class ServiceConfig(
    val type: LoadBalanceType = LoadBalanceType.ROUND_RIBBON,
    val nodeList: List<Node> = emptyList(),
)