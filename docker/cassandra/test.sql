DROP KEYSPACE IF EXISTS test;
CREATE KEYSPACE test
WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': 1 };
