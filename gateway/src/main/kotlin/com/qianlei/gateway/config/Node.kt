package com.qianlei.gateway.config

import kotlinx.serialization.Serializable

@Serializable
data class Node(
    val host: String,
    val port: Int,
)
