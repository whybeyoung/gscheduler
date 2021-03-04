CREATE TABLE `t_gs_process_definition` (
                                           `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'key',
                                           `name` varchar(255) DEFAULT NULL COMMENT 'process definition name',
                                           `version` int(11) DEFAULT NULL COMMENT 'process definition version',
                                           `release_state` tinyint(4) DEFAULT NULL COMMENT 'process definition release state：0:offline,1:online',
                                           `group_id` int(11) DEFAULT NULL COMMENT 'group_id分组id',
                                           `user_id` int(11) DEFAULT NULL COMMENT 'process definition creator id',
                                           `process_definition_json` longtext COMMENT 'process definition json content',
                                           `description` text,
                                           `global_params` text COMMENT 'global parameters',
                                           `flag` tinyint(4) DEFAULT NULL COMMENT '0 not available, 1 available',
                                           `locations` text COMMENT 'Node location information',
                                           `connects` text COMMENT 'Node connection information',
                                           `receivers` text COMMENT 'receivers',
                                           `receivers_cc` text COMMENT 'cc',
                                           `create_time` datetime DEFAULT NULL COMMENT 'create time',
                                           `timeout` int(11) DEFAULT '0' COMMENT 'time out',
                                           `tenant_id` int(11) NOT NULL DEFAULT '-1' COMMENT 'tenant id',
                                           `update_time` datetime DEFAULT NULL COMMENT 'update time',
                                           `modify_by` varchar(36) DEFAULT '' COMMENT 'modify user',
                                           `resource_ids` varchar(255) DEFAULT NULL COMMENT 'resource ids',
                                           PRIMARY KEY (`id`),
                                           UNIQUE KEY `process_definition_unique` (`name`,`group_id`),
                                           KEY `process_definition_index` (`group_id`,`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;



CREATE TABLE `t_gs_process_definition_version` (
                                                   `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'key',
                                                   `process_definition_id` int(11) NOT NULL COMMENT 'process definition id',
                                                   `version` int(11) DEFAULT NULL COMMENT 'process definition version',
                                                   `process_definition_json` longtext COMMENT 'process definition json content',
                                                   `description` text,
                                                   `global_params` text COMMENT 'global parameters',
                                                   `locations` text COMMENT 'Node location information',
                                                   `connects` text COMMENT 'Node connection information',
                                                   `receivers` text COMMENT 'receivers',
                                                   `receivers_cc` text COMMENT 'cc',
                                                   `create_time` datetime DEFAULT NULL COMMENT 'create time',
                                                   `timeout` int(11) DEFAULT '0' COMMENT 'time out',
                                                   `resource_ids` varchar(255) DEFAULT NULL COMMENT 'resource ids',
                                                   PRIMARY KEY (`id`),
                                                   UNIQUE KEY `process_definition_id_and_version` (`process_definition_id`,`version`) USING BTREE,
                                                   KEY `process_definition_index` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;



CREATE TABLE `t_gs_process_instance` (
                                         `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'key',
                                         `name` varchar(255) DEFAULT NULL COMMENT 'process instance name',
                                         `process_definition_id` int(11) DEFAULT NULL COMMENT 'process definition id',
                                         `state` tinyint(4) DEFAULT NULL COMMENT 'process instance Status: 0 commit succeeded, 1 running, 2 prepare to pause, 3 pause, 4 prepare to stop, 5 stop, 6 fail, 7 succeed, 8 need fault tolerance, 9 kill, 10 wait for thread, 11 wait for dependency to complete',
                                         `recovery` tinyint(4) DEFAULT NULL COMMENT 'process instance failover flag：0:normal,1:failover instance',
                                         `start_time` datetime DEFAULT NULL COMMENT 'process instance start time',
                                         `end_time` datetime DEFAULT NULL COMMENT 'process instance end time',
                                         `run_times` int(11) DEFAULT NULL COMMENT 'process instance run times',
                                         `host` varchar(45) DEFAULT NULL COMMENT 'process instance host',
                                         `command_type` tinyint(4) DEFAULT NULL COMMENT 'command type',
                                         `command_param` text COMMENT 'json command parameters',
                                         `task_depend_type` tinyint(4) DEFAULT NULL COMMENT 'task depend type. 0: only current node,1:before the node,2:later nodes',
                                         `max_try_times` tinyint(4) DEFAULT '0' COMMENT 'max try times',
                                         `failure_strategy` tinyint(4) DEFAULT '0' COMMENT 'failure strategy. 0:end the process when node failed,1:continue running the other nodes when node failed',
                                          `warning_type` tinyint(4) DEFAULT '0' COMMENT 'warning type. 0:no warning,1:warning if process success,2:warning if process failed,3:warning if success',
                                         `warning_group_id` int(11) DEFAULT NULL COMMENT 'warning group id',
                                         `schedule_time` datetime DEFAULT NULL COMMENT 'schedule time',
                                         `command_start_time` datetime DEFAULT NULL COMMENT 'command start time',
                                         `global_params` text COMMENT 'global parameters',
                                         `process_instance_json` longtext COMMENT 'process instance json(copy的process definition 的json)',
                                         `flag` tinyint(4) DEFAULT '1' COMMENT 'flag',
                                         `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                         `is_sub_process` int(11) DEFAULT '0' COMMENT 'flag, whether the process is sub process',
                                         `executor_id` int(11) NOT NULL COMMENT 'executor id',
                                         `locations` text COMMENT 'Node location information',
                                         `connects` text COMMENT 'Node connection information',
                                         `history_cmd` text COMMENT 'history commands of process instance operation',
                                         `dependence_schedule_times` text COMMENT 'depend schedule fire time',
                                         `process_instance_priority` int(11) DEFAULT NULL COMMENT 'process instance priority. 0 Highest,1 High,2 Medium,3 Low,4 Lowest',
                                         `worker_group` varchar(64) DEFAULT '' COMMENT 'worker group',
                                         `timeout` int(11) DEFAULT '0' COMMENT 'time out',
                                         `tenant_id` int(11) NOT NULL DEFAULT '-1' COMMENT 'tenant id',
                                         `var_pool` longtext,
                                         PRIMARY KEY (`id`),
                                         KEY `process_instance_index` (`process_definition_id`,`id`) USING BTREE,
                                         KEY `start_time_index` (`start_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;



CREATE TABLE `t_gs_group` (
                                `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'key',
                                `name` varchar(100) DEFAULT NULL COMMENT 'group name',
                                `description` varchar(200) DEFAULT NULL,
                                `user_id` int(11) DEFAULT NULL COMMENT 'creator id',
                                `flag` tinyint(4) DEFAULT '1' COMMENT '0 not available, 1 available',
                                `create_time` datetime DEFAULT NULL COMMENT 'create time',
                                `update_time` datetime DEFAULT NULL COMMENT 'update time',
                                PRIMARY KEY (`id`),
                                KEY `user_id_index` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;


