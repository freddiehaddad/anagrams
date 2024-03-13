# Multithreaded Group Anagrams

Program that converts a group of words into groups of anagrams.

The program works by reading words from the input array and creating a thread
per word that tokenizes the input. A token contains the original word as well as
a sorted copy of the word. The sorted copy serves as a key for the map. The idea
is that all anagrams sorted will have the same letters in the same order. The
token is written to a channel for which another thread is listening. Each word
received by the channel is organized in a map based on the token's key.

Once all words have been processed the map is converted to an array of string
arrays where each element is all the words that are anagrams.

## Example

Input

```text
["eat", "tea", "tan", "ate", "nat", "bat"]
```

Intermediary Representation

```text
 Key   | Value
-------+-----------------------
 "aet" | ["eat", "tea", "ate"]
 "ant" | ["tan", "nat"]
 "abt" | ["bat"]
```

Result

```text
[["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

## Thread Flow Diagram

```text
+-------------------------------+
| Group()                     G |        +--------------------------+
|                               |        | Tokenization           G |
|                               +------->|                          |----+
| ["eat", "tea", ..., "bat"]    |        |                          |  G |
+-------------------------------+        | "eat"                    |    |----+
                                         +------------+-------------+    |  G |
                                              | "tea"                    |    |
                                              +------------+-------------+    |
+---------------------------------+                | "bat"                    |
| Map generation                G |                +------------+-------------+
|                                 |                             |
| // token                        |<----------------------------+
| { key: "aet", value: "eat" }    |
+---------------------------------+

G: Go Routine
```
