## Day 2

2022-01-25 22:58:53

gorm使用用户对象创建表时，用户结构体对象里面的峰驼字段会变成下划线方式，例如UserName到数据库中会变成user_name。

另外，id字段设置bigint新增记录会报错

> Error 1062: Duplicate entry '1' for key

用户对象

```go
type User struct {
	UserName  string `form:"username" binding:"required,min=3,max=10"`
	Password  string `form:"password" binding:"required"`
	Phone     string `form:"phone" binding:"required"`
	Role      string
	AvatarUrl string
	CreateAt  string
	UpdateAt  string
}
```

数据库字段

```mysql
create table users
(
    id         int auto_increment
        primary key,
    user_name  varchar(255) charset latin1     not null,
    password   varchar(255) charset latin1     null,
    role       varchar(255) default 'userRole' not null,
    phone      varchar(255)                    null,
    avatar_url varchar(255)                    null,
    create_at  varchar(255)                    null,
    update_at  varchar(255)                    null,
    constraint users_id_uindex
        unique (id)
)
    engine = MyISAM;
```


## Day 1

2022-01-19 23:12:16

1. jwt实现方式还是不太清楚！
