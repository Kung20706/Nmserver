# 登入方式
使用MacID或其中一種+Time+Key
## 
 ``` golang 
cryptpassword := helper.CryptDIDTS(input.DeviceID + input.TimeStamp)
//傳來時間與字串 
	if cryptpassword == input.Checksum {
        //符合-->新增TOKEN進CACHE SERVER內 並給予一串到客戶端
    }
 ```
 checksum 是客戶端先用 MD5(自己選定其中一種ID+時間)+(KEY) 用來確認雙方密鑰及登入安全的方式
 這邊會先保留一份密鑰 (雙方都有的) key 
 客戶端選擇一個ID 及包進去的 constant.SecretKey ==>可以約定修改
 
## 環境參數設置 不同os有不同是配方式 WIN為例 以\來識別路徑 MAC 用/是配環境內的路徑
configFile := GetAppRoot() + `\config\project\` + GetAppEnv() + `\` + GetAppSite() + ".toml"
##
新增 本地網址方便瀏覽狀況
## DB schema 貼入 sql 內 使用localhost:3306 
-- Adminer 4.8.0 MySQL 5.6.51 dump

-- Adminer 4.8.0 MySQL 5.6.51 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) DEFAULT NULL,
  `children` int(10) DEFAULT NULL,
  `auth` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `device_id` varchar(255) DEFAULT NULL,
  `gid` varchar(50) DEFAULT NULL,
  `fid` varchar(50) DEFAULT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `alias` varchar(50) NOT NULL,
  `status` int(2) NOT NULL DEFAULT '0',
  `login_fail_count` int(11) NOT NULL DEFAULT '0',
  `is_freeze` tinyint(1) NOT NULL DEFAULT '0',
  `firstname` varchar(30) DEFAULT NULL,
  `lastname` varchar(30) DEFAULT NULL,
  `areacode` varchar(30) DEFAULT NULL,
  `birthday` varchar(30) DEFAULT NULL,
  `gender` varchar(30) DEFAULT NULL,
  `phone` varchar(30) DEFAULT NULL,
  `open` int(1) NOT NULL,
  `rewrite_password` int(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `reset_password` int(1) NOT NULL DEFAULT '0',
  `userid` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_admin_deleted_at` (`deleted_at`),
  KEY `idx_account_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `operation`;
CREATE TABLE `operation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL,
  `t` varchar(10) DEFAULT NULL,
  `t0` varchar(10) DEFAULT NULL,
  `t2` varchar(10) DEFAULT NULL,
  `t2y` varchar(10) DEFAULT NULL,
  `t2m` varchar(10) DEFAULT NULL,
  `t2d` varchar(10) DEFAULT NULL,
  `t2w` varchar(10) DEFAULT NULL,
  `t2l` varchar(10) DEFAULT NULL,
  `t2r` varchar(10) DEFAULT NULL,
  `t2n` varchar(10) DEFAULT NULL,
  `t2i` varchar(10) DEFAULT NULL,
  `o2` varchar(10) DEFAULT NULL,
  `p` varchar(10) DEFAULT NULL,
  `p2` varchar(10) DEFAULT NULL,
  `py` varchar(10) DEFAULT NULL,
  `pm` varchar(10) DEFAULT NULL,
  `pd` varchar(10) DEFAULT NULL,
  `pw` varchar(10) DEFAULT NULL,
  `pl` varchar(10) DEFAULT NULL,
  `z2r` varchar(10) DEFAULT NULL,
  `z2n` varchar(10) DEFAULT NULL,
  `c` varchar(10) DEFAULT NULL,
  `c2` varchar(10) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- 2021-06-22 09:11:18

-- 2021-06-12 08:14:18
#  run bot 
PROJECT_ENV=docker ./app schedule
## db 
docker-compose up -d db localhost:3306
docker-compose up -d adminer localhost:8099
adminer 提供db ui 可以下 sql cmd 
## redis
docker-compose up -d redis 
docker-compose up -d redis-adminer localhost:3200看redis 狀況用


# 用戶認證的流程
 ```
https://app.diagrams.net/?state=%7B%22folderId%22:%221JDykZy0xScEU1qdxmuDDTE92LyQU35_Z%22,%22action%22:%22create%22,%22userId%22:%22108743682938619798928%22%7D#G1cWUeJ0MxCJxTOp7UMkqMb1AhBZshqNUQ
 ```
# token 實作
https://kknews.cc/zh-tw/code/o3n62jm.html
# driver啟動
docker-compose up -d (db adminer redis) 

# 地端開發
swag init 
將產生的go docs/doc.go init  改為大寫 
go build -o app
PROJECT_ENV=local ./app server
# 容器building
docker-compose up -d app (依賴dockerfile產生)


# 本地building
go get github.com/swaggo/swag/cmd/swag 
swag init
sed "s/^func init/func Init/" docs/docs.go > docs/document.go 
go run main.go 出現顯示視窗


/Account/Internal/Login
登入帳號

 ``` json 
{
  "password": "qwer5678",
  "username": "beartest2021@gmail.com"
}
 ```
 回傳資訊 EX:
 ``` json 
{
  "error_text": "",
  "data": "8b80b4baa07b238658bb0de74584c6b6aa2c58787f7bf4cd933c6cd5fb6b54c2b8dff0a2e41a65e99722247e0c017025c91e48d18076d1e38ae9a95c80ba31264572d7965a90d42b4396f47d61b7aa38c645e5dae71853cf014a16935e2cead2ac1fac8e3b17d3fc88233da377cb85979ad11773976f5f704b292f9739f121f97ab6c0d0eeee82540c933bb58d9e1483edb14082"
}
 ```


##  

 

上面token↑
 /Account/List
放到json內

會依權限不同回傳不同的資訊



AREA CODE 改成地址 
POST 時間 
前端 