SET FOREIGN_KEY_CHECKS = 0

SELECT concat('DROP TABLE IF EXISTS `', table_name, '`;')
FROM information_schema.tables
WHERE table_schema = 'MyDatabaseName';

SET FOREIGN_KEY_CHECKS = 1
