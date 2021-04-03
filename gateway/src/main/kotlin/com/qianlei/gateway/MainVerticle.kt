package com.qianlei.gateway

import com.qianlei.gateway.config.Node
import com.qianlei.gateway.config.ServerConfig
import com.qianlei.gateway.constant.LoadBalanceType
import com.qianlei.gateway.router.Router
import com.qianlei.gateway.router.RouterMapping
import com.qianlei.gateway.service.Service
import io.vertx.core.Handler
import io.vertx.core.http.HttpServerRequest
import io.vertx.ext.web.client.WebClient
import io.vertx.kotlin.coroutines.CoroutineVerticle
import io.vertx.kotlin.coroutines.await
import kotlinx.coroutines.launch

/**
 *
 * @author qianlei
 */
class MainVerticle : CoroutineVerticle() {
    private val routerMapping = RouterMapping()
    private val serverConfig = ServerConfig()
    private lateinit var client: WebClient

    override suspend fun start() {
        routerMapping.addRouter(
            Router(
                "web1",
                path = "/web1/*",
                service = Service(
                    type = LoadBalanceType.HASH,
                    nodes = listOf(Node("localhost", 9001), Node("localhost", 9002)),
                    hashOn = "uri"
                )
            ),
            Router("web2", path = "/web2/*", service = Service(nodes = listOf(Node("localhost", 9002)))),
        )
        client = WebClient.create(vertx)
        vertx.createHttpServer()
            .requestHandler(MyHandler())
            .listen(serverConfig.port, serverConfig.host)
    }

    inner class MyHandler : Handler<HttpServerRequest> {
        override fun handle(req: HttpServerRequest) {
            val body = req.body()
            val router = routerMapping.getRouter(req.path())
            if (router == null) {
                req.response()
                    .setStatusCode(404)
                    .end("path ${req.path()} not found")
                return
            }
            launch {
                val loadBalance = router.service.loadBalance
                val node = loadBalance.getNodeAfterLoadBalance(req)
                val request = client.request(req.method(), node.port, node.host, req.uri())
                req.headers().forEach { (name, value) -> request.putHeader(name, value) }

                val response = request.sendBuffer(body.await()).await()
                req.response()
                    .setStatusCode(response.statusCode())
                    .setStatusMessage(response.statusMessage())
                    .also { response.headers().forEach { (name, value) -> it.putHeader(name, value) } }
                    .end(response.body())

            }.invokeOnCompletion { e -> e?.printStackTrace() }
        }
    }
}