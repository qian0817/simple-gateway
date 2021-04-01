package com.qianlei.gateway.config

data class ServerConfig(
    val port: Int = 80,
    val host: String = "0.0.0.0",
    val etcd: EtcdConfig = EtcdConfig()
)