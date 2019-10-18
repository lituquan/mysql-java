CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `pref_list` text NOT NULL,
  `latest_log_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;