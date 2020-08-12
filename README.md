# iqueue

A simple 2-lock queue for intervals

Operates at about 110ns / op single-threaded on a 2.8GHz i7

This follows the michael+scott 1996 PODC approach using locking.

This is used by a range locking package as overflow from collisions in an interval tree.

