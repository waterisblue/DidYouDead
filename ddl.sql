create table DetectionLog
(
    id         int auto_increment
        primary key,
    account    varchar(40)                         not null,
    heartrate  int                                 null,
    createdate timestamp default CURRENT_TIMESTAMP null
);

create table FireService
(
    id            int auto_increment
        primary key,
    account       varchar(50)                         not null,
    sex           char(2)                             not null,
    idNum         varchar(50)                         null,
    locate        varchar(200)                        null,
    phoneNum      varchar(20)                         null,
    orderTime     timestamp                           null,
    funeralParlor varchar(50)                         null,
    fireService   varchar(50)                         null,
    urnStyle      varchar(50)                         null,
    cemetery      varchar(50)                         null,
    createDate    timestamp default CURRENT_TIMESTAMP null,
    Name          varchar(50)                         null,
    age           int                                 null,
    isActive      bit       default b'1'              null
);

create index IDX_ACCOUNT
    on FireService (account);

create index IDX_IDNUM
    on FireService (idNum);

create table SysUser
(
    id            int auto_increment
        primary key,
    account       varchar(50)                          null,
    password      varchar(50)                          not null,
    createdate    datetime   default CURRENT_TIMESTAMP null,
    administrator tinyint(1) default 0                 not null,
    constraint account
        unique (account)
);

create table Testament
(
    id                int auto_increment
        primary key,
    account           varchar(50)                          not null,
    testamentDetail   text                                 null,
    testamentStyle    varchar(20)                          null,
    testamentFileName varchar(50)                          null,
    createDate        timestamp  default CURRENT_TIMESTAMP null,
    testamentName     varchar(50)                          null,
    isActive          tinyint(1) default 0                 null
);

create index idx_account
    on Testament (account);

