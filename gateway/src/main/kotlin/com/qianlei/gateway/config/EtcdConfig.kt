package com.qianlei.gateway.config

/**
 *
 * @author qianlei
 */
data class EtcdConfig(
    val endpoints: Array<String> = arrayOf("http://127.0.0.1:2379"),
) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as EtcdConfig

        if (!endpoints.contentEquals(other.endpoints)) return false

        return true
    }

    override fun hashCode(): Int {
        return endpoints.contentHashCode()
    }
}