/*
 LeastMall SQL
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` int(10) DEFAULT '0' COMMENT '用户编号',
  `phone` varchar(30) DEFAULT '' COMMENT '用户手机',
  `name` varchar(30) DEFAULT '' COMMENT '用户名字',
  `zipcode` varchar(20) DEFAULT '' COMMENT '邮政编码',
  `address` varchar(250) DEFAULT '' COMMENT '地址',
  `default_address` int(11) DEFAULT '0' COMMENT '默认地址',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='地址信息';

-- ----------------------------
-- Records of address
-- ----------------------------
BEGIN;
INSERT INTO `address` VALUES (1, 1, '13888888888', '张三', '600000', '成都市xxxx区xxxx街道xxxx号', 0, 0);
INSERT INTO `address` VALUES (2, 1, '13888888889', '王五', '100000', '北京xxxxx区xxxxx街道xxxx号', 1, 0);
COMMIT;

-- ----------------------------
-- Table structure for administrator
-- ----------------------------
DROP TABLE IF EXISTS `administrator`;
CREATE TABLE `administrator` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) DEFAULT '' COMMENT '用户名',
  `password` varchar(100) DEFAULT '' COMMENT '密码',
  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `role_id` int(10) DEFAULT '0' COMMENT '角色编号',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  `is_super` tinyint(4) DEFAULT '0' COMMENT '是否超级管理员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='管理员表';

-- ----------------------------
-- Records of administrator
-- ----------------------------
BEGIN;
INSERT INTO `administrator` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '1888888888', 'admin@163.com', 1, 1, 0, 1);
INSERT INTO `administrator` VALUES (2, '工程师', '', '13999999999', 'programmer1@163.com', 1, 2, 1603517980, 0);
COMMIT;

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `module_name` varchar(80) NOT NULL DEFAULT '',
  `action_name` varchar(80) DEFAULT '' COMMENT '操作名称',
  `type` tinyint(1) DEFAULT '0' COMMENT '节点类型',
  `url` varchar(250) DEFAULT '' COMMENT '跳转地址',
  `module_id` int(11) DEFAULT '0' COMMENT '模块编号',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `description` varchar(250) DEFAULT '' COMMENT '描述',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  `checked` tinyint(1) DEFAULT '0' COMMENT '是否检验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='权限控制';

-- ----------------------------
-- Records of auth
-- ----------------------------
BEGIN;
INSERT INTO `auth` VALUES (1, '系统管理员', 'get', 3, '3', 0, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (2, '组织部门', '', 3, '', 0, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (3, '模块管理', '', 3, '', 0, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (4, 'Banner管理', '', 3, '', 0, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (5, '商品管理', '', 3, '', 0, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (6, '订单管理', '', 3, '', 0, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (7, '设置管理', '', 3, '', 0, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (8, '管理员列表', '管理员列表', 2, '/administrator', 1, 0, '', 0, 0, 1);
INSERT INTO `auth` VALUES (9, '', '新增管理员', 2, '/administrator/add', 1, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (10, '', '部门列表', 2, '/role', 2, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (11, '', '新增部门', 2, '/role/add', 2, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (12, '', '新增权限', 2, '/auth/add', 3, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (13, '', '权限列表', 2, '/auth', 3, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (14, '', 'Banner列表', 2, '/banner', 4, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (15, '', '新增Banner', 2, '/banner/add', 4, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (16, '', '商品列表', 2, '/product', 5, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (17, '', '商品分类', 2, '/productCate', 5, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (18, '', '商品类型', 2, '/productType', 5, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (19, '', '订单列表', 2, '/order', 6, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (20, '', '导航管理', 2, '/menu', 7, 0, '', 0, 0, 0);
INSERT INTO `auth` VALUES (21, '', '商城设置', 2, '/setting', 7, 0, '', 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(50) DEFAULT '' COMMENT '标题',
  `banner_type` tinyint(4) DEFAULT '0' COMMENT '类型',
  `banner_img` varchar(100) DEFAULT '' COMMENT '图片地址',
  `link` varchar(200) DEFAULT '' COMMENT '连接',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` int(10) DEFAULT '0' COMMENT '状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='焦点图表';

-- ----------------------------
-- Records of banner
-- ----------------------------
BEGIN;
INSERT INTO `banner` VALUES (1, 'banner1', 1, 'static/upload/20201121/1605944396119681000.jpg', '/', 1000, 1, 1603504343);
INSERT INTO `banner` VALUES (2, 'banner2', 1, 'static/upload/20201024/1603504839411539000.jpg', '', 1000, 1, 1603504839);
COMMIT;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(250) DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) DEFAULT '0.00',
  `goods_version` varchar(50) DEFAULT '' COMMENT '版本',
  `num` int(11) DEFAULT '0' COMMENT '数量',
  `product_gift` varchar(100) DEFAULT '' COMMENT '商品礼物',
  `product_fitting` varchar(100) DEFAULT '' COMMENT '商品搭配',
  `product_color` varchar(50) DEFAULT '' COMMENT '商品颜色',
  `product_img` varchar(150) DEFAULT '' COMMENT '商品图片',
  `product_attr` varchar(100) DEFAULT '' COMMENT '商品属性',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='购物车';

-- ----------------------------
-- Records of cart
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `link` varchar(250) DEFAULT '' COMMENT '连接',
  `position` int(10) DEFAULT '0' COMMENT '位置',
  `is_opennew` int(10) DEFAULT '0' COMMENT '是否新打开',
  `relation` varchar(100) DEFAULT '' COMMENT '关系',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` int(10) DEFAULT '0' COMMENT '状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, 'PC电脑', '/category_1.html', 1, 1, '1', 10, 1, 0);
INSERT INTO `menu` VALUES (2, '手机', '', 1, 2, '1', 10, 1, 1603518198);
COMMIT;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `order_id` varchar(100) DEFAULT '' COMMENT '订单编号',
  `uid` int(11) DEFAULT '0' COMMENT '用户编号',
  `all_price` decimal(10,2) DEFAULT '0.00' COMMENT '价格',
  `phone` varchar(30) DEFAULT '' COMMENT '电话',
  `name` varchar(100) DEFAULT '' COMMENT '名字',
  `address` varchar(250) DEFAULT '' COMMENT '地址',
  `zipcode` varchar(30) DEFAULT '' COMMENT '邮编',
  `pay_status` tinyint(4) DEFAULT '0' COMMENT '支付状态',
  `pay_type` tinyint(4) DEFAULT '0' COMMENT '支付类型',
  `order_status` tinyint(4) DEFAULT '0' COMMENT '订单状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order
-- ----------------------------
BEGIN;
INSERT INTO `order` VALUES (1, '2020102316351850', 1, 66.66, '18899999999', '张三', '成都市高新区', '600000', 0, 0, 0, 1603442151);
INSERT INTO `order` VALUES (2, '2020112016527196', 1, 6666.66, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, 1605862363);
INSERT INTO `order` VALUES (3, '2021011810071850', 1, 1999.00, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, 1610935652);
INSERT INTO `order` VALUES (4, '2021011810151576', 1, 1999.00, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, 1610936113);
INSERT INTO `order` VALUES (5, '2021011811371850', 1, 10995.00, '13888888889', '王五', '北京xxxxx区xxxxx街道xxxx号', '100000', 0, 0, 0, 1610941064);
COMMIT;

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '订单编号',
  `order_id` int(11) DEFAULT '0' COMMENT '订单编号',
  `uid` int(11) DEFAULT '0' COMMENT '用户编号',
  `product_title` varchar(100) DEFAULT '' COMMENT '商品标题',
  `product_id` int(11) DEFAULT '0' COMMENT '商品编号',
  `product_img` varchar(200) DEFAULT '' COMMENT '商品图片',
  `product_price` decimal(10,2) DEFAULT '0.00' COMMENT '商品价格',
  `product_num` int(10) DEFAULT '0' COMMENT '商品数量',
  `goods_version` varchar(100) DEFAULT '' COMMENT '商品版本',
  `goods_color` varchar(100) DEFAULT '' COMMENT '商品颜色',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order_item
-- ----------------------------
BEGIN;
INSERT INTO `order_item` VALUES (1, 1, 1, 'go web from ', 1, 'static/upload/20201023/1603440321653795000.png', 66.66, 1, '2', 'red', 1603442151);
COMMIT;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `sub_title` varchar(100) DEFAULT '' COMMENT '子标题',
  `product_sn` varchar(50) DEFAULT '',
  `cate_id` int(10) DEFAULT '0' COMMENT '分类id',
  `click_count` int(10) DEFAULT '0' COMMENT '点击数',
  `product_number` int(10) DEFAULT '0' COMMENT '商品编号',
  `price` decimal(10,2) DEFAULT '0.00' COMMENT '价格',
  `market_price` decimal(10,2) DEFAULT '0.00' COMMENT '市场价格',
  `relation_product` varchar(100) DEFAULT '' COMMENT '关联商品',
  `product_attr` varchar(100) DEFAULT '' COMMENT '商品属性',
  `product_version` varchar(100) DEFAULT '' COMMENT '商品版本',
  `product_img` varchar(100) DEFAULT '' COMMENT '商品图片',
  `product_gift` varchar(100) DEFAULT '',
  `product_fitting` varchar(100) DEFAULT '',
  `product_color` varchar(100) DEFAULT '' COMMENT '商品颜色',
  `product_keywords` varchar(100) DEFAULT '' COMMENT '关键词',
  `product_desc` varchar(50) DEFAULT '' COMMENT '描述',
  `product_content` varchar(100) DEFAULT '' COMMENT '内容',
  `is_delete` tinyint(4) DEFAULT '0' COMMENT '是否删除',
  `is_hot` tinyint(4) DEFAULT '0' COMMENT '是否热门',
  `is_best` tinyint(4) DEFAULT '0' COMMENT '是否畅销',
  `is_new` tinyint(4) DEFAULT '0' COMMENT '是否新品',
  `product_type_id` tinyint(4) DEFAULT '0' COMMENT '商品类型编号',
  `sort` int(10) DEFAULT '0' COMMENT '商品分类',
  `status` tinyint(4) DEFAULT '0' COMMENT '商品状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='商品';

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` VALUES (1, '电脑', '', '', 1, 100, 0, 1999.00, 2599.00, '', '', '第一版', 'static/upload/20201203/1607005318313835000.jpg', '', '', '1', '', '', '', 0, 1, 1, 0, 0, 0, 1, 1603440139);
INSERT INTO `product` VALUES (2, '笔记本', '', '', 1, 1, 5, 3999.00, 4699.00, '23,24,39', ' 格式: 颜色:红色,白色,黄色 | 尺寸:41,42,43', '', 'static/upload/20210118/1610940762322803000.jpg', '', '', '1', '', '', '', 0, 1, 1, 1, 1, 0, 1, 0);
INSERT INTO `product` VALUES (3, '手机', '', '', 1, 100, 6, 999.00, 1299.00, '', '', '', 'static/upload/20210118/1610940776037983000.jpg', '', '', '1', '', '', '', 0, 1, 1, 1, 0, 1, 1, 1607005027);
COMMIT;

-- ----------------------------
-- Table structure for product_attr
-- ----------------------------
DROP TABLE IF EXISTS `product_attr`;
CREATE TABLE `product_attr` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT '0' COMMENT '商品编号',
  `attribute_cate_id` int(10) DEFAULT '0' COMMENT '属性分类编号',
  `attribute_id` int(10) DEFAULT '0' COMMENT '属性编号',
  `attribute_title` varchar(100) DEFAULT '' COMMENT '属性标题',
  `attribute_type` int(10) DEFAULT '0' COMMENT '属性类型',
  `attribute_value` varchar(100) DEFAULT '' COMMENT '属性值',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='商品属性';

-- ----------------------------
-- Records of product_attr
-- ----------------------------
BEGIN;
INSERT INTO `product_attr` VALUES (12, 2, 1, 1, '平板电脑', 1, '', 10, 1610940762, 1);
COMMIT;

-- ----------------------------
-- Table structure for product_cate
-- ----------------------------
DROP TABLE IF EXISTS `product_cate`;
CREATE TABLE `product_cate` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) DEFAULT '' COMMENT '分类名称',
  `cate_img` varchar(200) DEFAULT '' COMMENT '分类图片',
  `link` varchar(250) DEFAULT '' COMMENT '链接',
  `template` text COMMENT '模版',
  `pid` int(10) DEFAULT '0' COMMENT '父编号',
  `sub_title` varchar(100) DEFAULT '' COMMENT '子标题',
  `keywords` varchar(250) DEFAULT '' COMMENT '关键字',
  `description` text COMMENT '描述',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='商品分类';

-- ----------------------------
-- Records of product_cate
-- ----------------------------
BEGIN;
INSERT INTO `product_cate` VALUES (1, '手机', '', '', '', 0, '手机', '手机', '手机', 0, 1, 0);
INSERT INTO `product_cate` VALUES (2, 'PC电脑', '', '', '', 0, 'PC电脑', '', '', 0, 1, 0);
COMMIT;

-- ----------------------------
-- Table structure for product_color
-- ----------------------------
DROP TABLE IF EXISTS `product_color`;
CREATE TABLE `product_color` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `color_name` varchar(100) DEFAULT '' COMMENT '颜色名字',
  `color_value` varchar(100) DEFAULT '' COMMENT '颜色值',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `checked` tinyint(4) DEFAULT '0' COMMENT '是否检验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product_color
-- ----------------------------
BEGIN;
INSERT INTO `product_color` VALUES (1, '黑色', '#ffffff', 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for product_image
-- ----------------------------
DROP TABLE IF EXISTS `product_image`;
CREATE TABLE `product_image` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT '0' COMMENT '商品编号',
  `img_url` varchar(250) DEFAULT '' COMMENT '图片地址',
  `color_id` int(11) DEFAULT '0' COMMENT '颜色编号',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product_image
-- ----------------------------
BEGIN;
INSERT INTO `product_image` VALUES (1, 1, '/static/upload/20201024/1603519200684359000.jpg', 0, 10, 1603519201, 1);
INSERT INTO `product_image` VALUES (2, 1, '/static/upload/20201024/1603519285204437000.jpg', 0, 10, 1603519291, 1);
INSERT INTO `product_image` VALUES (6, 2, '/static/upload/20210118/1610940522542324000.jpg', 0, 10, 1610940523, 1);
INSERT INTO `product_image` VALUES (7, 2, '/static/upload/20210118/1610940522573123000.jpg', 0, 10, 1610940523, 1);
INSERT INTO `product_image` VALUES (8, 3, '/static/upload/20210118/1610940548355473000.jpg', 0, 10, 1610940548, 1);
COMMIT;

-- ----------------------------
-- Table structure for product_type
-- ----------------------------
DROP TABLE IF EXISTS `product_type`;
CREATE TABLE `product_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `description` varchar(500) DEFAULT '' COMMENT '描述',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product_type
-- ----------------------------
BEGIN;
INSERT INTO `product_type` VALUES (1, '电脑', '', 1, 0);
COMMIT;

-- ----------------------------
-- Table structure for product_type_attribute
-- ----------------------------
DROP TABLE IF EXISTS `product_type_attribute`;
CREATE TABLE `product_type_attribute` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cate_id` int(11) DEFAULT '0' COMMENT '分类编号',
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `attr_type` tinyint(4) DEFAULT '0' COMMENT '属性类型',
  `attr_value` varchar(100) DEFAULT '' COMMENT '属性值',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of product_type_attribute
-- ----------------------------
BEGIN;
INSERT INTO `product_type_attribute` VALUES (1, 1, '平板电脑', 1, '', 1, 10, 1603440086);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '标题名称',
  `description` varchar(500) DEFAULT '' COMMENT '描述',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, '超级管理员', '超级管理员', 1, 0);
INSERT INTO `role` VALUES (2, '技术部', '技术部', 1, 1603518015);
INSERT INTO `role` VALUES (3, '运营部', '运营部', 1, 1603518054);
COMMIT;

-- ----------------------------
-- Table structure for role_auth
-- ----------------------------
DROP TABLE IF EXISTS `role_auth`;
CREATE TABLE `role_auth` (
  `auth_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限编号',
  `role_id` int(11) DEFAULT '0' COMMENT '角色编号',
  PRIMARY KEY (`auth_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role_auth
-- ----------------------------
BEGIN;
INSERT INTO `role_auth` VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `site_title` varchar(100) DEFAULT '' COMMENT '商城名称',
  `site_logo` varchar(250) DEFAULT '' COMMENT '商城图标',
  `site_keywords` varchar(100) DEFAULT '' COMMENT '商城关键字',
  `site_description` varchar(500) DEFAULT '' COMMENT '商城描述',
  `no_picture` varchar(100) DEFAULT '' COMMENT '没有图片显示',
  `site_icp` varchar(50) DEFAULT '' COMMENT '商城ICP',
  `site_tel` varchar(50) DEFAULT '' COMMENT '商城手机号',
  `search_keywords` varchar(250) DEFAULT '' COMMENT '搜索关键字',
  `tongji_code` varchar(500) DEFAULT '' COMMENT '统计编码',
  `appid` varchar(50) DEFAULT '' COMMENT 'oss appid',
  `app_secret` varchar(80) DEFAULT '' COMMENT 'oss app_secret',
  `end_point` varchar(200) DEFAULT '' COMMENT 'oss 终端点',
  `bucket_name` varchar(200) DEFAULT '' COMMENT 'oss 桶名称',
  `oss_status` tinyint(4) DEFAULT '0' COMMENT 'oss 状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of setting
-- ----------------------------
BEGIN;
INSERT INTO `setting` VALUES (1, 'LeastMall商城', 'ab', 'LeastMall商城', 'LeastMall商城', 'a', 'LeastMall商城', '88888888', 'LeastMall商城', 'dd', 'b', 'c', 'd', 'a', 1);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(30) DEFAULT '' COMMENT '手机号',
  `password` varchar(80) DEFAULT '' COMMENT '密码',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  `last_ip` varchar(50) DEFAULT '' COMMENT '最近ip',
  `email` varchar(80) DEFAULT '' COMMENT '邮编',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '13888888888', 'e10adc3949ba59abbe56e057f20f883e', 0, '', 'admin@qq.com', 1);
INSERT INTO `user` VALUES (2, '18389999993', 'e10adc3949ba59abbe56e057f20f883e', 0, '127.0.0.1', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for user_sms
-- ----------------------------
DROP TABLE IF EXISTS `user_sms`;
CREATE TABLE `user_sms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(50) DEFAULT '' COMMENT 'ip地址',
  `phone` varchar(50) DEFAULT '' COMMENT '手机号',
  `send_count` int(10) DEFAULT '0' COMMENT '发送统计',
  `add_day` varchar(200) DEFAULT '' COMMENT '添加日期',
  `add_time` int(10) DEFAULT '0' COMMENT '添加时间',
  `sign` varchar(80) DEFAULT '' COMMENT '签名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_sms
-- ----------------------------
BEGIN;
INSERT INTO `user_sms` VALUES (1, '127.0.0.1', '18389999993', 1, '20201102', 1604288606, 'e178c966721a75236355d935ac3dd9ff');
INSERT INTO `user_sms` VALUES (2, '127.0.0.1', '18389999991', 1, '20201119', 1605759728, '200f0a43fbc9c0ae40f26432269cae91');
INSERT INTO `user_sms` VALUES (3, '127.0.0.1', '18389999992', 1, '20201119', 1605759900, '5abc5dca4b31c1cfe222693ee1c5bd1c');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
