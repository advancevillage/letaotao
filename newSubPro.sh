#!/bin/bash

export root=$(pwd)

#校验目录
function CheckProject() {
    local p=$1
    if test -z "$p"; then
        ehco "[error] Project Name Is Empty!"
        exit 1
    fi
    local d="$root/$p"
    if test -d "$d"; then
        echo "[error] Project Name Has Exist!"
        exit 1
    fi
}
#创建项目
function CreateProject() {
    local p="$root/$1"
    local tree=(docs config template src deploy)
    for i in ${tree[*]}; do
        if test -d "$p/$i"; then
            continue
        else
            mkdir -p "$p/$i"
        fi
    done
}
#初始化项目
function InitProject() {
    local p="$root/$1"
    local d="$2"
    local me="$p/README.md"
    : > ${me}
    cat << EOF > ${me}
#### 项目名称
 * `basename ${p}`

#### 项目描述
    `echo ${d}`

#### 代码规范
    1: 命名规范
       - 变量命名
         * 小驼峰  eg: Id service GoodsId
       - 常量命名
         * 前缀知意 eg: ErrorQueueXXX = errors.New("xxx")
       - 函数命名
         * 动词+名词 eg: QueryMerchandiseById() error
         * 至少返回一个error类型
    2: SQL规范
         * fmt.Sprintf()
         eg: sql := fmt.Sprintf("" +
                    "select * " +
                    "from %s m " +
                    "where m.a = '%s' " +
                    ";",
                    TableName, param)
    3: DSL规范

#### 目录结构
    `basename ${p}`
       |
       |--- docs      API文档 Swagger
       |--- config    配置文件
       |--- deploy    部署脚本
       |--- template  通知消息模版
       |--- src       源代码
             ｜
             ｜--- init  初始化
             ｜--- main  启动入口
             ｜--- route 路由
             ｜--- ...
       |--- README.md 说明文件

     每个目录组成
       - model.go    定义数据类型
       - subject.go  文件主题
     eg:  init
            |--- model.go
            |--- single.go      //初始化信号
            |--- log.go         //初始化日志
            |--- config.go      //初始化配置
            |--- args.go        //初始化命令行参数

#### 构建编译
    go build -o bin/`basename ${p}` -gcflags "-N -l" -ldflags "-X main.commit=4d399017 -X main.version=v5.0.0"  src/main/`basename ${p}`.go

#### 参考文件
    1: swag Download: https://github.com/swaggo/swag/releases  [1.6.2]
       $ cp swag  /usr/sbin/
EOF
    local mod="$p/go.mod"
    cat << EOF > ${mod}
module `basename ${p}`

go 1.13

require (
	github.com/advancevillage/3rd  v1.0.0
)
EOF
    local src=(init main route)
    for i in ${src[*]}; do
        if test -d "$p/src/$i"; then
            continue
        else
            mkdir -p "$p/src/$i"
        fi
    done
    local main=`basename ${p}`
    main="$p/src/main/$main.go"
    cat << EOF > ${main}
// @title `basename ${p}`
// @version 0.0.1
// @description `echo ${d}`
// @contact.name richard sun
// @contact.email cugriver@163.com
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:13147
// @BasePath /v1
// @schemes http https
package main

var (
    commit  = "000000"
    version = "v0.0.1"
)

func main() {

}
EOF
    local config=`basename ${p}`
    config="$p/config/$config.xml"
    cat << EOF > ${config}
<?xml version="1.0" encoding="UTF-8"?>
<configure>
	<`basename ${p}`>
		<host>0.0.0.0</host>
		<port>13147</port>
	</`basename ${p}`>
</configure>
EOF
    local ignore="$p/.gitignore"
    cat << 'EOF' > ${ignore}
.idea*
.idea/vcs.xml
.DS_Store
pkg
*.log
EOF
    local deploy=`basename ${p}`
    deploy="$p/deploy/Dockerfile"
    cat << 'EOF' > ${deploy}

EOF
    local docs=`basename ${p}`
    docs="$p/docs/swagger.yml"
    cat << 'EOF' > ${docs}

EOF
    local template=`basename ${p}`
    template="$p/template/email.html"
    cat << 'EOF' > ${template}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>

</body>
</html>
EOF
    local init=`basename ${p}`
    init="$p/src/init/model.go"
    cat << 'EOF' > ${init}
package init
EOF
    local route=`basename ${p}`
    route="$p/src/route/model.go"
    cat << 'EOF' > ${route}
package route
EOF
}

function BindRepo() {
    local p=$1
    cd ${root}/${p}
    git init
    git add .
    git commit -m "first commit"
    git remote add origin https://github.com/advancevillage/${p}.git
    git push -u origin master
}

function BindSubModule() {
    cd ${root}
    local p=$1
    local repo="https://github.com/advancevillage/${p}"
    git submodule add ${repo}
}

#捕获命令行输入 项目名称
read -p "Enter Project Name: " project
read -p "Enter Project Description: " description

CheckProject   ${project}
CreateProject  ${project}
InitProject    ${project} ${description}
BindRepo       ${project}
BindSubModule  ${project}