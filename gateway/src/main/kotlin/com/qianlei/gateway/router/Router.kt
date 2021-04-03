package com.qianlei.gateway.router

import com.qianlei.gateway.config.PluginConfig
import com.qianlei.gateway.config.Service
import com.qianlei.gateway.constant.HttpMethod

/**
 *
 * @author qianlei
 */
data class Router(
    val name: String,
    val labels: Map<String, String> = mapOf(),
    val version: String? = null,
    val description: String? = null,
    val publish: Boolean = false,
    val host: String? = null,
    val path: String = "/*",
    val methods: Set<HttpMethod> = emptySet(),
    val service: Service = Service(),
    val plugins: List<PluginConfig> = listOf()
)