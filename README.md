Disclaimer: This is a very early WIP project. As such, the structure of the repository and code may be subject to large changes in the future. The code here is not intended to be used for any project other than [PotionDB](https://github.com/AndreRijo/potionDB) and [PotionDB's TPC-H Client](https://github.com/AndreRijo/TPCH-Client) at this phase.

# PotionDB SQL

A (very WIP) implementation of the SQL language for PotionDB.
Most of the current work has been done in supporting specification for PotionDB's views through SQL.
In the future, it is intended for it to be possible to define tables, queries, updates and possibly indexes.
The current work already includes a Parser, a Lexer and also an Interpreter for a (very limited) subset of SQL.

This repository is required by both [PotionDB](https://github.com/AndreRijo/potionDB) and [PotionDB's TPC-H Client](https://github.com/AndreRijo/TPCH-Client) as both of those have been used to test the current implementation of this library.

Our SQL Parser and Lexer builds upon the [Antlr](https://www.antlr.org) library.