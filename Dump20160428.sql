CREATE DATABASE  IF NOT EXISTS `xin_cg` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `xin_cg`;
-- MySQL dump 10.13  Distrib 5.6.13, for Win32 (x86)
--
-- Host: 127.0.0.1    Database: xin_cg
-- ------------------------------------------------------
-- Server version	5.6.17-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin` (
  `id` mediumint(6) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(32) DEFAULT NULL COMMENT '密码',
  `roleid` smallint(5) DEFAULT '0' COMMENT '角色',
  `lastloginip` varchar(15) DEFAULT '0.0.0.0' COMMENT '最后登陆地址PI',
  `lastlogintime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登陆时间',
  `email` varchar(40) DEFAULT NULL COMMENT '邮箱',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `lang` varchar(6) NOT NULL DEFAULT 'zh-cn' COMMENT '语言',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1:允许登陆 0:禁止登陆 ',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `code` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `username` (`username`,`roleid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','e10adc3949ba59abbe56e057f20f883e',1,'','2016-04-19 09:34:34','abc@163.com','admin','zh-cn',1,'2014-01-17 23:58:58','01');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_panel`
--

DROP TABLE IF EXISTS `admin_panel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_panel` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
  `aid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `name` varchar(40) DEFAULT '' COMMENT '菜单名称',
  `url` varchar(255) DEFAULT '' COMMENT '菜单url',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '添加时间',
  UNIQUE KEY `uid` (`id`,`aid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='快捷面板';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_panel`
--

LOCK TABLES `admin_panel` WRITE;
/*!40000 ALTER TABLE `admin_panel` DISABLE KEYS */;
/*!40000 ALTER TABLE `admin_panel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gard`
--

DROP TABLE IF EXISTS `gard`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gard` (
  `id` mediumint(6) unsigned NOT NULL AUTO_INCREMENT,
  `dict` varchar(32) NOT NULL COMMENT '区',
  `street` varchar(32) DEFAULT NULL COMMENT '街道',
  `childrengard` varchar(64) DEFAULT NULL COMMENT '幼儿园',
  `code` varchar(10) DEFAULT '0' COMMENT '编码',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='幼儿园表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gard`
--

LOCK TABLES `gard` WRITE;
/*!40000 ALTER TABLE `gard` DISABLE KEYS */;
INSERT INTO `gard` VALUES (1,'罗湖','','','02','0000-00-00 00:00:00'),(2,'南山','','','01','2016-04-20 17:07:07'),(3,'福田','','','03','2016-04-20 17:08:08'),(4,'宝安','','','04','2016-04-20 18:16:16'),(5,'龙岗','','','05','2016-04-20 18:16:16'),(6,'南山','蛇口','','0101','2016-04-20 18:17:17');
/*!40000 ALTER TABLE `gard` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logs`
--

DROP TABLE IF EXISTS `logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` int(10) unsigned NOT NULL COMMENT 'uid',
  `module` varchar(50) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '模型',
  `url` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作对应的url',
  `action` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作对应的action',
  `ip` varchar(15) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0.0.0.0' COMMENT '操作者所在IP',
  `desc` text COLLATE utf8_unicode_ci NOT NULL COMMENT '操作说明',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`,`module`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='后台操作日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logs`
--

LOCK TABLES `logs` WRITE;
/*!40000 ALTER TABLE `logs` DISABLE KEYS */;
INSERT INTO `logs` VALUES (1,1,'Logs','Logs.DelAll','DelAll','','清空日志|^|日志管理','2016-04-17 16:14:14'),(2,1,'User','User.Login','Login','','登陆用户名:admin|^|登陆系统!|^|登陆ID:1','2016-04-17 16:46:46'),(3,1,'User','User.Login','Login','','登陆用户名:admin|^|登陆系统!|^|登陆ID:1','2016-04-19 09:34:34'),(4,1,'Role','Role.Edit','Edit','','编辑角色:网站管理员|^|角色管理|^|ID:2','2016-04-19 09:36:36'),(5,1,'Menu','Menu.Edit','Edit','','编辑菜单:幼儿园管理|^|菜单管理|^|ID:36','2016-04-19 18:36:36'),(6,1,'Menu','Menu.Edit','Edit','','编辑菜单:幼儿园管理|^|菜单管理|^|ID:37','2016-04-19 18:37:37'),(7,1,'User','User.Login','Login','','登陆用户名:admin|^|登陆系统!|^|登陆ID:1','2016-04-20 10:45:45'),(8,1,'User','User.Login','Login','127.0.0.1','登陆用户名:admin|^|登陆系统!|^|登陆ID:1','2016-04-20 14:10:10'),(9,1,'Gard','Gard.Add','Add','','添加幼儿园:dfaf|^|幼儿园管理','2016-04-20 17:07:07'),(10,1,'Gard','Gard.Add','Add','','添加幼儿园:010102|^|幼儿园管理','2016-04-20 17:08:08'),(11,1,'Gard','Gard.Edit','Edit','','编辑幼儿园:|^|管理员管理','2016-04-20 18:13:13'),(12,1,'Gard','Gard.Edit','Edit','','编辑幼儿园:|^|管理员管理','2016-04-20 18:14:14'),(13,1,'Gard','Gard.Edit','Edit','','编辑幼儿园:01|^|管理员管理','2016-04-20 18:15:15'),(14,1,'Gard','Gard.Edit','Edit','','编辑幼儿园:02|^|管理员管理','2016-04-20 18:15:15'),(15,1,'Gard','Gard.Edit','Edit','','编辑幼儿园:03|^|管理员管理','2016-04-20 18:16:16'),(16,1,'Gard','Gard.Add','Add','','添加幼儿园:04|^|幼儿园管理','2016-04-20 18:16:16'),(17,1,'Gard','Gard.Add','Add','','添加幼儿园:05|^|幼儿园管理','2016-04-20 18:16:16'),(18,1,'Gard','Gard.Add','Add','','添加幼儿园:0101|^|幼儿园管理','2016-04-20 18:17:17'),(19,1,'Gard','Gard.Add','Add','','添加幼儿园:54|^|幼儿园管理','2016-04-20 18:17:17'),(20,1,'Gard','Gard.Delete','Delete','','删除幼儿园|^|ID:7','2016-04-20 18:20:20'),(21,1,'Menu','Menu.Delete','Delete','','删除菜单|^|ID:38','2016-04-20 18:20:20'),(22,1,'Menu','Menu.Delete','Delete','','删除菜单|^|ID:39','2016-04-20 18:20:20'),(23,1,'User','User.Login','Login','','登陆用户名:admin|^|登陆系统!|^|登陆ID:1','2016-04-22 17:20:20'),(24,1,'Menu','Menu.Edit','Edit','','编辑菜单:基本信息|^|菜单管理|^|ID:3','2016-04-22 17:22:22'),(25,1,'Menu','Menu.Edit','Edit','','编辑菜单:基本情况|^|菜单管理|^|ID:24','2016-04-22 17:23:23'),(26,1,'User','User.Login','Login','','登陆用户名:admin|^|登陆系统!|^|登陆ID:1','2016-04-27 14:00:00');
/*!40000 ALTER TABLE `logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `pid` int(11) NOT NULL DEFAULT '0',
  `name` char(40) NOT NULL DEFAULT '' COMMENT '名称',
  `enname` char(40) NOT NULL DEFAULT '' COMMENT '英文名称',
  `url` char(100) NOT NULL DEFAULT '' COMMENT '功能地址',
  `data` char(100) DEFAULT '' COMMENT '附加参数',
  `order` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `display` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示，1:显示 0:不显示',
  PRIMARY KEY (`id`),
  KEY `listorder` (`order`),
  KEY `parentid` (`pid`),
  KEY `module` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8 COMMENT='后台菜单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (1,0,'我的面板','Panel','Panel','',10000,1),(2,0,'设置','Settings','Setting','',20000,1),(3,0,'基本信息','Basicinfo','Basic','',30000,1),(4,0,'内容','Content','Content','',40000,1),(5,0,'用户','Users','User','',50000,1),(6,0,'扩展','Extensions','Extend','',60000,1),(7,0,'界面','Templates','Style','',70000,1),(8,0,'应用','Plugin','Plugin','',80000,1),(9,2,'菜单设置','Menu Settings','javascript:;','',20100,1),(10,9,'菜单管理','Menu management','Menu','',20101,1),(11,1,'个人设置','Personal Settings','javascript:;','',10100,1),(12,11,'个人信息','Personal information','EditInfo','',10101,1),(13,11,'修改密码','Change password','EditPwd','',10102,1),(14,2,'管理员管理','Administrator manager','javascript:;','',20200,1),(15,14,'管理员管理','Administrator manager','Admin','',20201,1),(16,14,'角色管理','Role management','Role','',20202,1),(17,2,'日志管理','Log management','javascript:;','',20400,1),(18,17,'日志管理','Log management','Logs','',20401,1),(19,1,'快捷面板','Shortcut panel','javascript:;','',10200,1),(20,4,'内容管理','Content management','javascript:;','',40100,1),(21,4,'焦点图管理','Focus Management','javascript:;','',40200,1),(22,20,'栏目管理','Manage column','Category','',40101,1),(23,20,'内容管理','Manage content','Content','',40102,1),(24,3,'基本情况','Basicinfo','javascript:;','',30100,1),(25,24,'公告','Announcement','Announce','',30101,1),(26,6,'扩展','Extensions','javascript:;','',60100,1),(27,26,'来源管理','Source management','Copyfrom','',60101,1),(28,5,'会员管理','Manage user','javascript:;','',50100,1),(29,5,'会员组管理','Manage user group','javascript:;','',50200,1),(30,28,'会员管理','Manage user','User','',50101,1),(31,29,'管理会员组','Manage user group','Group','',50201,1),(32,7,'模板管理','Manage template','javascript:;','',70100,1),(33,32,'模板风格','Style template','Style','',70101,1),(34,21,'焦点图列表','Focus List','Focus','',40201,1),(35,21,'分类管理','Focus Cate','FocusCate','',40202,1),(36,2,'幼儿园管理','Children Gard Manage','javascript:;','',20300,1),(37,36,'幼儿园管理','Children Gard Manage','Gard','',20301,1),(40,24,'投诉建议','Complaints proposals','Complaints','',30102,1);
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `model`
--

DROP TABLE IF EXISTS `model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `model` (
  `id` mediumint(6) unsigned NOT NULL AUTO_INCREMENT,
  `dict` varchar(32) NOT NULL COMMENT '区',
  `street` varchar(32) DEFAULT NULL COMMENT '街道',
  `childrengard` varchar(64) DEFAULT NULL COMMENT '幼儿园',
  `code` varchar(10) DEFAULT '0' COMMENT '编码',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='幼儿园表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `model`
--

LOCK TABLES `model` WRITE;
/*!40000 ALTER TABLE `model` DISABLE KEYS */;
INSERT INTO `model` VALUES (1,'23','33','33','33','2010-01-01 00:00:00');
/*!40000 ALTER TABLE `model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` int(3) unsigned NOT NULL AUTO_INCREMENT,
  `rolename` varchar(50) NOT NULL COMMENT '角色名称',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '角色说明',
  `data` text NOT NULL COMMENT '菜单列表',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否启用',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `roleid` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'超级管理员','超级管理员','1,11,12,13,2,9,10,14,15,16,17,18,3,4,5,6,7,8',1,'2014-01-18 00:09:09'),(2,'网站管理员','网站管理员','1,11,12,13,2,9,10,14,15,16,17,18,3,5,6,7,8',1,'2014-02-10 22:18:18');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `email` char(32) NOT NULL DEFAULT '' COMMENT '电子邮箱',
  `username` char(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `encrypt` char(6) NOT NULL DEFAULT '' COMMENT '随机码',
  `nickname` char(20) NOT NULL DEFAULT '' COMMENT '昵称',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机',
  `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
  `regip` char(15) NOT NULL DEFAULT '' COMMENT '注册ip',
  `regdate` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '注册时间',
  `lastdate` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `lastip` char(15) NOT NULL DEFAULT '' COMMENT '上次登录ip',
  `loginnum` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '登陆次数',
  `groupid` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户组id',
  `areaid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '地区id',
  `amount` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '金钱总额',
  `point` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `ismessage` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否有短消息',
  `islock` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否定锁',
  `vip` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'vip等级',
  `overduedate` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'vip过期时间',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '审核状态 5:用户名已存在;4:拒绝;3:删除:2:忽略;0:未审核;1:通过',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `email` (`email`(20))
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_group`
--

DROP TABLE IF EXISTS `user_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_group` (
  `id` tinyint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '会员组id',
  `name` char(15) NOT NULL DEFAULT '' COMMENT '用户组名称',
  `issystem` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是系统组',
  `star` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '会员组星星数',
  `point` smallint(6) unsigned NOT NULL DEFAULT '0' COMMENT '积分范围',
  `allowmessage` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '允许发短消息数量',
  `allowvisit` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否允许访问',
  `allowpost` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否允许发稿',
  `allowpostverify` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否投稿不需审核',
  `allowsearch` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否允许搜索',
  `allowupgrade` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否允许自主升级',
  `allowsendmessage` tinyint(1) unsigned NOT NULL COMMENT '允许发送短消息',
  `allowpostnum` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '每天允许发文章数',
  `allowattachment` tinyint(1) NOT NULL COMMENT '是否允许上传附件',
  `priceyear` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '包年价格',
  `pricemonth` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '包月价格',
  `priceday` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '包天价格',
  `icon` char(30) NOT NULL COMMENT '用户组图标',
  `usernamecolor` char(7) NOT NULL COMMENT '用户名颜色',
  `desc` char(100) NOT NULL COMMENT '描述',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否禁用',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_group`
--

LOCK TABLES `user_group` WRITE;
/*!40000 ALTER TABLE `user_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'xin_cg'
--

--
-- Dumping routines for database 'xin_cg'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-04-28 16:49:33
