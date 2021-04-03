package com.qianlei.gateway

import com.qianlei.gateway.config.ServerConfig
import com.qianlei.gateway.etcd.EtcdClient
import com.qianlei.gateway.router.RouterMapping
import io.vertx.core.Handler
import io.vertx.ext.web.Router
import io.vertx.ext.web.RoutingContext
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
    private lateinit var etcdClient: EtcdClient
    private lateinit var webClient: WebClient

    override suspend fun start() {
        webClient = WebClient.create(vertx)

        etcdClient = EtcdClient(serverConfig.etcd)
        routerMapping.addRouter(*etcdClient.getAllRouter().toTypedArray())
        etcdClient.watchRouter({ routerMapping.addRouter(it) }, { routerMapping.deleteRouter(it.name) })

        val server = vertx.createHttpServer()
        val router = Router.router(vertx)
        router.route().handler(RoutingHandler())
        server.requestHandler(router).listen(serverConfig.port, serverConfig.host)
    }

    override suspend fun stop() {
        etcdClient.close()
    }

    inner class RoutingHandler : Handler<RoutingContext> {
        override fun handle(context: RoutingContext) {
            val req = context.request()
            val router = routerMapping.getRouter(req.path())
            if (router == null) {
                req.response()
                    .setStatusCode(404)
                    .end("path ${req.path()} not found")
                return
            }
            launch {
                val loadBalance = router.service.loadBalance()
                val node = loadBalance.getNodeAfterLoadBalance(req)
                val request = webClient.request(req.method(), node.port, node.host, req.uri())
                req.headers().forEach { (name, value) -> request.putHeader(name, value) }

                val response = request.sendBuffer(context.body).await()
                req.response()
                    .setStatusCode(response.statusCode())
                    .setStatusMessage(response.statusMessage())
                    .also { response.headers().forEach { (name, value) -> it.putHeader(name, value) } }
                    .end(response.body())

            }.invokeOnCompletion { e -> e?.printStackTrace() }
        }
    }
}