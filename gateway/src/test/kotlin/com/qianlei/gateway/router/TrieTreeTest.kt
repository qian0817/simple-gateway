package com.qianlei.gateway.router

import kotlin.test.Test
import kotlin.test.assertEquals

/**
 *
 * @author qianlei
 */
class TrieTreeTest {
    @Test
    fun testTrieTree() {
        val tree = TrieTree<String>()
        tree.add(listOf("a", "b"), "ab")
        tree.add(listOf("a", "b", "c"), "abc")
        tree.add(listOf("a", "*"), "a*")
        tree.add(listOf("*"), "*")
        tree.add(listOf("b", "a"), "ba")

        assertEquals(tree.get(listOf("a", "b")), "ab")
        assertEquals(tree.get(listOf("a", "b", "c")), "abc")
        assertEquals(tree.get(listOf("a", "b", "c", "d")), "a*")
        assertEquals(tree.get(listOf("b", "a")), "ba")
        assertEquals(tree.get(listOf("c")), "*")

        assertEquals(tree.del(listOf("a", "b")), "ab")
        assertEquals(tree.del(listOf("a", "b", "c")), "abc")
        assertEquals(tree.del(listOf("a", "*")), "a*")
        assertEquals(tree.del(listOf("*")), "*")
        assertEquals(tree.del(listOf("b", "a")), "ba")
    }
}