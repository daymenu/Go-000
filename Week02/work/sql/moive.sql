CREATE DATABASE moive;

use moive;

CREATE TABLE moive 
(
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '影片名',
    `to_star` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '主演',
    `love` BIGINT NOT NULL DEFAULT '' COMMENT '喜爱度',
    INDEX name(name)
) ENGINE = "InnoDB";

INSERT INTO moive
(`name`, `to_star`)
VALUES 
('喜爱夜蒲', '连')
,('一路向西', '向西')
,('女集中营', '洪玉兰')
,('酒帘', '陈萍')
;