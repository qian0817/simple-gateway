package com.qianlei.gateway.router

import java.util.concurrent.ConcurrentHashMap

/**
 * Trie Tree 实现
 * @author qianlei
 */
class TrieTree<T> {
    private val root = TrieTreeNode<T>()

    fun add(path: List<String>, value: T) {
        var node = root
        for (i in path.indices) {
            val subPath = path[i]
            if (subPath == "*") {
                val newNode = TrieTreeNode(value)
                node.child["*"] = newNode
                node = newNode
            } else {
                val child = node.child[subPath]
                if (child == null) {
                    val newNode = TrieTreeNode<T>()
                    node.child[subPath] = newNode
                    node = newNode
                } else {
                    node = child
                }
            }
            if (i == path.size - 1) {
                node.value = value
            }
        }
    }

    fun del(path: List<String>): T? {
        var ret: T? = null
        var node = root
        for (i in path.indices) {
            val subPath = path[i]
            if (subPath == "*") {
                return node.child.remove("*")?.value
            } else {
                val child = node.child[subPath] ?: return null
                node = child
            }
            if (i == path.size - 1) {
                ret = node.value
                node.value = null
            }
        }
        return ret
    }

    fun get(path: List<String>): T? {
        var ret: T? = null
        var node = root
        node.child["*"]?.let { ret = it.value }
        for (i in path.indices) {
            val subPath = path[i]
            node.child["*"]?.let { ret = it.value }
            node = node.child[subPath] ?: return ret
            if (i == path.size - 1 && node.value != null) {
                ret = node.value
            }
        }
        return ret
    }

    class TrieTreeNode<T>(
        var value: T? = null,
        val child: MutableMap<String, TrieTreeNode<T>> = ConcurrentHashMap()
    )
}