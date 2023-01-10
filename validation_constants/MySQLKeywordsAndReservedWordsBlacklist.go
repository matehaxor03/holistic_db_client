package validation_constants

func GetMySQLKeywordsAndReservedWordsInvalidWords() map[string]interface{} {
value := make(map[string]interface{})
    value["A"] = nil
    value["ACCESSIBLE"] = nil
    value["ACCOUNT"] = nil
    value["ACTION"] = nil
    value["ACTIVE"] = nil
    value["ACTIVE;"] = nil
    value["ADD"] = nil
    value["ADMIN"] = nil
    value["ADMIN;"] = nil
    value["AFTER"] = nil
    value["AGAINST"] = nil
    value["AGGREGATE"] = nil
    value["ALGORITHM"] = nil
    value["ALL"] = nil
    value["ALTER"] = nil
    value["ALWAYS"] = nil
    value["ANALYSE"] = nil
    value["ANALYSE;"] = nil
    value["ANALYZE"] = nil
    value["AND"] = nil
    value["ANY"] = nil
    value["ARRAY"] = nil
    value["ARRAY;"] = nil
    value["AS"] = nil
    value["ASC"] = nil
    value["ASCII"] = nil
    value["ASENSITIVE"] = nil
    value["AT"] = nil
    value["ATTRIBUTE"] = nil
    value["ATTRIBUTE;"] = nil
    value["AUTHENTICATION"] = nil
    value["AUTHENTICATION;"] = nil
    value["AUTOEXTEND_SIZE"] = nil
    value["AUTO_INCREMENT"] = nil
    value["AVG"] = nil
    value["AVG_ROW_LENGTH"] = nil
    value["B"] = nil
    value["BACKUP"] = nil
    value["BEFORE"] = nil
    value["BEGIN"] = nil
    value["BETWEEN"] = nil
    value["BIGINT"] = nil
    value["BINARY"] = nil
    value["BINLOG"] = nil
    value["BIT"] = nil
    value["BLOB"] = nil
    value["BLOCK"] = nil
    value["BOOL"] = nil
    value["BOOLEAN"] = nil
    value["BOTH"] = nil
    value["BTREE"] = nil
    value["BUCKETS"] = nil
    value["BUCKETS;"] = nil
    value["BY"] = nil
    value["BYTE"] = nil
    value["C"] = nil
    value["CACHE"] = nil
    value["CALL"] = nil
    value["CASCADE"] = nil
    value["CASCADED"] = nil
    value["CASE"] = nil
    value["CATALOG_NAME"] = nil
    value["CHAIN"] = nil
    value["CHALLENGE_RESPONSE"] = nil
    value["CHALLENGE_RESPONSE;"] = nil
    value["CHANGE"] = nil
    value["CHANGED"] = nil
    value["CHANNEL"] = nil
    value["CHAR"] = nil
    value["CHARACTER"] = nil
    value["CHARSET"] = nil
    value["CHECK"] = nil
    value["CHECKSUM"] = nil
    value["CIPHER"] = nil
    value["CLASS_ORIGIN"] = nil
    value["CLIENT"] = nil
    value["CLONE"] = nil
    value["CLONE;"] = nil
    value["CLOSE"] = nil
    value["COALESCE"] = nil
    value["CODE"] = nil
    value["COLLATE"] = nil
    value["COLLATION"] = nil
    value["COLUMN"] = nil
    value["COLUMNS"] = nil
    value["COLUMN_FORMAT"] = nil
    value["COLUMN_NAME"] = nil
    value["COMMENT"] = nil
    value["COMMIT"] = nil
    value["COMMITTED"] = nil
    value["COMPACT"] = nil
    value["COMPLETION"] = nil
    value["COMPONENT"] = nil
    value["COMPRESSED"] = nil
    value["COMPRESSION"] = nil
    value["CONCURRENT"] = nil
    value["CONDITION"] = nil
    value["CONNECTION"] = nil
    value["CONSISTENT"] = nil
    value["CONSTRAINT"] = nil
    value["CONSTRAINT_CATALOG"] = nil
    value["CONSTRAINT_NAME"] = nil
    value["CONSTRAINT_SCHEMA"] = nil
    value["CONTAINS"] = nil
    value["CONTEXT"] = nil
    value["CONTINUE"] = nil
    value["CONVERT"] = nil
    value["CPU"] = nil
    value["CREATE"] = nil
    value["CROSS"] = nil
    value["CUBE"] = nil
    value["CUME_DIST"] = nil
    value["CURRENT"] = nil
    value["CURRENT_DATE"] = nil
    value["CURRENT_TIME"] = nil
    value["CURRENT_TIMESTAMP"] = nil
    value["CURRENT_USER"] = nil
    value["CURSOR"] = nil
    value["CURSOR_NAME"] = nil
    value["D"] = nil
    value["DATA"] = nil
    value["DATABASE"] = nil
    value["DATABASES"] = nil
    value["DATAFILE"] = nil
    value["DATE"] = nil
    value["DATETIME"] = nil
    value["DAY"] = nil
    value["DAY_HOUR"] = nil
    value["DAY_MICROSECOND"] = nil
    value["DAY_MINUTE"] = nil
    value["DAY_SECOND"] = nil
    value["DEALLOCATE"] = nil
    value["DEC"] = nil
    value["DECIMAL"] = nil
    value["DECLARE"] = nil
    value["DEFAULT"] = nil
    value["DEFAULT_AUTH"] = nil
    value["DEFINER"] = nil
    value["DEFINITION"] = nil
    value["DEFINITION;"] = nil
    value["DELAYED"] = nil
    value["DELAY_KEY_WRITE"] = nil
    value["DELETE"] = nil
    value["DENSE_RANK"] = nil
    value["DESC"] = nil
    value["DESCRIBE"] = nil
    value["DESCRIPTION"] = nil
    value["DESCRIPTION;"] = nil
    value["DES_KEY_FILE"] = nil
    value["DES_KEY_FILE;"] = nil
    value["DETERMINISTIC"] = nil
    value["DIAGNOSTICS"] = nil
    value["DIRECTORY"] = nil
    value["DISABLE"] = nil
    value["DISCARD"] = nil
    value["DISK"] = nil
    value["DISTINCT"] = nil
    value["DISTINCTROW"] = nil
    value["DIV"] = nil
    value["DO"] = nil
    value["DOUBLE"] = nil
    value["DROP"] = nil
    value["DUAL"] = nil
    value["DUMPFILE"] = nil
    value["DUPLICATE"] = nil
    value["DYNAMIC"] = nil
    value["E"] = nil
    value["EACH"] = nil
    value["ELSE"] = nil
    value["ELSEIF"] = nil
    value["EMPTY"] = nil
    value["ENABLE"] = nil
    value["ENCLOSED"] = nil
    value["ENCRYPTION"] = nil
    value["END"] = nil
    value["ENDS"] = nil
    value["ENFORCED"] = nil
    value["ENFORCED;"] = nil
    value["ENGINE"] = nil
    value["ENGINES"] = nil
    value["ENGINE_ATTRIBUTE"] = nil
    value["ENGINE_ATTRIBUTE;"] = nil
    value["ENUM"] = nil
    value["ERROR"] = nil
    value["ERRORS"] = nil
    value["ESCAPE"] = nil
    value["ESCAPED"] = nil
    value["EVENT"] = nil
    value["EVENTS"] = nil
    value["EVERY"] = nil
    value["EXCEPT"] = nil
    value["EXCHANGE"] = nil
    value["EXCLUDE"] = nil
    value["EXCLUDE;"] = nil
    value["EXECUTE"] = nil
    value["EXISTS"] = nil
    value["EXIT"] = nil
    value["EXPANSION"] = nil
    value["EXPIRE"] = nil
    value["EXPLAIN"] = nil
    value["EXPORT"] = nil
    value["EXTENDED"] = nil
    value["EXTENT_SIZE"] = nil
    value["F"] = nil
    value["FACTOR"] = nil
    value["FACTOR;"] = nil
    value["FAILED_LOGIN_ATTEMPTS"] = nil
    value["FAILED_LOGIN_ATTEMPTS;"] = nil
    value["FALSE"] = nil
    value["FAST"] = nil
    value["FAULTS"] = nil
    value["FETCH"] = nil
    value["FIELDS"] = nil
    value["FILE"] = nil
    value["FILE_BLOCK_SIZE"] = nil
    value["FILTER"] = nil
    value["FINISH"] = nil
    value["FINISH;"] = nil
    value["FIRST"] = nil
    value["FIRST_VALUE"] = nil
    value["FIXED"] = nil
    value["FLOAT"] = nil
    value["FLOAT4"] = nil
    value["FLOAT8"] = nil
    value["FLUSH"] = nil
    value["FOLLOWING"] = nil
    value["FOLLOWING;"] = nil
    value["FOLLOWS"] = nil
    value["FOR"] = nil
    value["FORCE"] = nil
    value["FOREIGN"] = nil
    value["FORMAT"] = nil
    value["FOUND"] = nil
    value["FROM"] = nil
    value["FULL"] = nil
    value["FULLTEXT"] = nil
    value["FUNCTION"] = nil
    value["G"] = nil
    value["GENERAL"] = nil
    value["GENERATED"] = nil
    value["GEOMCOLLECTION"] = nil
    value["GEOMCOLLECTION;"] = nil
    value["GEOMETRY"] = nil
    value["GEOMETRYCOLLECTION"] = nil
    value["GET"] = nil
    value["GET_FORMAT"] = nil
    value["GET_MASTER_PUBLIC_KEY"] = nil
    value["GET_MASTER_PUBLIC_KEY;"] = nil
    value["GET_SOURCE_PUBLIC_KEY"] = nil
    value["GET_SOURCE_PUBLIC_KEY;"] = nil
    value["GLOBAL"] = nil
    value["GRANT"] = nil
    value["GRANTS"] = nil
    value["GROUP"] = nil
    value["GROUPING"] = nil
    value["GROUPS"] = nil
    value["GROUP_REPLICATION"] = nil
    value["GTID_ONLY"] = nil
    value["GTID_ONLY;"] = nil
    value["H"] = nil
    value["HANDLER"] = nil
    value["HASH"] = nil
    value["HAVING"] = nil
    value["HELP"] = nil
    value["HIGH_PRIORITY"] = nil
    value["HISTOGRAM"] = nil
    value["HISTOGRAM;"] = nil
    value["HISTORY"] = nil
    value["HISTORY;"] = nil
    value["HOST"] = nil
    value["HOSTS"] = nil
    value["HOUR"] = nil
    value["HOUR_MICROSECOND"] = nil
    value["HOUR_MINUTE"] = nil
    value["HOUR_SECOND"] = nil
    value["I"] = nil
    value["IDENTIFIED"] = nil
    value["IF"] = nil
    value["IGNORE"] = nil
    value["IGNORE_SERVER_IDS"] = nil
    value["IMPORT"] = nil
    value["IN"] = nil
    value["INACTIVE"] = nil
    value["INACTIVE;"] = nil
    value["INDEX"] = nil
    value["INDEXES"] = nil
    value["INFILE"] = nil
    value["INITIAL"] = nil
    value["INITIAL;"] = nil
    value["INITIAL_SIZE"] = nil
    value["INITIATE"] = nil
    value["INITIATE;"] = nil
    value["INNER"] = nil
    value["INOUT"] = nil
    value["INSENSITIVE"] = nil
    value["INSERT"] = nil
    value["INSERT_METHOD"] = nil
    value["INSTALL"] = nil
    value["INSTANCE"] = nil
    value["INT"] = nil
    value["INT1"] = nil
    value["INT2"] = nil
    value["INT3"] = nil
    value["INT4"] = nil
    value["INT8"] = nil
    value["INTEGER"] = nil
    value["INTERSECT"] = nil
    value["INTERVAL"] = nil
    value["INTO"] = nil
    value["INVISIBLE"] = nil
    value["INVOKER"] = nil
    value["IO"] = nil
    value["IO_AFTER_GTIDS"] = nil
    value["IO_BEFORE_GTIDS"] = nil
    value["IO_THREAD"] = nil
    value["IPC"] = nil
    value["IS"] = nil
    value["ISOLATION"] = nil
    value["ISSUER"] = nil
    value["ITERATE"] = nil
    value["J"] = nil
    value["JOIN"] = nil
    value["JSON"] = nil
    value["JSON_TABLE"] = nil
    value["JSON_VALUE"] = nil
    value["JSON_VALUE;"] = nil
    value["K"] = nil
    value["KEY"] = nil
    value["KEYRING"] = nil
    value["KEYRING;"] = nil
    value["KEYS"] = nil
    value["KEY_BLOCK_SIZE"] = nil
    value["KILL"] = nil
    value["L"] = nil
    value["LAG"] = nil
    value["LANGUAGE"] = nil
    value["LAST"] = nil
    value["LAST_VALUE"] = nil
    value["LATERAL"] = nil
    value["LEAD"] = nil
    value["LEADING"] = nil
    value["LEAVE"] = nil
    value["LEAVES"] = nil
    value["LEFT"] = nil
    value["LESS"] = nil
    value["LEVEL"] = nil
    value["LIKE"] = nil
    value["LIMIT"] = nil
    value["LINEAR"] = nil
    value["LINES"] = nil
    value["LINESTRING"] = nil
    value["LIST"] = nil
    value["LOAD"] = nil
    value["LOCAL"] = nil
    value["LOCALTIME"] = nil
    value["LOCALTIMESTAMP"] = nil
    value["LOCK"] = nil
    value["LOCKED"] = nil
    value["LOCKED;"] = nil
    value["LOCKS"] = nil
    value["LOGFILE"] = nil
    value["LOGS"] = nil
    value["LONG"] = nil
    value["LONGBLOB"] = nil
    value["LONGTEXT"] = nil
    value["LOOP"] = nil
    value["LOW_PRIORITY"] = nil
    value["M"] = nil
    value["MASTER"] = nil
    value["MASTER_AUTO_POSITION"] = nil
    value["MASTER_BIND"] = nil
    value["MASTER_COMPRESSION_ALGORITHMS"] = nil
    value["MASTER_COMPRESSION_ALGORITHMS;"] = nil
    value["MASTER_CONNECT_RETRY"] = nil
    value["MASTER_DELAY"] = nil
    value["MASTER_HEARTBEAT_PERIOD"] = nil
    value["MASTER_HOST"] = nil
    value["MASTER_LOG_FILE"] = nil
    value["MASTER_LOG_POS"] = nil
    value["MASTER_PASSWORD"] = nil
    value["MASTER_PORT"] = nil
    value["MASTER_PUBLIC_KEY_PATH"] = nil
    value["MASTER_PUBLIC_KEY_PATH;"] = nil
    value["MASTER_RETRY_COUNT"] = nil
    value["MASTER_SERVER_ID"] = nil
    value["MASTER_SERVER_ID;"] = nil
    value["MASTER_SSL"] = nil
    value["MASTER_SSL_CA"] = nil
    value["MASTER_SSL_CAPATH"] = nil
    value["MASTER_SSL_CERT"] = nil
    value["MASTER_SSL_CIPHER"] = nil
    value["MASTER_SSL_CRL"] = nil
    value["MASTER_SSL_CRLPATH"] = nil
    value["MASTER_SSL_KEY"] = nil
    value["MASTER_SSL_VERIFY_SERVER_CERT"] = nil
    value["MASTER_TLS_CIPHERSUITES"] = nil
    value["MASTER_TLS_CIPHERSUITES;"] = nil
    value["MASTER_TLS_VERSION"] = nil
    value["MASTER_USER"] = nil
    value["MASTER_ZSTD_COMPRESSION_LEVEL"] = nil
    value["MASTER_ZSTD_COMPRESSION_LEVEL;"] = nil
    value["MATCH"] = nil
    value["MAXVALUE"] = nil
    value["MAX_CONNECTIONS_PER_HOUR"] = nil
    value["MAX_QUERIES_PER_HOUR"] = nil
    value["MAX_ROWS"] = nil
    value["MAX_SIZE"] = nil
    value["MAX_UPDATES_PER_HOUR"] = nil
    value["MAX_USER_CONNECTIONS"] = nil
    value["MEDIUM"] = nil
    value["MEDIUMBLOB"] = nil
    value["MEDIUMINT"] = nil
    value["MEDIUMTEXT"] = nil
    value["MEMBER"] = nil
    value["MEMBER;"] = nil
    value["MEMORY"] = nil
    value["MERGE"] = nil
    value["MESSAGE_TEXT"] = nil
    value["MICROSECOND"] = nil
    value["MIDDLEINT"] = nil
    value["MIGRATE"] = nil
    value["MINUTE"] = nil
    value["MINUTE_MICROSECOND"] = nil
    value["MINUTE_SECOND"] = nil
    value["MIN_ROWS"] = nil
    value["MOD"] = nil
    value["MODE"] = nil
    value["MODIFIES"] = nil
    value["MODIFY"] = nil
    value["MONTH"] = nil
    value["MULTILINESTRING"] = nil
    value["MULTIPOINT"] = nil
    value["MULTIPOLYGON"] = nil
    value["MUTEX"] = nil
    value["MYSQL_ERRNO"] = nil
    value["N"] = nil
    value["NAME"] = nil
    value["NAMES"] = nil
    value["NATIONAL"] = nil
    value["NATURAL"] = nil
    value["NCHAR"] = nil
    value["NDB"] = nil
    value["NDBCLUSTER"] = nil
    value["NESTED"] = nil
    value["NESTED;"] = nil
    value["NETWORK_NAMESPACE"] = nil
    value["NETWORK_NAMESPACE;"] = nil
    value["NEVER"] = nil
    value["NEW"] = nil
    value["NEXT"] = nil
    value["NO"] = nil
    value["NODEGROUP"] = nil
    value["NONE"] = nil
    value["NOT"] = nil
    value["NOWAIT"] = nil
    value["NOWAIT;"] = nil
    value["NO_WAIT"] = nil
    value["NO_WRITE_TO_BINLOG"] = nil
    value["NTH_VALUE"] = nil
    value["NTILE"] = nil
    value["NULL"] = nil
    value["NULLS"] = nil
    value["NULLS;"] = nil
    value["NUMBER"] = nil
    value["NUMERIC"] = nil
    value["NVARCHAR"] = nil
    value["O"] = nil
    value["OF"] = nil
    value["OFF"] = nil
    value["OFF;"] = nil
    value["OFFSET"] = nil
    value["OJ"] = nil
    value["OJ;"] = nil
    value["OLD"] = nil
    value["OLD;"] = nil
    value["ON"] = nil
    value["ONE"] = nil
    value["ONLY"] = nil
    value["OPEN"] = nil
    value["OPTIMIZE"] = nil
    value["OPTIMIZER_COSTS"] = nil
    value["OPTION"] = nil
    value["OPTIONAL"] = nil
    value["OPTIONAL;"] = nil
    value["OPTIONALLY"] = nil
    value["OPTIONS"] = nil
    value["OR"] = nil
    value["ORDER"] = nil
    value["ORDINALITY"] = nil
    value["ORDINALITY;"] = nil
    value["ORGANIZATION"] = nil
    value["ORGANIZATION;"] = nil
    value["OTHERS"] = nil
    value["OTHERS;"] = nil
    value["OUT"] = nil
    value["OUTER"] = nil
    value["OUTFILE"] = nil
    value["OVER"] = nil
    value["OWNER"] = nil
    value["P"] = nil
    value["PACK_KEYS"] = nil
    value["PAGE"] = nil
    value["PARSER"] = nil
    value["PARSE_GCOL_EXPR"] = nil
    value["PARTIAL"] = nil
    value["PARTITION"] = nil
    value["PARTITIONING"] = nil
    value["PARTITIONS"] = nil
    value["PASSWORD"] = nil
    value["PASSWORD_LOCK_TIME"] = nil
    value["PASSWORD_LOCK_TIME;"] = nil
    value["PATH"] = nil
    value["PATH;"] = nil
    value["PERCENT_RANK"] = nil
    value["PERSIST"] = nil
    value["PERSIST;"] = nil
    value["PERSIST_ONLY"] = nil
    value["PERSIST_ONLY;"] = nil
    value["PHASE"] = nil
    value["PLUGIN"] = nil
    value["PLUGINS"] = nil
    value["PLUGIN_DIR"] = nil
    value["POINT"] = nil
    value["POLYGON"] = nil
    value["PORT"] = nil
    value["PRECEDES"] = nil
    value["PRECEDING"] = nil
    value["PRECEDING;"] = nil
    value["PRECISION"] = nil
    value["PREPARE"] = nil
    value["PRESERVE"] = nil
    value["PREV"] = nil
    value["PRIMARY"] = nil
    value["PRIVILEGES"] = nil
    value["PRIVILEGE_CHECKS_USER"] = nil
    value["PRIVILEGE_CHECKS_USER;"] = nil
    value["PROCEDURE"] = nil
    value["PROCESS"] = nil
    value["PROCESS;"] = nil
    value["PROCESSLIST"] = nil
    value["PROFILE"] = nil
    value["PROFILES"] = nil
    value["PROXY"] = nil
    value["PURGE"] = nil
    value["Q"] = nil
    value["QUARTER"] = nil
    value["QUERY"] = nil
    value["QUICK"] = nil
    value["R"] = nil
    value["RANDOM"] = nil
    value["RANDOM;"] = nil
    value["RANGE"] = nil
    value["RANK"] = nil
    value["READ"] = nil
    value["READS"] = nil
    value["READ_ONLY"] = nil
    value["READ_WRITE"] = nil
    value["REAL"] = nil
    value["REBUILD"] = nil
    value["RECOVER"] = nil
    value["RECURSIVE"] = nil
    value["REDOFILE"] = nil
    value["REDOFILE;"] = nil
    value["REDO_BUFFER_SIZE"] = nil
    value["REDUNDANT"] = nil
    value["REFERENCE"] = nil
    value["REFERENCE;"] = nil
    value["REFERENCES"] = nil
    value["REGEXP"] = nil
    value["REGISTRATION"] = nil
    value["REGISTRATION;"] = nil
    value["RELAY"] = nil
    value["RELAYLOG"] = nil
    value["RELAY_LOG_FILE"] = nil
    value["RELAY_LOG_POS"] = nil
    value["RELAY_THREAD"] = nil
    value["RELEASE"] = nil
    value["RELOAD"] = nil
    value["REMOTE"] = nil
    value["REMOTE;"] = nil
    value["REMOVE"] = nil
    value["RENAME"] = nil
    value["REORGANIZE"] = nil
    value["REPAIR"] = nil
    value["REPEAT"] = nil
    value["REPEATABLE"] = nil
    value["REPLACE"] = nil
    value["REPLICA"] = nil
    value["REPLICA;"] = nil
    value["REPLICAS"] = nil
    value["REPLICAS;"] = nil
    value["REPLICATE_DO_DB"] = nil
    value["REPLICATE_DO_TABLE"] = nil
    value["REPLICATE_IGNORE_DB"] = nil
    value["REPLICATE_IGNORE_TABLE"] = nil
    value["REPLICATE_REWRITE_DB"] = nil
    value["REPLICATE_WILD_DO_TABLE"] = nil
    value["REPLICATE_WILD_IGNORE_TABLE"] = nil
    value["REPLICATION"] = nil
    value["REQUIRE"] = nil
    value["REQUIRE_ROW_FORMAT"] = nil
    value["REQUIRE_ROW_FORMAT;"] = nil
    value["RESET"] = nil
    value["RESIGNAL"] = nil
    value["RESOURCE"] = nil
    value["RESOURCE;"] = nil
    value["RESPECT"] = nil
    value["RESPECT;"] = nil
    value["RESTART"] = nil
    value["RESTART;"] = nil
    value["RESTORE"] = nil
    value["RESTRICT"] = nil
    value["RESUME"] = nil
    value["RETAIN"] = nil
    value["RETAIN;"] = nil
    value["RETURN"] = nil
    value["RETURNED_SQLSTATE"] = nil
    value["RETURNING"] = nil
    value["RETURNING;"] = nil
    value["RETURNS"] = nil
    value["REUSE"] = nil
    value["REUSE;"] = nil
    value["REVERSE"] = nil
    value["REVOKE"] = nil
    value["RIGHT"] = nil
    value["RLIKE"] = nil
    value["ROLE"] = nil
    value["ROLE;"] = nil
    value["ROLLBACK"] = nil
    value["ROLLUP"] = nil
    value["ROTATE"] = nil
    value["ROUTINE"] = nil
    value["ROW"] = nil
    value["ROWS"] = nil
    value["ROW_COUNT"] = nil
    value["ROW_FORMAT"] = nil
    value["ROW_NUMBER"] = nil
    value["RTREE"] = nil
    value["S"] = nil
    value["SAVEPOINT"] = nil
    value["SCHEDULE"] = nil
    value["SCHEMA"] = nil
    value["SCHEMAS"] = nil
    value["SCHEMA_NAME"] = nil
    value["SECOND"] = nil
    value["SECONDARY"] = nil
    value["SECONDARY;"] = nil
    value["SECONDARY_ENGINE"] = nil
    value["SECONDARY_ENGINE;"] = nil
    value["SECONDARY_ENGINE_ATTRIBUTE"] = nil
    value["SECONDARY_ENGINE_ATTRIBUTE;"] = nil
    value["SECONDARY_LOAD"] = nil
    value["SECONDARY_LOAD;"] = nil
    value["SECONDARY_UNLOAD"] = nil
    value["SECONDARY_UNLOAD;"] = nil
    value["SECOND_MICROSECOND"] = nil
    value["SECURITY"] = nil
    value["SELECT"] = nil
    value["SENSITIVE"] = nil
    value["SEPARATOR"] = nil
    value["SERIAL"] = nil
    value["SERIALIZABLE"] = nil
    value["SERVER"] = nil
    value["SESSION"] = nil
    value["SET"] = nil
    value["SHARE"] = nil
    value["SHOW"] = nil
    value["SHUTDOWN"] = nil
    value["SIGNAL"] = nil
    value["SIGNED"] = nil
    value["SIMPLE"] = nil
    value["SKIP"] = nil
    value["SKIP;"] = nil
    value["SLAVE"] = nil
    value["SLOW"] = nil
    value["SMALLINT"] = nil
    value["SNAPSHOT"] = nil
    value["SOCKET"] = nil
    value["SOME"] = nil
    value["SONAME"] = nil
    value["SOUNDS"] = nil
    value["SOURCE"] = nil
    value["SOURCE_AUTO_POSITION"] = nil
    value["SOURCE_AUTO_POSITION;"] = nil
    value["SOURCE_BIND"] = nil
    value["SOURCE_BIND;"] = nil
    value["SOURCE_COMPRESSION_ALGORITHMS"] = nil
    value["SOURCE_COMPRESSION_ALGORITHMS;"] = nil
    value["SOURCE_CONNECT_RETRY"] = nil
    value["SOURCE_CONNECT_RETRY;"] = nil
    value["SOURCE_DELAY"] = nil
    value["SOURCE_DELAY;"] = nil
    value["SOURCE_HEARTBEAT_PERIOD"] = nil
    value["SOURCE_HEARTBEAT_PERIOD;"] = nil
    value["SOURCE_HOST"] = nil
    value["SOURCE_HOST;"] = nil
    value["SOURCE_LOG_FILE"] = nil
    value["SOURCE_LOG_FILE;"] = nil
    value["SOURCE_LOG_POS"] = nil
    value["SOURCE_LOG_POS;"] = nil
    value["SOURCE_PASSWORD"] = nil
    value["SOURCE_PASSWORD;"] = nil
    value["SOURCE_PORT"] = nil
    value["SOURCE_PORT;"] = nil
    value["SOURCE_PUBLIC_KEY_PATH"] = nil
    value["SOURCE_PUBLIC_KEY_PATH;"] = nil
    value["SOURCE_RETRY_COUNT"] = nil
    value["SOURCE_RETRY_COUNT;"] = nil
    value["SOURCE_SSL"] = nil
    value["SOURCE_SSL;"] = nil
    value["SOURCE_SSL_CA"] = nil
    value["SOURCE_SSL_CA;"] = nil
    value["SOURCE_SSL_CAPATH"] = nil
    value["SOURCE_SSL_CAPATH;"] = nil
    value["SOURCE_SSL_CERT"] = nil
    value["SOURCE_SSL_CERT;"] = nil
    value["SOURCE_SSL_CIPHER"] = nil
    value["SOURCE_SSL_CIPHER;"] = nil
    value["SOURCE_SSL_CRL"] = nil
    value["SOURCE_SSL_CRL;"] = nil
    value["SOURCE_SSL_CRLPATH"] = nil
    value["SOURCE_SSL_CRLPATH;"] = nil
    value["SOURCE_SSL_KEY"] = nil
    value["SOURCE_SSL_KEY;"] = nil
    value["SOURCE_SSL_VERIFY_SERVER_CERT"] = nil
    value["SOURCE_SSL_VERIFY_SERVER_CERT;"] = nil
    value["SOURCE_TLS_CIPHERSUITES"] = nil
    value["SOURCE_TLS_CIPHERSUITES;"] = nil
    value["SOURCE_TLS_VERSION"] = nil
    value["SOURCE_TLS_VERSION;"] = nil
    value["SOURCE_USER"] = nil
    value["SOURCE_USER;"] = nil
    value["SOURCE_ZSTD_COMPRESSION_LEVEL"] = nil
    value["SOURCE_ZSTD_COMPRESSION_LEVEL;"] = nil
    value["SPATIAL"] = nil
    value["SPECIFIC"] = nil
    value["SQL"] = nil
    value["SQLEXCEPTION"] = nil
    value["SQLSTATE"] = nil
    value["SQLWARNING"] = nil
    value["SQL_AFTER_GTIDS"] = nil
    value["SQL_AFTER_MTS_GAPS"] = nil
    value["SQL_BEFORE_GTIDS"] = nil
    value["SQL_BIG_RESULT"] = nil
    value["SQL_BUFFER_RESULT"] = nil
    value["SQL_CACHE"] = nil
    value["SQL_CACHE;"] = nil
    value["SQL_CALC_FOUND_ROWS"] = nil
    value["SQL_NO_CACHE"] = nil
    value["SQL_SMALL_RESULT"] = nil
    value["SQL_THREAD"] = nil
    value["SQL_TSI_DAY"] = nil
    value["SQL_TSI_HOUR"] = nil
    value["SQL_TSI_MINUTE"] = nil
    value["SQL_TSI_MONTH"] = nil
    value["SQL_TSI_QUARTER"] = nil
    value["SQL_TSI_SECOND"] = nil
    value["SQL_TSI_WEEK"] = nil
    value["SQL_TSI_YEAR"] = nil
    value["SRID"] = nil
    value["SRID;"] = nil
    value["SSL"] = nil
    value["STACKED"] = nil
    value["START"] = nil
    value["STARTING"] = nil
    value["STARTS"] = nil
    value["STATS_AUTO_RECALC"] = nil
    value["STATS_PERSISTENT"] = nil
    value["STATS_SAMPLE_PAGES"] = nil
    value["STATUS"] = nil
    value["STOP"] = nil
    value["STORAGE"] = nil
    value["STORED"] = nil
    value["STRAIGHT_JOIN"] = nil
    value["STREAM"] = nil
    value["STREAM;"] = nil
    value["STRING"] = nil
    value["SUBCLASS_ORIGIN"] = nil
    value["SUBJECT"] = nil
    value["SUBPARTITION"] = nil
    value["SUBPARTITIONS"] = nil
    value["SUPER"] = nil
    value["SUSPEND"] = nil
    value["SWAPS"] = nil
    value["SWITCHES"] = nil
    value["SYSTEM"] = nil
    value["T"] = nil
    value["TABLE"] = nil
    value["TABLES"] = nil
    value["TABLESPACE"] = nil
    value["TABLE_CHECKSUM"] = nil
    value["TABLE_NAME"] = nil
    value["TEMPORARY"] = nil
    value["TEMPTABLE"] = nil
    value["TERMINATED"] = nil
    value["TEXT"] = nil
    value["THAN"] = nil
    value["THEN"] = nil
    value["THREAD_PRIORITY"] = nil
    value["THREAD_PRIORITY;"] = nil
    value["TIES"] = nil
    value["TIES;"] = nil
    value["TIME"] = nil
    value["TIMESTAMP"] = nil
    value["TIMESTAMPADD"] = nil
    value["TIMESTAMPDIFF"] = nil
    value["TINYBLOB"] = nil
    value["TINYINT"] = nil
    value["TINYTEXT"] = nil
    value["TLS"] = nil
    value["TLS;"] = nil
    value["TO"] = nil
    value["TRAILING"] = nil
    value["TRANSACTION"] = nil
    value["TRIGGER"] = nil
    value["TRIGGERS"] = nil
    value["TRUE"] = nil
    value["TRUNCATE"] = nil
    value["TYPE"] = nil
    value["TYPES"] = nil
    value["U"] = nil
    value["UNBOUNDED"] = nil
    value["UNBOUNDED;"] = nil
    value["UNCOMMITTED"] = nil
    value["UNDEFINED"] = nil
    value["UNDO"] = nil
    value["UNDOFILE"] = nil
    value["UNDO_BUFFER_SIZE"] = nil
    value["UNICODE"] = nil
    value["UNINSTALL"] = nil
    value["UNION"] = nil
    value["UNIQUE"] = nil
    value["UNKNOWN"] = nil
    value["UNLOCK"] = nil
    value["UNREGISTER"] = nil
    value["UNREGISTER;"] = nil
    value["UNSIGNED"] = nil
    value["UNTIL"] = nil
    value["UPDATE"] = nil
    value["UPGRADE"] = nil
    value["USAGE"] = nil
    value["USE"] = nil
    value["USER"] = nil
    value["USER_RESOURCES"] = nil
    value["USE_FRM"] = nil
    value["USING"] = nil
    value["UTC_DATE"] = nil
    value["UTC_TIME"] = nil
    value["UTC_TIMESTAMP"] = nil
    value["V"] = nil
    value["VALIDATION"] = nil
    value["VALUE"] = nil
    value["VALUES"] = nil
    value["VARBINARY"] = nil
    value["VARCHAR"] = nil
    value["VARCHARACTER"] = nil
    value["VARIABLES"] = nil
    value["VARYING"] = nil
    value["VCPU"] = nil
    value["VCPU;"] = nil
    value["VIEW"] = nil
    value["VIRTUAL"] = nil
    value["VISIBLE"] = nil
    value["W"] = nil
    value["WAIT"] = nil
    value["WARNINGS"] = nil
    value["WEEK"] = nil
    value["WEIGHT_STRING"] = nil
    value["WHEN"] = nil
    value["WHERE"] = nil
    value["WHILE"] = nil
    value["WINDOW"] = nil
    value["WITH"] = nil
    value["WITHOUT"] = nil
    value["WORK"] = nil
    value["WRAPPER"] = nil
    value["WRITE"] = nil
    value["X"] = nil
    value["X509"] = nil
    value["XA"] = nil
    value["XID"] = nil
    value["XML"] = nil
    value["XOR"] = nil
    value["Y"] = nil
    value["YEAR"] = nil
    value["YEAR_MONTH"] = nil
    value["Z"] = nil
    value["ZEROFILL"] = nil
    value["ZONE"] = nil
    value["ZONE;"] = nil
    value["_FILENAME"] = nil
return value 
}