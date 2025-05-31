Disclaimer: This is still a WIP project. As such, the structure of the repository and code may be subject to large changes in the future. The code here is not intended to be used by any project other than [PotionDB](https://github.com/AndreRijo/potionDB) at this phase.

# PotionDB SQL

A WIP implementation of the SQL language for PotionDB.
Most of the work done so far focuses on automatically specifying PotionDB's views through SQL.
In the future, it is intended for it to be possible to define tables, direct queries, updates and possibly indexes through SQL.
The current work already includes a Parser, a Lexer and also an Interpreter of SQL.

Our SQL Parser and Lexer builds upon the [Antlr](https://www.antlr.org) library.