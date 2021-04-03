package com.qianlei.gateway.config

import kotlinx.serialization.Serializable

@Serializable
data class PluginConfig(
    val name: String,
    val enable: Boolean = false,
    val data: String
)

