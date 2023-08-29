/*
 Navicat Premium Data Transfer

 Source Server         : payment-dev
 Source Server Type    : MySQL
 Source Server Version : 50735 (5.7.35-log)
 Source Host           : rm-8vbjy34g96075qpoklo.mysql.zhangbei.rds.aliyuncs.com:3408
 Source Schema         : car_broke

 Target Server Type    : MySQL
 Target Server Version : 50735 (5.7.35-log)
 File Encoding         : 65001

 Date: 29/08/2023 10:41:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for alarm
-- ----------------------------
DROP TABLE IF EXISTS `alarm`;
CREATE TABLE `alarm`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `faultChargeableDeviceNum` int(11) NULL DEFAULT NULL COMMENT '可充电储能装置故障总数',
  `faultChargeableDeviceList` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '可充电储能装置故障代码列表',
  `faultDriveMotorNum` int(11) NULL DEFAULT NULL COMMENT '驱动电机故障总数',
  `faultDriveMotorList` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '驱动电机故障代码列表',
  `faultEngineNum` int(11) NULL DEFAULT NULL COMMENT '发动机故障总数',
  `faultEngineList` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '发动机故障列表',
  `faultOthersNum` int(11) NULL DEFAULT NULL COMMENT '其他故障总数',
  `faultOthersList` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '其他故障代码列表',
  `generalAlarmFlag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '通用报警标志',
  `maxAlarmLevel` int(11) NULL DEFAULT NULL COMMENT '最高报警等级',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '报警数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for chargeabletemp
-- ----------------------------
DROP TABLE IF EXISTS `chargeabletemp`;
CREATE TABLE `chargeabletemp`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `number` int(11) NULL DEFAULT NULL,
  `data` json NULL COMMENT '可充电储能子系统电压信息列表',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '可充电储能装置温度数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for chargeablevoltage
-- ----------------------------
DROP TABLE IF EXISTS `chargeablevoltage`;
CREATE TABLE `chargeablevoltage`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `number` int(11) NULL DEFAULT NULL COMMENT '可充电储能子系统个数',
  `data` json NULL COMMENT '可充电储能子系统电压信息列表',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '可充电储能装置电压数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for command
-- ----------------------------
DROP TABLE IF EXISTS `command`;
CREATE TABLE `command`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` json NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for drivemotor
-- ----------------------------
DROP TABLE IF EXISTS `drivemotor`;
CREATE TABLE `drivemotor`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `number` int(11) NULL DEFAULT NULL COMMENT '驱动电机个数',
  `data` json NULL COMMENT '驱动电机数据',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '驱动电机数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for engine
-- ----------------------------
DROP TABLE IF EXISTS `engine`;
CREATE TABLE `engine`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `crankshaftSpeed` int(11) NULL DEFAULT NULL COMMENT '曲轴转速；有效范围：0-60 000（表示0 r/min-60 000 r/min)，最小计量单元：1 r/min,0xFF/0xFE表示异常，0xFF,0xFF表示无效',
  `fuelConsumption` int(11) NULL DEFAULT NULL COMMENT '燃料消耗率；有效范围:0~60 000(表示0r/min一60 000 r/min)，最小计量单元:1r/min，0xFF/OxFE表示异常，0xFF表示无效',
  `status` int(11) NULL DEFAULT NULL COMMENT '发动机状态；0x01启动状态；0x02关闭状态，0xFE表示异常，0xFF表示无效',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '发动机数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for extreme
-- ----------------------------
DROP TABLE IF EXISTS `extreme`;
CREATE TABLE `extreme`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `maxBatteryVoltage` int(11) NULL DEFAULT NULL COMMENT '电池单体电压最高值',
  `maxTemp` int(11) NULL DEFAULT NULL COMMENT '最高温度值',
  `maxTempProbeNo` int(11) NULL DEFAULT NULL COMMENT '最高温度探针序号',
  `maxTempSubsysNo` int(11) NULL DEFAULT NULL COMMENT '最高温度子系统号',
  `maxVoltageBatteryCode` int(11) NULL DEFAULT NULL COMMENT '最高电压电池单体代号',
  `maxVoltageBatterySubsysNo` int(11) NULL DEFAULT NULL COMMENT '最高电压电池子系统号',
  `minBatteryVoltage` int(11) NULL DEFAULT NULL COMMENT '电池单体电压最低值',
  `minTemp` int(11) NULL DEFAULT NULL COMMENT '最低温度值',
  `minTempProbeNo` int(11) NULL DEFAULT NULL COMMENT '最低温度探针序号',
  `minTempSubsysNo` int(11) NULL DEFAULT NULL COMMENT '最低温度子系统号',
  `minVoltageBatteryCode` int(11) NULL DEFAULT NULL COMMENT '最低电压电池单体代号',
  `minVoltageBatterySubsysNo` int(11) NULL DEFAULT NULL COMMENT '最低电压电池子系统号',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '极值数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for fuelcell
-- ----------------------------
DROP TABLE IF EXISTS `fuelcell`;
CREATE TABLE `fuelcell`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `cellCurrent` int(11) NULL DEFAULT NULL COMMENT '燃料电池电流',
  `cellVoltage` int(11) NULL DEFAULT NULL COMMENT '燃料电池电压',
  `dcStatus` int(11) NULL DEFAULT NULL COMMENT '高压DC/DC状态',
  `fuelConsumption` int(11) NULL DEFAULT NULL COMMENT '燃料消耗率',
  `h_MaxConc` int(11) NULL DEFAULT NULL COMMENT '氢气最高浓度',
  `h_ConcSensorCode` int(11) NULL DEFAULT NULL COMMENT '氢气最高浓度传感器代号',
  `h_MaxPress` int(11) NULL DEFAULT NULL COMMENT '氢气最高压力',
  `h_PressSensorCode` int(11) NULL DEFAULT NULL COMMENT '氢气最高压力传感器代号',
  `h_MaxTemp` int(11) NULL DEFAULT NULL COMMENT '氢系统中最高温度',
  `h_TempProbeCode` int(11) NULL DEFAULT NULL COMMENT '氢系统中最高温度探针代号',
  `probeNum` int(11) NULL DEFAULT NULL COMMENT '燃料电池温度探针总数',
  `probeTemps` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '探针温度值',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '燃料电池数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for location
-- ----------------------------
DROP TABLE IF EXISTS `location`;
CREATE TABLE `location`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `latitude` int(11) NULL DEFAULT NULL COMMENT '纬度；以度为单位的纬度值乘以10^6，精确到百万分之一度',
  `longitude` int(11) NULL DEFAULT NULL COMMENT '经度；以度为单位的纬度值乘以10^6，精确到百万分之一度',
  `status` int(11) NULL DEFAULT NULL,
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '车辆位置数据' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for msg_log
-- ----------------------------
DROP TABLE IF EXISTS `msg_log`;
CREATE TABLE `msg_log`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `head` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cmdCFlag` int(3) NULL DEFAULT NULL,
  `cmdRsp` int(3) NULL DEFAULT NULL,
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `encrypt` int(2) NULL DEFAULT NULL,
  `len` int(5) NULL DEFAULT NULL,
  `data` blob NULL,
  `code` int(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 90 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '消息log' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for vehicle
-- ----------------------------
DROP TABLE IF EXISTS `vehicle`;
CREATE TABLE `vehicle`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `acceleratorPedal` int(11) NULL DEFAULT NULL,
  `brakePedal` int(11) NULL DEFAULT NULL,
  `charging` int(11) NULL DEFAULT NULL COMMENT '充电；0x01;停车充电;0x02;行驶充电;0x03;未充电状态;x04:充电完成;“0xFE”表示异常，“0xFE”表示无效',
  `current` int(11) NULL DEFAULT NULL COMMENT '总电流；有效值范围:0~20 000(偏移量1 000 A，表示一1000 A~+1 000 A)最小计量单元01A，“0xFFOXFE”表示异常“0xFF0xFF”表示无效',
  `dc` int(11) NULL DEFAULT NULL COMMENT '0x01：工作0x02：断开，\"OxFE”表示异常，\"OxFF”表示无效',
  `gear` int(11) NULL DEFAULT NULL COMMENT '挡位；',
  `mileage` int(11) NULL DEFAULT NULL COMMENT '累计里程；有效值范围:0~9 999 999(表示0 km~999 999.9 km)最小计量单元;0.1 km\n“0xFF，0xFF，0xFF,OxFE”表示异常，“0xFF,OxFF,OxFF0xFF”示无效',
  `mode` int(11) NULL DEFAULT NULL COMMENT '运行模式；0x01:纯电;0x02;混动;0x03:燃油;0xFE表示异常;0xFF 表示无效',
  `resistance` int(11) NULL DEFAULT NULL COMMENT '绝缘电阻；有效范围0~60 000(表示 0 Q~60 000 k2)最小计量单元;1 k2',
  `soc` int(11) NULL DEFAULT NULL COMMENT '有效值范围;0~100(表示 0%~100%),最小计量单元:1%“0xFE”表示异常“0xFF”表示无效',
  `speed` int(11) NULL DEFAULT NULL COMMENT '车速；有效值范围;0~100(表示 0%~100%),最小计量单元:1%，“0xFE”表示异常“0xFF”表示无效',
  `status` int(11) NULL DEFAULT NULL COMMENT '车辆状态；0x01:启动，0x02:熄火，0x03:其他，0xFE表示异常，0xFF表示无效',
  `voltage` int(11) NULL DEFAULT NULL COMMENT '总电压；有效值范围:0~10 000(表示 0 V~1000 V)，小计量单元:0.1V，“0xFF,xFE”示异常，“0xFFxFF”表示无效',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '整车信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for vehicle_conn
-- ----------------------------
DROP TABLE IF EXISTS `vehicle_conn`;
CREATE TABLE `vehicle_conn`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `host` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `port` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '8080',
  `connStr` int(32) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL COMMENT '1在线',
  `onlineDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `offlineDate` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `i_vin`(`vin`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9911 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '车辆链接注册表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for vlogin
-- ----------------------------
DROP TABLE IF EXISTS `vlogin`;
CREATE TABLE `vlogin`  (
  `vin` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `seq` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '车辆登录的流水号',
  `iccId` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'iccid',
  `num` int(11) NULL DEFAULT NULL COMMENT '可充电储能子系统数',
  `length` int(11) NULL DEFAULT NULL COMMENT '可充电储能系统编码长度。',
  `energyId` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '可充电储能系统编码列表。',
  `created_time` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '车辆登入' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for vlogout
-- ----------------------------
DROP TABLE IF EXISTS `vlogout`;
CREATE TABLE `vlogout`  (
  `vin` varchar(17) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int(11) NULL DEFAULT NULL,
  `month` int(11) NULL DEFAULT NULL,
  `day` int(11) NULL DEFAULT NULL,
  `hour` int(11) NULL DEFAULT NULL,
  `minutes` int(11) NULL DEFAULT NULL,
  `seconds` int(11) NULL DEFAULT NULL,
  `seq` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '车辆登录的流水号',
  `createdTime` datetime NULL DEFAULT NULL COMMENT '上报时间',
  PRIMARY KEY (`vin`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '车辆登出' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
