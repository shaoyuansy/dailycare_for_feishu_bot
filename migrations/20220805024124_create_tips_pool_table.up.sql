CREATE TABLE informal_tips_pool (
	id INT auto_increment NOT NULL,
	`type` varchar(100) NOT NULL,
	content VARCHAR(255) NOT NULL,
	created_at datetime NULL COMMENT '创建时间',
	updated_at datetime NULL COMMENT '更新时间',
	deleted_at datetime NULL COMMENT '删除时间',
	CONSTRAINT informal_tips_pool_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
