package com.qianlei.gateway.config

/**
 *
 * @author qianlei
 */
data class EtcdConfig(
    val host: String = "127.0.0.1",
    val port: Int = 2379
)