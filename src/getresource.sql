-- --------------------------------------------------
--  Table Structure for `model/collect.Pods`
-- --------------------------------------------------
DROP TABLE IF EXISTS `list_pods`;
DROP TABLE IF EXISTS `list_nodes`;
DROP TABLE IF EXISTS `list_services`;
DROP TABLE IF EXISTS `dashboard_service_second`;
DROP TABLE IF EXISTS `dashboard_service_minute`;
DROP TABLE IF EXISTS `dashboard_service_hour`;
DROP TABLE IF EXISTS `dashboard_service_day`;
DROP TABLE IF EXISTS `dashboard_service_week`;
CREATE TABLE `list_pods` (
  `id`                    BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `pod_name`              VARCHAR(30)           NOT NULL DEFAULT '',
  `pod_hostIP`            VARCHAR(30)           NOT NULL DEFAULT '',
  `containers_numbers`    VARCHAR(30)           NOT NULL DEFAULT '',
  `Creat_time`            TIMESTAMP             NOT NULL DEFAULT NOW(),
  `Record_time`           TIMESTAMP             NOT NULL DEFAULT NOW(),
  `All_pod_numbers`       VARCHAR(30)           NOT NULL DEFAULT '',
  `All_container_numbers` VARCHAR(30)           NOT NULL DEFAULT ''
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;
;

-- --------------------------------------------------
--  Table Structure for `model/collect.Nodes`
-- --------------------------------------------------
CREATE TABLE `list_nodes` (
  `id`               BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `Node_name`        VARCHAR(30)           NOT NULL DEFAULT '',
  `Numbers_cpu_core` VARCHAR(30)           NOT NULL DEFAULT '',
  `Numbers_gpu_core` VARCHAR(30)           NOT NULL DEFAULT '',
  `Memory_size`      VARCHAR(30)           NOT NULL DEFAULT '',
  `Pod_limit`        VARCHAR(30)           NOT NULL DEFAULT '',
  `Create_time`      TIMESTAMP             NOT NULL DEFAULT NOW(),
  `Record_time`      TIMESTAMP             NOT NULL DEFAULT NOW()
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------
--  Table Structure for `model/collect.Services`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `list_services` (
  `uid`             BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `service_name`    VARCHAR(30)           NOT NULL DEFAULT '',
  `service_numbers` VARCHAR(30)           NOT NULL DEFAULT '',
  `creat_time`      TIMESTAMP             NOT NULL DEFAULT NOW(),
  `record_time`     TIMESTAMP             NOT NULL DEFAULT NOW()
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;
-- --------------------------------------------------
--  Table Structure for `model/dashboard`
-- --------------------------------------------------

CREATE TABLE `dashboard_service_second` (
  `id`               BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `Service_numbers`  VARCHAR(30)           NOT NULL DEFAULT '',
  `pod_number`       VARCHAR(30)           NOT NULL DEFAULT '',
  `container_number` VARCHAR(30)           NOT NULL DEFAULT '',
  `Record_time`      TIMESTAMP             NOT NULL DEFAULT NOW()
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------
--  Table Structure for `model/dashboard`
-- --------------------------------------------------

CREATE TABLE `dashboard_service_minute` (
  `id`               BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `Service_numbers`  VARCHAR(30)           NOT NULL DEFAULT '',
  `pod_number`       VARCHAR(30)           NOT NULL DEFAULT '',
  `container_number` VARCHAR(30)           NOT NULL DEFAULT '',
  `Record_time`      TIMESTAMP             NOT NULL DEFAULT NOW()
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------
--  Table Structure for `model/dashboard`
-- --------------------------------------------------

CREATE TABLE `dashboard_service_hour` (
  `id`               BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `Service_numbers`  VARCHAR(30)           NOT NULL DEFAULT '',
  `pod_number`       VARCHAR(30)           NOT NULL DEFAULT '',
  `container_number` VARCHAR(30)           NOT NULL DEFAULT '',
  `Record_time`      TIMESTAMP             NOT NULL DEFAULT NOW()
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------
--  Table Structure for `model/dashboard`
-- --------------------------------------------------

CREATE TABLE `dashboard_service_day` (
  `id`               BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `Service_numbers`  VARCHAR(30)           NOT NULL DEFAULT '',
  `pod_number`       VARCHAR(30)           NOT NULL DEFAULT '',
  `container_number` VARCHAR(30)           NOT NULL DEFAULT '',
  `Record_time`      TIMESTAMP             NOT NULL DEFAULT NOW()
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;
-- --------------------------------------------------
--  Table Structure for `model/dashboard`
-- --------------------------------------------------

CREATE TABLE `dashboard_service_week` (
  `id`               BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `Service_numbers`  VARCHAR(30)           NOT NULL DEFAULT '',
  `pod_number`       VARCHAR(30)           NOT NULL DEFAULT '',
  `container_number` VARCHAR(30)           NOT NULL DEFAULT '',
  `Record_time`      TIMESTAMP             NOT NULL DEFAULT NOW()
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;
