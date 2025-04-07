CREATE KEYSPACE IF NOT EXISTS catalog WITH REPLICATION =
    { 'class' : 'SimpleStrategy',
        'replication_factor' : '1'
        };