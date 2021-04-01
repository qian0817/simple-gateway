package com.qianlei.gateway.config


data class PluginConfig(
    val name: String,
    val enable: Boolean = false,
    val data: Map<String, Any>
)

