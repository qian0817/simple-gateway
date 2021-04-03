package com.qianlei.gateway.etcd

import com.qianlei.gateway.config.EtcdConfig
import com.qianlei.gateway.config.Node
import com.qianlei.gateway.constant.LoadBalanceType
import com.qianlei.gateway.router.Router
import com.qianlei.gateway.service.Service
import io.etcd.jetcd.ByteSequence
import io.etcd.jetcd.Client
import io.etcd.jetcd.Watch
import io.etcd.jetcd.options.GetOption
import io.etcd.jetcd.options.WatchOption
import io.etcd.jetcd.watch.WatchEvent
import io.etcd.jetcd.watch.WatchResponse
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import mu.KotlinLogging
import kotlin.text.Charsets.UTF_8

/**
 *
 * @author qianlei
 */
fun main() {
    EtcdClient(EtcdConfig()).use { client ->
        val router1 = Router(
            "web1",
            path = "/web1/*",
            service = Service(
                type = LoadBalanceType.HASH,
                nodes = listOf(Node("localhost", 9001), Node("localhost", 9002)),
                hashOn = "uri"
            )
        )
        val router2 = Router("web2", path = "/web2/*", service = Service(nodes = listOf(Node("localhost", 9002))))
        client.addRouter(router1)
        client.addRouter(router2)
    }
}

@Suppress("UnstableApiUsage")
class EtcdClient(config: EtcdConfig) : AutoCloseable {
    private val logger = KotlinLogging.logger { }
    private val client = Client.builder().endpoints(*config.endpoints).keepaliveWithoutCalls(true).build()

    fun addRouter(router: Router) {
        client.kvClient.use { client ->
            client.put(
                ByteSequence.from("gateway:router:${router.name}", UTF_8),
                ByteSequence.from(Json.encodeToString(router).toByteArray())
            ).get()
        }
    }

    fun getAllRouter(): List<Router> {
        return client.kvClient.use { client ->
            val key = ByteSequence.from("gateway:router:", UTF_8)
            val response = client.get(key, GetOption.newBuilder().withPrefix(key).build()).get()
            response.kvs.map { kv -> Json.decodeFromString(kv.value.bytes.decodeToString()) }
        }
    }

    fun watchRouter(onPut: suspend (router: Router) -> Unit, onDelete: suspend (router: Router) -> Unit) {
        val client = client.watchClient
        val key = ByteSequence.from("gateway:router:", UTF_8)
        val watchOption = WatchOption.newBuilder().withPrefix(key).withPrevKV(true).build()
        client.watch(key, watchOption, object : Watch.Listener {
            override fun onNext(response: WatchResponse) {
                response.events.forEach { event ->
                    GlobalScope.launch {
                        when (event.eventType) {
                            WatchEvent.EventType.DELETE -> {
                                val value = event.prevKV.value.bytes.decodeToString()
                                val router = Json.decodeFromString<Router>(value)
                                onDelete(router)
                            }
                            WatchEvent.EventType.PUT -> {
                                val value = event.keyValue.value.bytes.decodeToString()
                                val router = Json.decodeFromString<Router>(value)
                                onPut(router)
                            }
                            WatchEvent.EventType.UNRECOGNIZED, null -> logger.warn { "etcd watch unrecognized" }
                        }
                    }.invokeOnCompletion { e -> e?.let { logger.error(e) { } } }
                }
            }

            override fun onError(e: Throwable) {
                logger.error(e) {}
            }

            override fun onCompleted() {}
        }
        )
    }

    override fun close() {
        client.close()
    }
}