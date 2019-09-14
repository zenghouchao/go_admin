/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : go_web

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-09-14 09:51:04
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for go_admin
-- ----------------------------
DROP TABLE IF EXISTS `go_admin`;
CREATE TABLE `go_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `pass` varchar(70) NOT NULL COMMENT '密码',
  `time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '新增时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='管理员表';

-- ----------------------------
-- Records of go_admin
-- ----------------------------
INSERT INTO `go_admin` VALUES ('1', 'admin', '86050074ffe28cd5528504df63d8e873', '1566743400');

-- ----------------------------
-- Table structure for go_article
-- ----------------------------
DROP TABLE IF EXISTS `go_article`;
CREATE TABLE `go_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cateId` int(10) unsigned NOT NULL COMMENT '文章栏目ID',
  `title` varchar(100) NOT NULL COMMENT '文章标题',
  `content` text NOT NULL,
  `time` int(11) unsigned NOT NULL COMMENT '发布时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '文章状态',
  `author` varchar(35) NOT NULL COMMENT '作者',
  PRIMARY KEY (`id`),
  KEY `index_cate` (`cateId`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=21 DEFAULT CHARSET=utf8 COMMENT='文章表';

-- ----------------------------
-- Records of go_article
-- ----------------------------
INSERT INTO `go_article` VALUES ('2', '26', '产业互联网：企业前所未有的战略机遇', '在2019中国500强企业高峰论坛“产业互联网，企业成长新动能”分论坛上，金蝶国际软件集团董事会主席徐少春认为，产业互联网具有以下三个方面的特征：第一，互联网等新技术重在重塑各行各业的价值链；第二，产业互联网正在重塑每个企业的核心竞争能力；第三，新资源方面，现在更多强调人才和无形资源。他指出，产业互联网的新技术使得一切的应用都变成了云化、移动化、智能化、场景化、生态化和人文化', '1567353600', '1', 'admin');
INSERT INTO `go_article` VALUES ('8', '37', '勇做敢于斗争善于斗争的战士', '9月3日，习近平总书记在中央党校（国家行政学院）中青年干部培训班开班式上发表重要讲话强调，广大干部特别是年轻干部要经受严格的思想淬炼、政治历练、实践锻炼，发扬斗争精神，增强斗争本领，为实现“两个一百年”奋斗目标、实现中华民族伟大复兴的中国梦而顽强奋斗。', '1567526400', '1', '央视快评');
INSERT INTO `go_article` VALUES ('9', '24', '教育部、公安部召开紧急电视电话会议再部署学校安全工作', '9月4日，教育部、公安部联合召开全国学校安全工作紧急电视电话会议，对中小学、幼儿园安全工作再次作出明确部署。\n会议提出了六点要求。一要把安全红线意识树牢。教育系统各级领导和干部一定要把确保安全作为教育事业发展和学生成长成才的最底线要求，把广大师生的生命安全放在第一位。不管什么时候谈起教育，首先要谈安全；到学校检查指导工作，一定要先去看安全工作到不到位', '1567526400', '0', '新京报');
INSERT INTO `go_article` VALUES ('4', '27', '强国图志——国运兴体育兴', '体育强则中国强， 国运兴则体育兴。 祝福祖国国家强盛！民族振兴！[玫瑰][玫瑰]', '1567267200', '1', '今日头条');
INSERT INTO `go_article` VALUES ('10', '26', '华为鲲鹏高频亮相，进入英特尔独大的服务器芯片市场', 'ARM在移动端一骑绝尘，但是在服务器领域构建生态是难题，虽然多年来高通、英伟达、三星等大厂均尝试建立ARM生态，但从成效来看，几乎可以忽略不计。\n\n这里面除了技术和性能的因素外，更重要的还有成本与生态的考验。弗罗斯特研究（ForresterResearch）首席分析师戴鲲曾对记者表示，对于这个领域的新入局者，实现规模化效益存在重重阻力。“服务器芯片市场需要长期的技术投资与软硬件生态系统的广泛支持。', '1567353600', '1', '第一财经');
INSERT INTO `go_article` VALUES ('11', '38', '刘涛带儿子女儿拍杂志，两个宝贝长得也太像王珂了吧', '近日，刘涛带着女儿和儿子一同拍的杂志封面曝光。照片中刘涛身穿一套浅色森女风的衣服，刘涛的儿子穿黄色衬衫搭配黑色裤子，而刘涛的女儿则穿一条格子连衣长裙，和弟弟的穿搭很配', '1567526400', '0', '谈资');
INSERT INTO `go_article` VALUES ('12', '37', '郭台铭：希望出来服务 港媒：与国民党接近摊牌', '郭台铭3日称“希望出来为大家服务”、“我出来就是希望把台湾经济翻转”——就差宣布参选了。\n中评社今发文说，郭台铭、柯文哲、王金平是否“结盟”，底牌在9月17日就要揭晓，也是郭、王和国民党的摊牌日。\n\n目前国民党仍在努力拉回郭。有人提醒，他参加了国民党初选，输了。', '1567612800', '1', '你好台湾网');
INSERT INTO `go_article` VALUES ('13', '26', '研究了微博最新推出的绿洲App后，一位实习生写下了这些', '微博又在社交领域里有新动作了。\n\n8月29日，由微博团队全新打造的社交App——绿洲，正式开启内测。虽然只是内测，但还是吸引了众多的关注，并于9月3日登上了App Store 社交类排行榜第一（最新消息：9月4日下午2点，绿洲已从App Store下架，原因为logo涉嫌抄袭）', '1567526400', '1', '司林');
INSERT INTO `go_article` VALUES ('14', '26', '已成功还款218亿元是回国征兆？贾跃亭辞任法拉第未来CEO', '9月3日，电动汽车公司法拉第未来宣布正式任命毕福康博士（Dr·Carsten Breitfeld）为全球CEO，而FF创始人贾跃亭辞去CEO职务，出任CPUO（首席产品和用户官）', '1567612800', '0', '聚牛科技');
INSERT INTO `go_article` VALUES ('15', '24', '中小学老师是高薪职业吗？深圳应届生税前年收入30万+', '中小学老师是高薪职业吗？\n\n9月3日，在教育部举行的新闻发布会上，教育部教师工作司司长任友群透露，我国教师工资由上世纪80年代之前在国民经济各行业排行倒数后三位，上升到目前全国19大行业排名第7位。\n\n对于大学老师而言，在工资之外，收入来源一般包括科研经费、项目经费以及科研奖励等，待遇的个体差异较大，难以做出直接比较，但对于中小学老师而言，待遇究竟如何？', '1567526400', '1', '21世纪经济报道');
INSERT INTO `go_article` VALUES ('16', '24', '中科院研究生被杀的背后 不止嫉妒这么简单', '2019年6月14日，小雨下了一整天，重庆垫江的一处墓地上，雷燕和谢中华在雨中不停落泪，今天是儿子去世一周年忌日，他们眼前的墓碑上写着\"中科院研究生谢雕之墓\"。\n\n雷燕哭喊着儿子的名字，不停用手擦去遗像上的雨滴，雨滴又不停地落上。\n\n谢中华默默流泪，他将事先准备好的菜和水果摆在儿子碑前，拿出了白酒和酒杯。\n\n2018年春节，谢雕本想和父亲喝一杯，但考虑到父亲的鼻咽癌不能饮酒，所以父子俩约定，等父亲的病好了再喝', '1567526400', '1', 'Tina心理');
INSERT INTO `go_article` VALUES ('17', '24', '商务部：中方坚决反对贸易战升级', '今天（9月5日）下午三点，商务部召开例行新闻发布会。商务部新闻发言人高峰就近期热点话题回答中外记者提问。\n\n高峰表示，中方坚决反对贸易战升级，这不利于中国，不利于美国，也不利于全世界，相信大家已经注意到，在今天上午的通话中，双方一致认为，应共同努力，采取实际行动，为下一步磋商创造良好的条件。', '1567612800', '1', '央视网');
INSERT INTO `go_article` VALUES ('18', '26', '中国8.54亿网民学历结构：约九成网民学历不足本科', '中国经济周刊-经济网讯 8月30日，中国互联网络信息中心CNNIC发布《中国互联网络发展状况统计报告》，截止到2019年6月，中国网民规模达8.54亿。\n\n报告显示我国网民的人均每周上网时长为27.9小时，较2018年底增加0.3小时。\n\n其中近5成的网民年龄为30岁以下年轻群体，本科以下学历(不含本科)网民占比为90.3%，而月收入在5000元以下的网民群体合计占比为72.8%', '1567612800', '0', '中国经济周刊');
INSERT INTO `go_article` VALUES ('19', '26', '习近平对全国道德模范表彰活动作出重要指示', '新华社北京9月5日电 中共中央总书记、国家主席、中央军委主席习近平近日对全国道德模范表彰活动作出重要指示，向受表彰的全国道德模范致以热烈的祝贺。\n\n习近平指出，在新中国成立70周年之际，中央文明委评选表彰新一届全国道德模范，这对倡导好风尚、弘扬正能量、促进全社会向上向善具有十分重要的意义', '1567612800', '1', '新华网客户端');
INSERT INTO `go_article` VALUES ('20', '24', '数次会见德国总理 习近平讲话释放这些重要信息', '9月6日，习近平主席在北京会见德国总理默克尔。“中方扩大开放说到做到”“欢迎德国企业参与长江经济带建设”……在会谈阶段，习近平提出多项务实的合作项目和合作方向，再一次阐明中国坚持扩大开放的决心。\n\n“当前，我们正处在一个挑战层出不穷、风险日益增多的时代，单边主义、保护主义严重威胁世界和平稳定，任何国家都不能独善其身。”长久以来，中国和德国作为两个负责任的大国，双方朝着互利合作的方向携手前行，一系列合作成果惠及两国人民，中德关系也随之不断迈向新高度', '1567872000', '0', '央视网新闻');

-- ----------------------------
-- Table structure for go_cate
-- ----------------------------
DROP TABLE IF EXISTS `go_cate`;
CREATE TABLE `go_cate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '栏目名称',
  `status` tinyint(2) unsigned NOT NULL COMMENT '栏目状态',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=54 DEFAULT CHARSET=utf8 COMMENT='文章栏目表';

-- ----------------------------
-- Records of go_cate
-- ----------------------------
INSERT INTO `go_cate` VALUES ('24', '新闻栏目', '1');
INSERT INTO `go_cate` VALUES ('39', '政治新闻', '0');
INSERT INTO `go_cate` VALUES ('25', '快递栏目', '1');
INSERT INTO `go_cate` VALUES ('26', '图腾栏目', '1');
INSERT INTO `go_cate` VALUES ('27', '体育栏目', '1');
INSERT INTO `go_cate` VALUES ('42', '影视资讯', '0');
INSERT INTO `go_cate` VALUES ('41', '社会新闻', '0');
INSERT INTO `go_cate` VALUES ('40', '农业新闻', '0');
INSERT INTO `go_cate` VALUES ('37', '政治论坛', '1');
INSERT INTO `go_cate` VALUES ('38', '娱乐新闻', '1');
INSERT INTO `go_cate` VALUES ('43', '教育栏目', '1');
INSERT INTO `go_cate` VALUES ('44', '培训智讯', '0');
INSERT INTO `go_cate` VALUES ('45', '工程技术', '0');
INSERT INTO `go_cate` VALUES ('46', '文学艺术', '0');
INSERT INTO `go_cate` VALUES ('47', '社会科学', '0');
INSERT INTO `go_cate` VALUES ('48', '生物科学', '0');
INSERT INTO `go_cate` VALUES ('49', '地理频道', '0');
INSERT INTO `go_cate` VALUES ('50', '大千世界', '0');
INSERT INTO `go_cate` VALUES ('51', '百度一下', '0');
INSERT INTO `go_cate` VALUES ('52', '诗词歌赋', '1');
INSERT INTO `go_cate` VALUES ('53', '人生哲学', '1');
