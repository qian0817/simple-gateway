package com.qianlei.gateway.router

import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock

/**
 * 路由映射
 * 使用 Trie Tree 树作为路由匹配的数据结构
 * 其中 add 操作和 del 操作需要进行加锁，get 操作无需加锁
 *
 * @author qianlei
 */
class RouterMapping {
    private val trieTree = TrieTree<Router>()
    private val mutex = Mutex()

    /**
     * 添加路由
     */
    suspend fun addRouter(vararg routers: Router) {
        routers.forEach { require(isPattern(it.path)) }
        mutex.withLock {
            routers.forEach { router ->
                val path = router.path
                val subPaths = path.split("/").filter { it != "" }
                trieTree.add(subPaths, router)
            }
        }
    }

    /**
     * 删除[paths]路径下的路由
     */
    suspend fun deleteRouter(vararg paths: String) {
        mutex.withLock {
            paths.forEach { path ->
                val subPaths = path.split("/").filter { it != "" }
                trieTree.del(subPaths)
            }
        }
    }

    /**
     * 获取[path]路径下的路由
     */
    fun getRouter(path: String): Router? {
        val subPaths = path.split("/").filter { it != "" }
        return trieTree.get(subPaths)
    }

    private fun isPattern(pattern: String): Boolean {
        if (pattern.isEmpty()) {
            return true
        }
        val index = pattern.indexOf('*')
        if (index != -1 && index != pattern.length - 1) {
            return false
        }
        for (i in 1 until pattern.length) {
            if (pattern[i] == '/' && pattern[i - 1] == '/') {
                return false
            }
        }
        if (index == pattern.length - 1 && pattern.length != 1) {
            return pattern[pattern.length - 2] == '/'
        }
        return true
    }
}