## 数据库创建
### 创建数据库
```
create database blog_service default character set utf8mb4 default collate utf8mb4_general_ci
```

### 创建标签表
```
create table `blog_tag`(
    `id` int(10) unsigned not null auto_increment,
    `name` varchar(100) default '' comment '标签名称',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除，0表示未删除，1表示删除',
    `state` tinyint(3) unsigned default '0' comment '状态，0为禁用，1表示启用',
    primary key(`id`) 
) engine=InnoDB default charset=utf8mb4 comment '标签管理'
```

### 创建文章表
```
create table `blog_article`(
    `id` int(10) unsigned not null auto_increment,
    `title` varchar(100) default '' comment '文章标题',
    `desc` varchar(100) default '' comment '文章简述',
    `cover_image_url` varchar(255) default '' comment '封面图片地址',
    `content` longtext comment '文章内容',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除，0表示未删除，1表示删除',
    `state` tinyint(3) unsigned default '0' comment '状态，0为禁用，1表示启用',
    primary key(`id`) 
) engine=InnoDB default charset=utf8mb4 comment '文章管理'
```

### 创建文章标签关联表
```
create table `blog_article_tag`(
    `id` int(10) unsigned not null auto_increment,
    `article_id` int(10) not null comment '文章ID',
    `tag_id` int(10) not null comment '标签ID',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除，0表示未删除，1表示删除',
    primary key(`id`) 
) engine=InnoDB default charset=utf8mb4 comment '标签管理'
```

### 创建认证表
```
create table `blog_auth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key` varchar(20) DEFAULT '' COMMENT 'key',
    `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除，0表示未删除，1表示删除', 
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT="认证管理";
```