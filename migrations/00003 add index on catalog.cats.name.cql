CREATE CUSTOM INDEX cat_name_idx ON catalog.cats (name)
    USING 'org.apache.cassandra.index.sasi.SASIIndex'