gin-scaffold
===

Go project generator, create PROJECT from SQL.

### How does it work?

Just like golang json encoding, when you want to marshal/unmarshal an object to json string, we add tag description for the object field. As an example:

````go
type JsonSomething struct{
  AField  int64     `json:"x"`
  BField  string    `json:"y"`
}
````
Similarly, scaffold use database table schemas' field comment as the tag description for generating code. For example,

````sql
CREATE TABLE `user_accounts` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT COMMENT 'caption:"No"',
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"Name" column:"y" update:"y" query:"like" widget:"text" valid:"required(),min(6),max(16)"',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '' COMMENT 'caption:"Email" column:"y" query:"like" widget:"email" valid:"required(),email()"',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"Sex" column:"y" update:"y" widget:"selection" relation:"user_accounts_sex"',
  `description` VARCHAR(256)     NOT NULL  DEFAULT '' COMMENT 'caption:"Description" update:"y" widget:"textarea"',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '' COMMENT 'caption:"Password" update:"y" widget:"password" valid:"required()"',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '' COMMENT 'caption:"Header Image" update:"y" widget:"file"',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'caption:"Status" column:"y" update:"y" query:"eq" widget:"selection" relation:"user_accounts_status"',
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT 'caption:"Create Time" widget:"datetime"',
  `updated_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'caption:"Update Time" column:"y" widget:"datetime"',
  `deleted_at`   TIMESTAMP       NULL  DEFAULT NULL  COMMENT 'caption:"Delete Time" gotype:"*time.Time" ignore:"y" widget:"datetime"'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'caption:"Members" index:"y" import:"y" export:"y"';

````

##### Quick Start

````shell
# define a [PROJECT_NAME] in Makefile
# generate a project
$: make start

# change mysql settings in [PROJECT_NAME]/configs/config.toml

# start the generated project
$: cd [GO_PATH]/src/[PROJECT_NAME]
$: make start
````

#### Thanks 

[liujianping/scaffold](https://github.com/liujianping/scaffold) go code generator

[LyricTian/gin-admin](https://github.com/LyricTian/gin-admin) awesome RBAC admin system