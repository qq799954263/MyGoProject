/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50722
Source Host           : localhost:3306
Source Database       : student

Target Server Type    : MYSQL
Target Server Version : 50722
File Encoding         : 65001

Date: 2018-08-18 17:45:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `studentlist`
-- ----------------------------
DROP TABLE IF EXISTS `studentlist`;
CREATE TABLE `studentlist` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `username` char(16) DEFAULT NULL,
  `class` char(16) DEFAULT NULL,
  `sex` char(16) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of studentlist
-- ----------------------------
INSERT INTO `studentlist` VALUES ('1', '小红', '2班', '女');
INSERT INTO `studentlist` VALUES ('2', '小明', '1班', '男');
INSERT INTO `studentlist` VALUES ('3', '张三', '3班', '男');
INSERT INTO `studentlist` VALUES ('5', '赵四', '2班', '男');
